package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

// diffSummary 输出 commit message 用的简洁摘要（按 conf 规则集统计 values 变化）：
//   第一行：总 values 增减（+add -del）
//   后续行：每个规则集一行  tag: values N (+add -del)
func diffSummary(oldPath, newPath string) (string, error) {
	oldRep, err := readBuildReport(oldPath)
	if err != nil {
		return "", fmt.Errorf("read old report: %w", err)
	}
	newRep, err := readBuildReport(newPath)
	if err != nil {
		return "", fmt.Errorf("read new report: %w", err)
	}
	oldByTag := map[string]ruleSetReport{}
	for _, rs := range oldRep.RuleSets {
		oldByTag[rs.Tag] = rs
	}
	totalAdd, totalDel := 0, 0
	var lines []string
	for _, rs := range newRep.RuleSets {
		old, ok := oldByTag[rs.Tag]
		add, del := rs.ValueCount, 0
		if ok {
			add = rs.ValueCount - old.ValueCount
			del = 0
			if add < 0 {
				del = -add
				add = 0
			}
		}
		totalAdd += add
		totalDel += del
		lines = append(lines, fmt.Sprintf("%s: values %d (+%d -%d)", rs.Tag, rs.ValueCount, add, del))
	}
	if len(lines) == 0 {
		return "", nil
	}
	return fmt.Sprintf("+%d -%d\n%s", totalAdd, totalDel, strings.Join(lines, "\n")), nil
}

// diffReport 对比新旧 report.json，输出上游活跃程度与规则质量变化摘要（markdown）。
func diffReports(oldPath, newPath string) error {
	oldRep, err := readBuildReport(oldPath)
	if err != nil {
		return fmt.Errorf("read old report: %w", err)
	}
	newRep, err := readBuildReport(newPath)
	if err != nil {
		return fmt.Errorf("read new report: %w", err)
	}

	oldByTag := map[string]ruleSetReport{}
	for _, rs := range oldRep.RuleSets {
		oldByTag[rs.Tag] = rs
	}

	fmt.Println("### Per Rule Set (new vs previous)")
	fmt.Println()
	fmt.Println("| tag | values | text-rm | sem-rm | excl-rm | src bytes | bin bytes |")
	fmt.Println("|---|---|---|---|---|---|---|")
	for _, rs := range newRep.RuleSets {
		old, ok := oldByTag[rs.Tag]
		if !ok {
			fmt.Printf("| %s (new) | %d | %d | %d | %d | %d | %d |\n",
				rs.Tag, rs.ValueCount, rs.TextRemoved, rs.SemanticRemoved, rs.ExcludeRemoved,
				rs.SourceBytes, rs.BinaryBytes)
			continue
		}
		fmt.Printf("| %s | %d (%+d) | %d (%+d) | %d (%+d) | %d (%+d) | %d (%+d) | %d (%+d) |\n",
			rs.Tag,
			rs.ValueCount, rs.ValueCount-old.ValueCount,
			rs.TextRemoved, rs.TextRemoved-old.TextRemoved,
			rs.SemanticRemoved, rs.SemanticRemoved-old.SemanticRemoved,
			rs.ExcludeRemoved, rs.ExcludeRemoved-old.ExcludeRemoved,
			rs.SourceBytes, rs.SourceBytes-old.SourceBytes,
			rs.BinaryBytes, rs.BinaryBytes-old.BinaryBytes)
	}

	fmt.Println()
	fmt.Println("### Upstream Sources (activity & quality)")
	fmt.Println()
	fmt.Println("| rule set | source | fmt | status | lines | parsed | skip rate |")
	fmt.Println("|---|---|---|---|---|---|---|")
	for _, rs := range newRep.RuleSets {
		for _, s := range rs.Sources {
			status := "ok"
			if !s.OK {
				status = "**FAIL**"
			}
			skipRate := "-"
			if s.Lines > 0 {
				skipRate = fmt.Sprintf("%.1f%%", float64(s.Lines-s.Parsed)*100/float64(s.Lines))
			}
			fmt.Printf("| %s | %s | %s | %s | %d | %d | %s |\n",
				rs.Tag, baseName(s.Name), s.Format, status, s.Lines, s.Parsed, skipRate)
		}
	}

	// 源可达性变化
	var failNew, failOld []string
	for _, rs := range newRep.RuleSets {
		for _, s := range rs.Sources {
			if !s.OK {
				failNew = append(failNew, s.Name)
			}
		}
	}
	for _, rs := range oldRep.RuleSets {
		for _, s := range rs.Sources {
			if !s.OK {
				failOld = append(failOld, s.Name)
			}
		}
	}
	if len(failNew) > 0 {
		sort.Strings(failNew)
		fmt.Println()
		fmt.Println("**failed sources this run:** " + strings.Join(failNew, ", "))
	}
	return nil
}

func readBuildReport(path string) (buildReport, error) {
	var rep buildReport
	content, err := os.ReadFile(path)
	if err != nil {
		return rep, err
	}
	err = json.Unmarshal(content, &rep)
	return rep, err
}
