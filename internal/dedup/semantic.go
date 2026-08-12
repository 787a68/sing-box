package dedup

import (
	"net/netip"
	"sort"
	"strings"
)

// DomainSuffix 语义去重：父域覆盖子域（保留更宽的）；被 keyword 覆盖的删除。
// 保持相对顺序。
func DomainSuffix(values []string) []string {
	// 按"标签数从少到多"排序后保留最宽的，再还原原序
	idx := sortedByWidth(values, suffixWidth)
	var kept []int
	for _, i := range idx {
		covered := false
		for _, k := range kept {
			if SuffixCovers(values[k], values[i]) {
				covered = true
				break
			}
		}
		if !covered {
			kept = append(kept, i)
		}
	}
	return reorder(values, kept)
}

// SuffixByKeyword 删除被任意 keyword 覆盖的后缀。
func SuffixByKeyword(suffixes, keywords []string) []string {
	if len(keywords) == 0 {
		return suffixes
	}
	var out []string
	for _, s := range suffixes {
		covered := false
		for _, kw := range keywords {
			if KeywordCovers(kw, s) {
				covered = true
				break
			}
		}
		if !covered {
			out = append(out, s)
		}
	}
	return out
}

// Domain 语义去重：被任意 keyword 或 suffix 覆盖的删除。
func Domain(domains, keywords, suffixes []string) []string {
	var out []string
	for _, d := range domains {
		covered := false
		for _, kw := range keywords {
			if KeywordCovers(kw, d) {
				covered = true
				break
			}
		}
		if !covered {
			for _, s := range suffixes {
				if SuffixCovers(s, d) {
					covered = true
					break
				}
			}
		}
		if !covered {
			out = append(out, d)
		}
	}
	return out
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

func sortedByWidth(values []string, width func(string) int) []int {
	idx := make([]int, len(values))
	for i := range idx {
		idx[i] = i
	}
	sort.SliceStable(idx, func(a, b int) bool {
		return width(values[idx[a]]) < width(values[idx[b]])
	})
	return idx
}

func suffixWidth(s string) int {
	return strings.Count(strings.TrimPrefix(s, "."), ".") + 1
}

func reorder(values []string, kept []int) []string {
	out := make([]string, 0, len(kept))
	for _, i := range kept {
		out = append(out, values[i])
	}
	return out
}
