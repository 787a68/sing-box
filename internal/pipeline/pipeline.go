package pipeline

import (
	"bufio"
	"bytes"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"sing-box-rules/internal/confparse"
	"sing-box-rules/internal/dedup"
	"sing-box-rules/internal/fetch"
	"sing-box-rules/internal/render"
	"sing-box-rules/internal/transform"

	"github.com/sagernet/sing-box/option"
)

// Config 构建参数。
type Config struct {
	ConfDir     string
	OutDir      string
	UseExamples bool
	Strict      bool
	Timeout     time.Duration
	Retries     int
}

// SourceStat 单个上游源的统计（用于活跃程度与质量评估）。
type SourceStat struct {
	Name     string         `json:"name"`
	Format   string         `json:"format"`
	OK       bool           `json:"ok"`       // 抓取是否成功
	Lines    int            `json:"lines"`    // 清洗后行数
	Parsed   int            `json:"parsed"`   // 成功解析数
	Skipped  map[string]int `json:"skipped"`  // 跳过分类统计
}

// Report 一次构建的统计。
type Report struct {
	Tag             string
	Skipped         map[string]int
	Sources         []SourceStat
	RuleCount       int // 输出 headless 规则条数
	ValueCount      int // 输出匹配值总数（字段值）
	TextRemoved     int // 文本去重删除数
	SemanticRemoved int // 语义去重删除数
	ExcludeRemoved  int // exclude 值级过滤删除数
	SourceJSON      string
	BinarySRS       string
}

