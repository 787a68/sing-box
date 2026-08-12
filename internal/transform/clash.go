package transform

import (
	"regexp"
	"strings"

	"github.com/sagernet/sing-box/option"
)

// clash 行解析：DOMAIN[,6] / DOMAIN-SUFFIX / DOMAIN-KEYWORD / DOMAIN-REGEX /
// IP-CIDR / IP-CIDR6 / SRC-IP-CIDR / DST-PORT / SRC-PORT；no-resolve 后缀剥除；
// GEOIP / IP-ASN / PROCESS-* 等平台限定或无等价类别跳过。
func init() {
	Register("clash", ParseClashLine)
}

var clashPattern = regexp.MustCompile(`^\s*(?i)([A-Z][A-Z0-9-]*)\s*,\s*(.+?)\s*$`)

// ParseClashLine 解析一行 clash 规则。
func ParseClashLine(line string) (option.HeadlessRule, error) {
	m := clashPattern.FindStringSubmatch(line)
	if m == nil {
		return option.HeadlessRule{}, Skip("unknown", line)
	}
	ruleType := strings.ToUpper(m[1])
	value := strings.TrimSpace(m[2])
	if i := strings.Index(value, ","); i >= 0 {
		// 剥掉尾部策略名 / no-resolve 等附加参数
		value = strings.TrimSpace(value[:i])
	}

	switch ruleType {
	case "DOMAIN":
		return MakeRule("domain", []string{value}), nil
	case "DOMAIN-SUFFIX":
		return MakeRule("domain_suffix", []string{value}), nil
	case "DOMAIN-KEYWORD":
		return MakeRule("domain_keyword", []string{value}), nil
	case "DOMAIN-REGEX":
		return MakeRule("domain_regex", []string{value}), nil
	case "IP-CIDR", "IP-CIDR6":
		return MakeRule("ip_cidr", []string{value}), nil
	case "SRC-IP-CIDR":
		return MakeRule("source_ip_cidr", []string{value}), nil
	case "DST-PORT":
		if _, ok := ParseUint16(value); !ok {
			return option.HeadlessRule{}, Skip("unknown", line)
		}
		return MakeRule("port", []string{value}), nil
	case "SRC-PORT":
		if _, ok := ParseUint16(value); !ok {
			return option.HeadlessRule{}, Skip("unknown", line)
		}
		return MakeRule("source_port", []string{value}), nil
	case "GEOIP":
		return option.HeadlessRule{}, Skip("geoip", line)
	case "IP-ASN":
		return option.HeadlessRule{}, Skip("ip-asn", line)
	case "PROCESS-NAME", "PROCESS-PATH", "PROCESS-PATH-REGEX":
		return option.HeadlessRule{}, Skip("process", line)
	case "NETWORK", "RULE-SET", "MATCH", "USER-AGENT", "URL-REGEX", "IN-PORT", "IN-USER", "IN-NAME", "IN-TYPE":
		return option.HeadlessRule{}, Skip("unsupported", line)
	default:
		return option.HeadlessRule{}, Skip("unknown", line)
	}
}
