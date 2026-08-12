package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"sing-box-rules/internal/pipeline"
)

// sourceStatReport 单个上游源统计（JSON 输出用）。
type sourceStatReport struct {
	Name     string         `json:"name"`
	Format   string         `json:"format"`
	OK       bool           `json:"ok"`
	Lines    int            `json:"lines"`
	Parsed   int            `json:"parsed"`
	Skipped  map[string]int `json:"skipped,omitempty"`
}

// ruleSetReport 单个规则集的统计。
type ruleSetReport struct {
	Tag             string            `json:"tag"`
	RuleCount       int               `json:"rules"`
	ValueCount      int               `json:"values"`
	Skipped         map[string]int    `json:"skipped,omitempty"`
	TextRemoved     int               `json:"text_removed"`
	SemanticRemoved int               `json:"semantic_removed"`
	ExcludeRemoved  int               `json:"exclude_removed"`
	SourceBytes     int               `json:"source_bytes"`
	BinaryBytes     int               `json:"binary_bytes"`
	Sources         []sourceStatReport `json:"sources"`
}

// buildReport 汇总所有规则集的报告。
type buildReport struct {
	GeneratedAt    time.Time       `json:"generated_at"`
	SingBoxVersion string          `json:"sing_box_version"`
	RuleSets       []ruleSetReport `json:"rule_sets"`
}

// writeReportJSON 输出机器可读报告（report.json）+ 规则集表格（rulesets.md）。
func writeReportJSON(outDir string, reports []*pipeline.Report) error {
	rep := buildReport{
		GeneratedAt:    time.Now().UTC(),
		SingBoxVersion: singBoxVersion(),
	}
	for _, r := range reports {
		ruleSet := ruleSetReport{
			Tag:             r.Tag,
			RuleCount:       r.RuleCount,
			ValueCount:      r.ValueCount,
			Skipped:         r.Skipped,
			TextRemoved:     r.TextRemoved,
			SemanticRemoved: r.SemanticRemoved,
			ExcludeRemoved:  r.ExcludeRemoved,
		}
		if fi, err := os.Stat(r.SourceJSON); err == nil {
			ruleSet.SourceBytes = int(fi.Size())
		}
		if fi, err := os.Stat(r.BinarySRS); err == nil {
			ruleSet.BinaryBytes = int(fi.Size())
		}
		for _, s := range r.Sources {
			ruleSet.Sources = append(ruleSet.Sources, sourceStatReport{
				Name:    s.Name,
				Format:  s.Format,
				OK:      s.OK,
				Lines:   s.Lines,
				Parsed:  s.Parsed,
				Skipped: s.Skipped,
			})
		}
		rep.RuleSets = append(rep.RuleSets, ruleSet)
	}
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(outDir, "report.json"), data, 0o644); err != nil {
		return err
	}
	return writeRuleSetsMD(outDir, rep)
}

// writeRuleSetsMD 输出 CHANGES.md 的规则集表格片段（消除 CI 对 jq 的依赖）。
func writeRuleSetsMD(outDir string, rep buildReport) error {
	var b strings.Builder
	b.WriteString("| tag | rules | values | text-rm | sem-rm | excl-rm | skipped | src bytes | bin bytes | sources |\n")
	b.WriteString("|---|---|---|---|---|---|---|---|---|---|\n")
	for _, rs := range rep.RuleSets {
		skipped := ""
		if len(rs.Skipped) > 0 {
			parts := make([]string, 0, len(rs.Skipped))
			for k, v := range rs.Skipped {
				parts = append(parts, fmt.Sprintf("%s:%d", k, v))
			}
			sort.Strings(parts)
			skipped = strings.Join(parts, ", ")
		}
		srcInfo := ""
		{
			var parts []string
			for _, s := range rs.Sources {
				mark := "ok"
				if !s.OK {
					mark = "fail"
				}
				parts = append(parts, fmt.Sprintf("%s(%s,%s,%d/%d)", baseName(s.Name), s.Format, mark, s.Parsed, s.Lines))
			}
			srcInfo = strings.Join(parts, "; ")
		}
		fmt.Fprintf(&b, "| %s | %d | %d | %d | %d | %d | %s | %d | %d | %s |\n",
			rs.Tag, rs.RuleCount, rs.ValueCount,
			rs.TextRemoved, rs.SemanticRemoved, rs.ExcludeRemoved,
			skipped, rs.SourceBytes, rs.BinaryBytes, srcInfo)
	}
	return os.WriteFile(filepath.Join(outDir, "rulesets.md"), []byte(b.String()), 0o644)
}

func baseName(path string) string {
	if i := strings.LastIndexAny(path, "/\\"); i >= 0 {
		return path[i+1:]
	}
	return path
}

// singBoxVersion 返回 go.mod 中固定的 sing-box 版本。
func singBoxVersion() string {
	content, err := os.ReadFile("go.mod")
	if err != nil {
		return "unknown"
	}
	for _, line := range strings.Split(string(content), "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[0] == "github.com/sagernet/sing-box" {
			return fields[1]
		}
	}
	return "unknown"
}
