package transform

import (
	"fmt"
	"strings"

	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/json/badoption"
)

// dnsTypeMap 常用 DNS 查询类型名 → 数值（与 sing-box mDNS.StringToType 一致）。
var dnsTypeMap = map[string]option.DNSQueryType{
	"A": 1, "NS": 2, "CNAME": 5, "SOA": 6, "PTR": 12, "MX": 15,
	"TXT": 16, "AAAA": 28, "SRV": 33, "NAPTR": 35, "ANY": 255,
	"CAA": 257, "SVCB": 64, "HTTPS": 65, "AXFR": 252, "IXFR": 251,
	"TSIG": 250, "TKEY": 249,
}

// MakeRule 按白名单字段构造单字段规则。
// field 必须是白名单字段；port 系列传入的是字符串（uint16 解析失败则丢弃）。
func MakeRule(field string, values []string) option.HeadlessRule {
	rule := DefaultRule()
	d := &rule.DefaultOptions
	appendStr := func(target *badoption.Listable[string]) {
		*target = append(*target, values...)
	}
	switch field {
	case "domain":
		appendStr(&d.Domain)
	case "domain_suffix":
		appendStr(&d.DomainSuffix)
	case "domain_keyword":
		appendStr(&d.DomainKeyword)
	case "domain_regex":
		appendStr(&d.DomainRegex)
	case "ip_cidr":
		appendStr(&d.IPCIDR)
	case "source_ip_cidr":
		appendStr(&d.SourceIPCIDR)
	case "port":
		for _, v := range values {
			if n, ok := ParseUint16(v); ok {
				d.Port = append(d.Port, n)
			}
		}
	case "source_port":
		for _, v := range values {
			if n, ok := ParseUint16(v); ok {
				d.SourcePort = append(d.SourcePort, n)
			}
		}
	case "port_range":
		appendStr(&d.PortRange)
	case "source_port_range":
		appendStr(&d.SourcePortRange)
	case "network":
		appendStr(&d.Network)
	case "query_type":
		for _, v := range values {
			if qt, ok := ParseQueryType(v); ok {
				d.QueryType = append(d.QueryType, qt)
			}
		}
	}
	return rule
}

// ParseQueryType 解析 DNS 查询类型（类型名或数字）。
func ParseQueryType(s string) (option.DNSQueryType, bool) {
	if qt, err := queryTypeByName(s); err == nil {
		return qt, true
	}
	var n uint16
	if _, err := fmt.Sscanf(s, "%d", &n); err == nil {
		return option.DNSQueryType(n), true
	}
	return 0, false
}

func queryTypeByName(s string) (option.DNSQueryType, error) {
	qt, loaded := dnsTypeMap[strings.ToUpper(s)]
	if !loaded {
		return 0, fmt.Errorf("unknown query type: %s", s)
	}
	return qt, nil
}

// ParseUint16 解析 0-65535 的十进制字符串。
func ParseUint16(s string) (uint16, bool) {
	if s == "" {
		return 0, false
	}
	var n uint16
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + uint16(c-'0')
		if n > 65535 {
			return 0, false
		}
	}
	return n, true
}

// ValidateHeadlessRule 校验规则只含白名单字段；
// 违反返回 *SkipError（分类 whitelist），否则返回 nil。
func ValidateHeadlessRule(rule option.HeadlessRule) error {
	switch rule.Type {
	case "", "default":
		return validateDefault(rule.DefaultOptions)
	case "logical":
		for _, r := range rule.LogicalOptions.Rules {
			if err := ValidateHeadlessRule(r); err != nil {
				return err
			}
		}
		return nil
	default:
		return Skip("whitelist", "unknown rule type")
	}
}

func validateDefault(d option.DefaultHeadlessRule) error {
	has := func(list badoption.Listable[string]) bool { return len(list) > 0 }
	switch {
	case has(d.ProcessName):
		return Skip("whitelist", "process_name")
	case has(d.ProcessPath):
		return Skip("whitelist", "process_path")
	case has(d.ProcessPathRegex):
		return Skip("whitelist", "process_path_regex")
	case has(d.PackageName):
		return Skip("whitelist", "package_name")
	case has(d.PackageNameRegex):
		return Skip("whitelist", "package_name_regex")
	case has(d.WIFISSID):
		return Skip("whitelist", "wifi_ssid")
	case has(d.WIFIBSSID):
		return Skip("whitelist", "wifi_bssid")
	case len(d.NetworkType) > 0:
		return Skip("whitelist", "network_type")
	case d.NetworkIsExpensive || d.NetworkIsConstrained:
		return Skip("whitelist", "network_is_*")
	case d.NetworkInterfaceAddress != nil && d.NetworkInterfaceAddress.Size() > 0:
		return Skip("whitelist", "network_interface_address")
	case len(d.DefaultInterfaceAddress) > 0:
		return Skip("whitelist", "default_interface_address")
	}
	return nil
}
