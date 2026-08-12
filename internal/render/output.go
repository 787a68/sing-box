package render

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sagernet/sing-box/common/srs"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/option"
)

// WriteSource 输出 rule-set source 格式（JSON，version 5，缩进美化）。
func WriteSource(path string, rules []option.HeadlessRule) error {
	compat := option.PlainRuleSetCompat{
		Version: C.RuleSetVersion5,
		Options: option.PlainRuleSet{Rules: rules},
	}
	data, err := json.Marshal(&compat)
	if err != nil {
		return err
	}
	var indented bytes.Buffer
	if err := json.Indent(&indented, data, "", "  "); err != nil {
		return err
	}
	return writeFile(path, indented.Bytes())
}

// WriteSRS 输出 rule-set binary（.srs，当前版本 5）。
func WriteSRS(path string, rules []option.HeadlessRule) error {
	var buf bytes.Buffer
	if err := srs.Write(&buf, option.PlainRuleSet{Rules: rules}, C.RuleSetVersion5); err != nil {
		return err
	}
	return writeFile(path, buf.Bytes())
}

// CheckSRS 回读校验 .srs 合法且可升级到当前版本。
func CheckSRS(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	compat, err := srs.Read(bytes.NewReader(content), false)
	if err != nil {
		return err
	}
	if _, err := compat.Upgrade(); err != nil {
		return err
	}
	return nil
}

func writeFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return err
	}
	fmt.Println("  wrote", path, fmt.Sprintf("(%d bytes)", len(data)))
	return nil
}
