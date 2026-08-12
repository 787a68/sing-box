package dedup

import "strings"

// Strings 文本去重：完全相同的值只保留第一个，保持顺序。
func Strings(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	out := make([]string, 0, len(values))
	for _, v := range values {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

// SuffixCovers 判断 parent 是否覆盖 child（child 是 parent 或 parent 的子域）。
// 忽略前导点。
func SuffixCovers(parent, child string) bool {
	p := strings.TrimPrefix(parent, ".")
	c := strings.TrimPrefix(child, ".")
	return c == p || strings.HasSuffix(c, "."+p)
}

// KeywordCovers 判断 kw 是否覆盖域名（kw 是域名的任意位置子串）。
func KeywordCovers(kw, domain string) bool {
	return strings.Contains(strings.TrimPrefix(domain, "."), kw)
}
