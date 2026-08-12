package render

import (
	"net/netip"
	"regexp"

	"github.com/sagernet/sing-box/option"

	"sing-box-rules/internal/dedup"
)

// Excludes 是一组用于值级过滤的 headless 规则（只允许白名单字段）。
type Excludes []option.HeadlessRule

// compiledExclude 一条 exclude 规则 + 预编译正则（避免对每个候选值重复编译）。
type compiledExclude struct {
	d  option.DefaultHeadlessRule
	re []*regexp.Regexp
}

type compiledExcludes []compiledExclude

func (ex Excludes) compile() compiledExcludes {
	out := make([]compiledExclude, 0, len(ex))
	for _, e := range ex {
		if e.Type != "default" && e.Type != "" {
			continue
		}
		c := compiledExclude{d: e.DefaultOptions}
		for _, re := range e.DefaultOptions.DomainRegex {
			if r, err := regexp.Compile(re); err == nil {
				c.re = append(c.re, r)
			}
		}
		out = append(out, c)
	}
	return out
}

// Apply 对单条默认规则做值级过滤：某个值被任一 exclude 条件覆盖即删除。
func (ex Excludes) Apply(r *option.HeadlessRule) {
	if len(ex) == 0 {
		return
	}
	compiled := ex.compile()
	d := &r.DefaultOptions
	d.Domain = filterList(d.Domain, func(v string) bool { return !compiled.coveredDomain(v) })
	d.DomainSuffix = filterList(d.DomainSuffix, func(v string) bool { return !compiled.coveredDomain(v) })
	d.DomainKeyword = filterList(d.DomainKeyword, func(v string) bool { return !compiled.coveredDomain(v) })
	d.DomainRegex = filterList(d.DomainRegex, func(v string) bool { return !compiled.coveredDomain(v) })
	d.IPCIDR = filterList(d.IPCIDR, func(v string) bool { return !compiled.coveredIP(v) })
	d.SourceIPCIDR = filterList(d.SourceIPCIDR, func(v string) bool { return !compiled.coveredIP(v) })
	d.Port = filterList(d.Port, func(v uint16) bool { return !compiled.coveredPort(v) })
	d.SourcePort = filterList(d.SourcePort, func(v uint16) bool { return !compiled.coveredPort(v) })
	d.PortRange = filterList(d.PortRange, func(v string) bool { return !compiled.coveredStr(v) })
	d.SourcePortRange = filterList(d.SourcePortRange, func(v string) bool { return !compiled.coveredStr(v) })
	d.Network = filterList(d.Network, func(v string) bool { return !compiled.coveredStr(v) })
}

// coveredDomain 检查候选域名字符串是否被任一 exclude 的域名系字段覆盖。
func (cs compiledExcludes) coveredDomain(value string) bool {
	for _, c := range cs {
		d := c.d
		for _, v := range d.Domain {
			if v == value {
				return true
			}
		}
		for _, v := range d.DomainSuffix {
			if dedup.SuffixCovers(v, value) {
				return true
			}
		}
		for _, v := range d.DomainKeyword {
			if dedup.KeywordCovers(v, value) {
				return true
			}
		}
		for _, re := range c.re {
			if re.MatchString(value) {
				return true
			}
		}
	}
	return false
}

// coveredIP 检查候选 CIDR/IP 是否被任一 exclude 的 ip_cidr 字段包含。
func (cs compiledExcludes) coveredIP(value string) bool {
	prefix, ok := parsePrefix(value)
	if !ok {
		return false
	}
	for _, c := range cs {
		for _, v := range c.d.IPCIDR {
			p, ok := parsePrefix(v)
			if ok && p.Contains(prefix.Addr()) {
				return true
			}
		}
	}
	return false
}

func (cs compiledExcludes) coveredPort(v uint16) bool {
	for _, c := range cs {
		for _, p := range c.d.Port {
			if p == v {
				return true
			}
		}
	}
	return false
}

// coveredStr 精确匹配（port_range / network 等字符串字段）。
func (cs compiledExcludes) coveredStr(value string) bool {
	for _, c := range cs {
		d := c.d
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

func filterList[T ~[]E, E any](list T, keep func(E) bool) T {
	var out T
	for _, v := range list {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}