// Build 处理一个 .conf 文件，产出 source JSON + binary SRS。
// confPath 的文件名（去扩展名）即规则集名。
// 返回每个 output 的 Report（多 output 拆分时可能有多个）。
func Build(confPath string, cfg Config) ([]*Report, error) {
	conf, err := confparse.Parse(confPath)
	if err != nil {
		return nil, fmt.Errorf("parse conf: %w", err)
	}
	tag := strings.TrimSuffix(filepath.Base(confPath), filepath.Ext(confPath))

	// 校验格式注册
	for _, src := range conf.Sources {
		if _, ok := transform.Get(src.Format); !ok {
			return nil, fmt.Errorf("unknown source format %q (available: %s)", src.Format, strings.Join(transform.Names(), ", "))
		}
	}
	// 校验 exclude 白名单
	for _, ex := range conf.Excludes {
		if err := transform.ValidateHeadlessRule(ex); err != nil {
			return nil, fmt.Errorf("exclude rule contains non-whitelist field: %v", err)
		}
	}

	// 抓取 + 清洗 + 解析
	results := fetch.FetchAll(conf.Sources, cfg.ConfDir, cfg.UseExamples, cfg.Timeout, cfg.Retries)
	skipped := map[string]int{}
	var rules []option.HeadlessRule
	var sources []SourceStat
	for i, res := range results {
		src := conf.Sources[i]
		stat := SourceStat{Name: res.Name, Format: src.Format, OK: res.Err == nil}
		if res.Err != nil {
			sources = append(sources, stat)
			if cfg.Strict {
				return nil, fmt.Errorf("fetch %s: %w", res.Name, res.Err)
			}
			fmt.Printf("  [warn] fetch %s: %v\n", res.Name, res.Err)
			continue
		}
		parseLines(res.Body, src.Format, &rules, &stat)
		mergeSkipped(skipped, stat.Skipped)
		sources = append(sources, stat)
	}

	// 头部规则拆分：可归并组（domain 系 / ip 系 / port 系）合并进主规则参与语义去重；
	// 独立组（query_type / logical 等）保持独立、输出在最前。
	mergeableHead, standaloneHead := splitHeadRules(conf.HeadRules)

	// 归并 + 去重 + 排除
	merged := render.Merge(append(mergeableHead, rules...))
	d := &merged.DefaultOptions
	textRemoved := removeCount(d.Domain, dedup.Strings(d.Domain))
	d.Domain = dedup.Strings(d.Domain)
	textRemoved += removeCount(d.DomainSuffix, dedup.Strings(d.DomainSuffix))
	d.DomainSuffix = dedup.Strings(d.DomainSuffix)
	textRemoved += removeCount(d.DomainKeyword, dedup.Strings(d.DomainKeyword))
	d.DomainKeyword = dedup.Strings(d.DomainKeyword)
	textRemoved += removeCount(d.DomainRegex, dedup.Strings(d.DomainRegex))
	d.DomainRegex = dedup.Strings(d.DomainRegex)
	textRemoved += removeCount(d.IPCIDR, dedup.Strings(d.IPCIDR))
	d.IPCIDR = dedup.Strings(d.IPCIDR)
	textRemoved += removeCount(d.SourceIPCIDR, dedup.Strings(d.SourceIPCIDR))
	d.SourceIPCIDR = dedup.Strings(d.SourceIPCIDR)

	semanticRemoved := removeCount(d.DomainSuffix, dedup.DomainSuffix(d.DomainSuffix))
	d.DomainSuffix = dedup.DomainSuffix(d.DomainSuffix)
	semanticRemoved += removeCount(d.DomainSuffix, dedup.SuffixByKeyword(d.DomainSuffix, d.DomainKeyword))
	d.DomainSuffix = dedup.SuffixByKeyword(d.DomainSuffix, d.DomainKeyword)
	semanticRemoved += removeCount(d.Domain, dedup.Domain(d.Domain, d.DomainKeyword, d.DomainSuffix))
	d.Domain = dedup.Domain(d.Domain, d.DomainKeyword, d.DomainSuffix)
	semanticRemoved += removeCount(d.IPCIDR, dedup.IPCIDR(d.IPCIDR))
	d.IPCIDR = dedup.IPCIDR(d.IPCIDR)
	semanticRemoved += removeCount(d.SourceIPCIDR, dedup.IPCIDR(d.SourceIPCIDR))
	d.SourceIPCIDR = dedup.IPCIDR(d.SourceIPCIDR)

	ex := render.Excludes(conf.Excludes)
	beforeValues := countValues([]option.HeadlessRule{merged})
	ex.Apply(&merged)
	excludeRemoved := beforeValues - countValues([]option.HeadlessRule{merged})

	// 输出变体：缺省 = 单输出（tag = conf 文件名，全部字段）
	outputs := conf.Outputs
	if len(outputs) == 0 {
		outputs = []confparse.OutputSpec{{Tag: tag}}
	}

	var reports []*Report
	for _, out := range outputs {
		final := render.FilterRuleSet(standaloneHead, merged, out.Fields)
		if len(final) == 0 {
			fmt.Printf("  [warn] output %s: empty after field filter, skipped\n", out.Tag)
			continue
		}
		// 校验最终产物白名单
		for _, r := range final {
			if err := transform.ValidateHeadlessRule(r); err != nil {
				return nil, fmt.Errorf("output rule contains non-whitelist field: %v", err)
			}
		}
		srcPath := filepath.Join(cfg.OutDir, "source", out.Tag+".json")
		if err := render.WriteSource(srcPath, final); err != nil {
			return nil, fmt.Errorf("write source: %w", err)
		}
		srsPath := filepath.Join(cfg.OutDir, "binary", out.Tag+".srs")
		if err := render.WriteSRS(srsPath, final); err != nil {
			return nil, fmt.Errorf("write srs: %w", err)
		}
		if err := render.CheckSRS(srsPath); err != nil {
			return nil, fmt.Errorf("srs self-check failed: %w", err)
		}
		reports = append(reports, &Report{
			Tag:             out.Tag,
			Skipped:         skipped,
			Sources:         sources,
			RuleCount:       countRules(final),
			ValueCount:      countValues(final),
			TextRemoved:     textRemoved,
			SemanticRemoved: semanticRemoved,
			ExcludeRemoved:  excludeRemoved,
			SourceJSON:      srcPath,
			BinarySRS:       srsPath,
		})
	}
	if len(reports) == 0 {
		return nil, fmt.Errorf("no output produced for %s", tag)
	}
	return reports, nil
}

