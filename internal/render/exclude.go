package render

import (
	"net/netip"
	"regexp"
	"strings"

	"github.com/sagernet/sing-box/option"
)

// Excludes 是一组用于值级过滤的 headless 规则（只允许白名单字段）。
type Excludes []option.HeadlessRule

// Apply 对单条默认规则做值级过滤：某个值被任一 exclude 条件覆盖即删除。
func (ex Excludes) Apply(r *option.HeadlessRule) {
	if len(ex) == 0 {
		return
	}
	d := &r.DefaultOptions
	d.Domain = filterStr(d.Domain, func(v string) bool { return !ex.coveredDomain(v) })
	d.DomainSuffix = filterStr(d.DomainSuffix, func(v string) bool { return !ex.coveredDomain(v) })
	d.DomainKeyword = filterStr(d.DomainKeyword, func(v string) bool { return !ex.coveredDomain(v) })
	d.DomainRegex = filterStr(d.DomainRegex, func(v string) bool { return !ex.coveredDomain(v) })
	d.IPCIDR = filterStr(d.IPCIDR, func(v string) bool { return !ex.coveredIP(v) })
	d.SourceIPCIDR = filterStr(d.SourceIPCIDR, func(v string) bool { return !ex.coveredIP(v) })
	d.Port = filterUint16(d.Port, func(v uint16) bool { return !ex.coveredPort(v) })
	d.SourcePort = filterUint16(d.SourcePort, func(v uint16) bool { return !ex.coveredPort(v) })
	d.PortRange = filterStr(d.PortRange, func(v string) bool { return !ex.coveredStr(v) })
	d.SourcePortRange = filterStr(d.SourcePortRange, func(v string) bool { return !ex.coveredStr(v) })
	d.Network = filterStr(d.Network, func(v string) bool { return !ex.coveredStr(v) })
}

// coveredDomain 检查候选域名字符串是否被任一 exclude 的域名系字段覆盖。
func (ex Excludes) coveredDomain(value string) bool {
	for _, e := range ex {
		if e.Type != "default" && e.Type != "" {
			continue
		}
		d := e.DefaultOptions
		for _, v := range d.Domain {
			if v == value {
				return true
			}
		}
		for _, v := range d.DomainSuffix {
			value = strings.TrimPrefix(value, ".")
			s := strings.TrimPrefix(v, ".")
			if value == s || strings.HasSuffix(value, "."+s) {
				return true
			}
		}
		for _, v := range d.DomainKeyword {
			if strings.Contains(strings.TrimPrefix(value, "."), v) {
				return true
			}
		}
		for _, v := range d.DomainRegex {
			if re, err := regexp.Compile(v); err == nil && re.MatchString(value) {
				return true
			}
		}
	}
	return false
}

// coveredIP 检查候选 CIDR/IP 是否被任一 exclude 的 ip_cidr 字段包含。
func (ex Excludes) coveredIP(value string) bool {
	prefix, ok := parsePrefix(value)
	if !ok {
		return false
	}
	for _, e := range ex {
		if e.Type != "default" && e.Type != "" {
			continue
		}
		for _, v := range e.DefaultOptions.IPCIDR {
			p, ok := parsePrefix(v)
			if ok && p.Contains(prefix.Addr()) {
				return true
			}
		}
	}
	return false
}

func (ex Excludes) coveredPort(v uint16) bool {
	for _, e := range ex {
		if e.Type != "default" && e.Type != "" {
			continue
		}
		for _, p := range e.DefaultOptions.Port {
			if p == v {
				return true
			}
		}
	}
	return false
}

// coveredStr 精确匹配（port_range / network 等字符串字段）。
func (ex Excludes) coveredStr(value string) bool {
	for _, e := range ex {
		if e.Type != "default" && e.Type != "" {
			continue
		}
		d := e.DefaultOptions
		if containsStr(d.PortRange, value) || containsStr(d.SourcePortRange, value) ||
			containsStr(d.Network, value) {
			return true
		}
	}
	return false
}

func parsePrefix(s string) (netip.Prefix, bool) {
	if p, err := netip.ParsePrefix(s); err == nil {
		return p.Masked(), true
	}
	if a, err := netip.ParseAddr(s); err == nil {
		return netip.PrefixFrom(a, a.BitLen()).Masked(), true
	}
	return netip.Prefix{}, false
}

func containsStr(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}

func filterStr(list []string, keep func(string) bool) []string {
	var out []string
	for _, v := range list {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func filterUint16(list []uint16, keep func(uint16) bool) []uint16 {
	var out []uint16
	for _, v := range list {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}
