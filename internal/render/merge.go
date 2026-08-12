package render

import (
	"github.com/sagernet/sing-box/option"

	"sing-box-rules/internal/transform"
)

// Merge 将多条同类型默认规则归并为一条（同字段值追加），保持出现顺序。
// logical 规则保留原样（由调用方置于独立位置）。
func Merge(rules []option.HeadlessRule) option.HeadlessRule {
	merged := transform.DefaultRule()
	for _, r := range rules {
		if r.Type != "default" && r.Type != "" {
			continue
		}
		d := &merged.DefaultOptions
		s := &r.DefaultOptions
		d.Domain = append(d.Domain, s.Domain...)
		d.DomainSuffix = append(d.DomainSuffix, s.DomainSuffix...)
		d.DomainKeyword = append(d.DomainKeyword, s.DomainKeyword...)
		d.DomainRegex = append(d.DomainRegex, s.DomainRegex...)
		d.IPCIDR = append(d.IPCIDR, s.IPCIDR...)
		d.SourceIPCIDR = append(d.SourceIPCIDR, s.SourceIPCIDR...)
		d.Port = append(d.Port, s.Port...)
		d.SourcePort = append(d.SourcePort, s.SourcePort...)
		d.PortRange = append(d.PortRange, s.PortRange...)
		d.SourcePortRange = append(d.SourcePortRange, s.SourcePortRange...)
		d.Network = append(d.Network, s.Network...)
		d.QueryType = append(d.QueryType, s.QueryType...)
	}
	return merged
}

// IsEmpty 判断归并后的规则是否没有任何匹配字段。
func IsEmpty(r option.HeadlessRule) bool {
	d := r.DefaultOptions
	return len(d.Domain) == 0 && len(d.DomainSuffix) == 0 &&
		len(d.DomainKeyword) == 0 && len(d.DomainRegex) == 0 &&
		len(d.IPCIDR) == 0 && len(d.SourceIPCIDR) == 0 &&
		len(d.Port) == 0 && len(d.SourcePort) == 0 &&
		len(d.PortRange) == 0 && len(d.SourcePortRange) == 0 &&
		len(d.Network) == 0 && len(d.QueryType) == 0
}