// splitHeadRules 把头部规则拆成可归并组与独立组。
func splitHeadRules(headRules []option.HeadlessRule) (mergeable, standalone []option.HeadlessRule) {
	for _, r := range headRules {
		if r.Type == "logical" {
			standalone = append(standalone, r)
			continue
		}
		// query_type 保持独立（输出在最前）
		if len(r.DefaultOptions.QueryType) > 0 {
			standalone = append(standalone, r)
			continue
		}
		mergeable = append(mergeable, r)
	}
	return
}

// removeCount 返回 before 与 after 的数量差（用于去重统计）。
func removeCount(before, after []string) int {
	return len(before) - len(after)
}

// mergeSkipped 将单个 source 的跳过统计并入全局。
func mergeSkipped(total, part map[string]int) {
	for k, v := range part {
		total[k] += v
	}
}

// countRules 统计 headless 规则条数（含 logical 嵌套）。
func countRules(rules []option.HeadlessRule) int {
	n := 0
	for _, r := range rules {
		n++
		if r.Type == "logical" {
			n += countRules(r.LogicalOptions.Rules)
		}
	}
	return n
}

// countValues 统计匹配值总数（各字段值之和）。
func countValues(rules []option.HeadlessRule) int {
	n := 0
	for _, r := range rules {
		if r.Type == "logical" {
			n += countValues(r.LogicalOptions.Rules)
			continue
		}
		d := r.DefaultOptions
		n += len(d.Domain) + len(d.DomainSuffix) + len(d.DomainKeyword) + len(d.DomainRegex) +
			len(d.IPCIDR) + len(d.SourceIPCIDR) + len(d.Port) + len(d.SourcePort) +
			len(d.PortRange) + len(d.SourcePortRange) + len(d.Network) + len(d.QueryType)
	}
	return n
}

// parseLines 清洗并解析一个 source 的所有行，统计填充到 stat。
func parseLines(content []byte, format string, rules *[]option.HeadlessRule, stat *SourceStat) {
	parser, _ := transform.Get(format)
	stat.Skipped = map[string]int{}
	for _, line := range cleanLines(content) {
		stat.Lines++
		rule, err := parser(line)
		if err != nil {
			if skipErr, ok := err.(*transform.SkipError); ok {
				stat.Skipped[skipErr.Category]++
			} else {
				stat.Skipped["error"]++
			}
			continue
		}
		stat.Parsed++
		*rules = append(*rules, rule)
	}
}

// cleanLines 去除空行、注释行、行尾注释。
func cleanLines(content []byte) []string {
	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader(content))
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") || strings.HasPrefix(line, ";") {
			continue
		}
		line = stripTrailingComment(line)
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

// stripTrailingComment 剥离行尾注释（识别 URL 中的 // 不误伤）。
func stripTrailingComment(line string) string {
	inURL := false
	for i := 0; i < len(line); i++ {
		if i > 0 && line[i-1] == ':' && line[i] == '/' && i+1 < len(line) && line[i+1] == '/' {
			inURL = true
			i += 2
			continue
		}
		if !inURL && line[i] == ' ' && i+1 < len(line) && (line[i+1] == '#' || line[i+1] == ';' || (line[i+1] == '/' && i+2 < len(line) && line[i+2] == '/')) {
			return line[:i]
		}
		if inURL && line[i] == ' ' {
			inURL = false
		}
	}
	return line
}

// Summary 打印跳过统计（按分类排序）。
func (r *Report) Summary() string {
	if len(r.Skipped) == 0 {
		return fmt.Sprintf("%s: no skipped lines", r.Tag)
	}
	keys := make([]string, 0, len(r.Skipped))
	for k := range r.Skipped {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s(%d)", k, r.Skipped[k]))
	}
	return fmt.Sprintf("%s: skipped %d [%s]", r.Tag, len(r.Skipped), strings.Join(parts, " "))
}
