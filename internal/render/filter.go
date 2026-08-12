package render

import (
	"github.com/sagernet/sing-box/option"
)

// 输出变体字段白名单分组。
var (
	// DomainFields 域名类字段（含 DNS 专用 query_type）。
	DomainFields = []string{"domain", "domain_suffix", "domain_keyword", "domain_regex", "query_type"}
	// IPFields IP 类字段。
	IPFields = []string{"ip_cidr", "source_ip_cidr"}
)

// FilterFields 按白名单过滤一条 headless 规则：
// 只保留 fields 中列出的字段，其余丢弃；logical 规则递归处理。
// 返回 false 表示过滤后规则为空（无任何匹配字段），不应输出。
func FilterFields(rule option.HeadlessRule, fields []string) (option.HeadlessRule, bool) {
	// 空 fields = 全部保留（单输出缺省）
	if len(fields) == 0 {
		return rule, !IsEmpty(rule)
	}
	if rule.Type == "logical" {
		var kept []option.HeadlessRule
		for _, r := range rule.LogicalOptions.Rules {
			if fr, ok := FilterFields(r, fields); ok {
				kept = append(kept, fr)
			}
		}
		rule.LogicalOptions.Rules = kept
		return rule, len(kept) > 0
	}
	keep := make(map[string]bool, len(fields))
	for _, f := range fields {
		keep[f] = true
	}
	d := &rule.DefaultOptions
	if !keep["domain"] {
		d.Domain = nil
	}
	if !keep["domain_suffix"] {
		d.DomainSuffix = nil
	}
	if !keep["domain_keyword"] {
		d.DomainKeyword = nil
	}
	if !keep["domain_regex"] {
		d.DomainRegex = nil
	}
	if !keep["query_type"] {
		d.QueryType = nil
	}
	if !keep["ip_cidr"] {
		d.IPCIDR = nil
	}
	if !keep["source_ip_cidr"] {
		d.SourceIPCIDR = nil
	}
	if !keep["port"] {
		d.Port = nil
	}
	if !keep["source_port"] {
		d.SourcePort = nil
	}
	if !keep["port_range"] {
		d.PortRange = nil
	}
	if !keep["source_port_range"] {
		d.SourcePortRange = nil
	}
	if !keep["network"] {
		d.Network = nil
	}
	return rule, !IsEmpty(rule)
}

// FilterRuleSet 对整个规则集（head_rules + 主规则）按白名单过滤，
// 返回非空规则列表（顺序保持：head_rules 在前）。
func FilterRuleSet(headRules []option.HeadlessRule, main option.HeadlessRule, fields []string) []option.HeadlessRule {
	var out []option.HeadlessRule
	for _, r := range headRules {
		if fr, ok := FilterFields(r, fields); ok {
			out = append(out, fr)
		}
	}
	if fm, ok := FilterFields(main, fields); ok {
		out = append(out, fm)
	}
	return out
}
