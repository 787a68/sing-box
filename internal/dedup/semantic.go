package dedup

import (
	"net/netip"
	"sort"
	"strings"
)

// 语义去重（与 quantumult-x 的 SemanticDedup 对齐）：
// 先按 key 排序（祖先在前），再单遍扫描——每条只与"更靠前且已保留"的条目比对，
// 被覆盖即删除（首条保留）。输出按原输入顺序还原。
// 复杂度 O(n·标签数 + n·keyword数)，无 O(n²) 全量扫描。

// DedupKeywords 语义去重：keyword 被更早(排序后靠前)的 keyword 子串覆盖则删除。
// 空串 keyword 直接丢弃（否则会误伤所有条目）。
func DedupKeywords(values []string) []string {
	idx := sortedByKeys(keys(values, identityKey))
	var kept []string
	var keptIdx []int
	for _, i := range idx {
		v := values[i]
		if v == "" {
			continue
		}
		covered := false
		for _, kw := range kept {
			if len(kw) <= len(v) && strings.Contains(v, kw) {
				covered = true
				break
			}
		}
		if !covered {
			kept = append(kept, v)
			keptIdx = append(keptIdx, i)
		}
	}
	sort.Ints(keptIdx) // 按原始顺序输出
	return reorder(values, keptIdx)
}

// DedupSuffixes 语义去重：被自身祖先 suffix 或被任一 keyword 覆盖则删除。
// keywords 需为已语义去重的列表。
func DedupSuffixes(values, keywords []string) []string {
	idx := sortedByKeys(keys(values, reversedDomainKey))
	set := make(map[string]struct{}, len(values))
	var keptIdx []int
	for _, i := range idx {
		v := values[i]
		if coveredBySuffix(v, set) || containsKeyword(v, keywords) {
			continue
		}
		set[trimDot(v)] = struct{}{}
		keptIdx = append(keptIdx, i)
	}
	sort.Ints(keptIdx) // 按原始顺序输出
	return reorder(values, keptIdx)
}

// DedupDomains 语义去重：被重复 domain（精确）、祖先 suffix 或任一 keyword 覆盖则删除。
// domain 是精确匹配，不能被祖先 domain 覆盖（与 QX hostSet 一致）。
// keywords / suffixes 需为已语义去重的列表。
func DedupDomains(values, keywords, suffixes []string) []string {
	set := make(map[string]struct{}, len(suffixes))
	for _, s := range suffixes {
		set[trimDot(s)] = struct{}{}
	}
	idx := sortedByKeys(keys(values, reversedDomainKey))
	domainSet := make(map[string]struct{}, len(values))
	var keptIdx []int
	for _, i := range idx {
		v := values[i]
		if coveredBySuffix(v, set) || coveredByDomain(v, domainSet) || containsKeyword(v, keywords) {
			continue
		}
		domainSet[trimDot(v)] = struct{}{}
		keptIdx = append(keptIdx, i)
	}
	sort.Ints(keptIdx) // 按原始顺序输出
	return reorder(values, keptIdx)
}

// coveredByDomain 判断 v 是否与 set 中的 domain 精确相等（domain 不匹配子域）。
func coveredByDomain(v string, set map[string]struct{}) bool {
	_, ok := set[trimDot(v)]
	return ok
}

// IPCIDR 语义去重：宽网段包含窄网段时删除窄段（保留更宽的），保持相对顺序。
func IPCIDR(values []string) []string {
	parsed := make([]netip.Prefix, len(values))
	valid := make([]bool, len(values))
	for i, v := range values {
		if p, err := netip.ParsePrefix(v); err == nil {
			parsed[i] = p.Masked()
			valid[i] = true
			continue
		}
		if a, err := netip.ParseAddr(v); err == nil {
			parsed[i] = netip.PrefixFrom(a, a.BitLen()).Masked()
			valid[i] = true
		}
	}
	idx := make([]int, 0, len(values))
	for i := range values {
		if valid[i] {
			idx = append(idx, i)
		}
	}
	// 宽度升序（宽网段在前），保证"宽覆盖窄"无论出现顺序都保留宽段
	sort.SliceStable(idx, func(a, b int) bool {
		return parsed[idx[a]].Bits() < parsed[idx[b]].Bits()
	})
	var kept []int
	for _, i := range idx {
		covered := false
		for _, k := range kept {
			if parsed[k].Contains(parsed[i].Addr()) {
				covered = true
				break
			}
		}
		if !covered {
			kept = append(kept, i)
		}
	}
	sort.Ints(kept) // 按原始顺序输出
	return reorder(values, kept)
}

// coveredBySuffix 判断 v 自身或其任意祖先标签域是否已在 set 中。
func coveredBySuffix(v string, set map[string]struct{}) bool {
	v = trimDot(v)
	for pos := 0; ; {
		if _, ok := set[v[pos:]]; ok {
			return true
		}
		dot := strings.IndexByte(v[pos:], '.')
		if dot < 0 {
			return false
		}
		pos += dot + 1
	}
}

// containsKeyword 判断 v 是否包含任一 keyword（长度预过滤，长 keyword 不可能命中短值）。
func containsKeyword(v string, keywords []string) bool {
	v = trimDot(v)
	for _, kw := range keywords {
		if kw != "" && len(kw) <= len(v) && strings.Contains(v, kw) {
			return true
		}
	}
	return false
}

// reversedDomainKey "sub.example.com" → "com.example.sub"（祖先按字典序在前）。
func reversedDomainKey(v string) string {
	v = trimDot(v)
	labels := strings.Split(v, ".")
	for i, j := 0, len(labels)-1; i < j; i, j = i+1, j-1 {
		labels[i], labels[j] = labels[j], labels[i]
	}
	return strings.Join(labels, ".")
}

func identityKey(v string) string {
	return v
}

// keys 预计算每条值的排序 key（避免排序比较器内重复计算）。
func keys(values []string, key func(string) string) []string {
	out := make([]string, len(values))
	for i, v := range values {
		out[i] = key(v)
	}
	return out
}

// sortedByKeys 按预计算的 key 稳定排序，返回原下标。
func sortedByKeys(keys []string) []int {
	idx := make([]int, len(keys))
	for i := range idx {
		idx[i] = i
	}
	sort.SliceStable(idx, func(a, b int) bool {
		return keys[idx[a]] < keys[idx[b]]
	})
	return idx
}

func trimDot(s string) string {
	return strings.TrimPrefix(s, ".")
}

func reorder(values []string, kept []int) []string {
	out := make([]string, 0, len(kept))
	for _, i := range kept {
		out = append(out, values[i])
	}
	return out
}
