package transform

import (
	"net/netip"
	"strings"

	"github.com/sagernet/sing-box/option"
)

// host 行解析：每行一个域名（.google.com 或 google.com）→ domain_suffix；
// 或每行一个 IP / CIDR（含 ip-cidr, 前缀行）→ ip_cidr。
func init() {
	Register("host", ParseHostLine)
}

// ParseHostLine 解析一行 host 列表。
func ParseHostLine(line string) (option.HeadlessRule, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return option.HeadlessRule{}, Skip("unknown", line)
	}
	// 兼容 ip-cidr,1.2.3.0/24 写法（显式 IP 校验，非法值跳过）
	if idx := strings.Index(line, ","); idx >= 0 {
		head := strings.ToLower(strings.TrimSpace(line[:idx]))
		if head == "ip-cidr" || head == "ip6-cidr" {
			value := strings.TrimSpace(line[idx+1:])
			if _, err := netip.ParsePrefix(value); err != nil {
				if _, err := netip.ParseAddr(value); err != nil {
					return option.HeadlessRule{}, Skip("unknown", line)
				}
			}
			return MakeRule("ip_cidr", []string{value}), nil
		}
		return option.HeadlessRule{}, Skip("unknown", line)
	}
	// IP / CIDR
	trimmed := strings.TrimPrefix(line, ".")
	if _, err := netip.ParsePrefix(trimmed); err == nil {
		return MakeRule("ip_cidr", []string{trimmed}), nil
	}
	if _, err := netip.ParseAddr(trimmed); err == nil {
		return MakeRule("ip_cidr", []string{trimmed}), nil
	}
	// 域名 → domain_suffix（剥前导点）
	if !validDomain(trimmed) {
		return option.HeadlessRule{}, Skip("unknown", line)
	}
	return MakeRule("domain_suffix", []string{trimmed}), nil
}

func validDomain(s string) bool {
	if s == "" || len(s) > 253 {
		return false
	}
	label := false
	for _, c := range s {
		switch {
		case c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '-' || c == '_':
			label = true
		case c == '.':
			if !label {
				return false
			}
			label = false
		default:
			return false
		}
	}
	return label
}
