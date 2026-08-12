package transform

import (
	"strings"

	"github.com/sagernet/sing-box/option"
)

// qx 行解析：Quantumult X 规则行 `host-suffix, value, policy`。
// 支持 host / host-suffix / host-keyword / host-wildcard / host-regex /
// ip-cidr / ip6-cidr；geoip / ip-asn 跳过。
func init() {
	Register("qx", ParseQXLine)
}

// ParseQXLine 解析一行 QX 规则。
func ParseQXLine(line string) (option.HeadlessRule, error) {
	parts := strings.Split(line, ",")
	if len(parts) < 2 {
		return option.HeadlessRule{}, Skip("unknown", line)
	}
	ruleType := strings.ToLower(strings.TrimSpace(parts[0]))
	value := strings.TrimSpace(parts[1]) // 第 3 段起为策略名（,reject / ,no-resolve），已按首个逗号切分剥除

	switch ruleType {
	case "host":
		return MakeRule("domain", []string{value}), nil
	case "host-suffix":
		return MakeRule("domain_suffix", []string{value}), nil
	case "host-keyword":
		return MakeRule("domain_keyword", []string{value}), nil
	case "host-wildcard":
		re, ok := wildcardToRegex(value)
		if !ok {
			return option.HeadlessRule{}, Skip("unknown", line)
		}
		return MakeRule("domain_regex", []string{re}), nil
	case "host-regex":
		return MakeRule("domain_regex", []string{value}), nil
	case "ip-cidr", "ip6-cidr":
		return MakeRule("ip_cidr", []string{value}), nil
	case "geoip":
		return option.HeadlessRule{}, Skip("geoip", line)
	case "ip-asn":
		return option.HeadlessRule{}, Skip("ip-asn", line)
	default:
		return option.HeadlessRule{}, Skip("unknown", line)
	}
}

// wildcardToRegex 将 QX 通配符（* 任意、? 单字符）转成锚定的 RE2 正则。
func wildcardToRegex(wildcard string) (string, bool) {
	if wildcard == "" || !strings.ContainsAny(wildcard, "*?") {
		return "", false
	}
	var b strings.Builder
	b.WriteString("^")
	escaped := false
	for _, c := range wildcard {
		if escaped {
			b.WriteString(regexpEscape(string(c)))
			escaped = false
			continue
		}
		switch c {
		case '\\':
			escaped = true
		case '*':
			b.WriteString(".*")
		case '?':
			b.WriteString(".")
		case '.', '+', '(', ')', '[', ']', '{', '}', '|', '^', '$':
			b.WriteString("\\")
			b.WriteRune(c)
		default:
			b.WriteRune(c)
		}
	}
	b.WriteString("$")
	return b.String(), true
}

func regexpEscape(s string) string {
	const meta = `\.+*?()|[]{}^$`
	if strings.ContainsRune(meta, rune(s[0])) {
		return "\\" + s
	}
	return s
}
