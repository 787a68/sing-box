package transform

import (
	"fmt"
	"sort"

	"github.com/sagernet/sing-box/option"
)

// Parser 将一行输入解析为一条 headless 规则。
// 支持的行返回规则；不支持的行返回 *SkipError（计入报告而非致命）。
type Parser func(line string) (option.HeadlessRule, error)

// SkipError 表示该行不适用当前格式（平台限定、无等价、无法识别等）。
type SkipError struct {
	Category string // 统计分类：geoip / process / unsupported / unknown ...
	Line     string
}

func (e *SkipError) Error() string {
	return fmt.Sprintf("%s: %q", e.Category, e.Line)
}

// Skip 构造跳过错误。
func Skip(category, line string) error {
	return &SkipError{Category: category, Line: line}
}

var registry = map[string]Parser{}

// Register 注册一种行格式 parser。
func Register(name string, p Parser) {
	if p == nil {
		panic("transform: nil parser")
	}
	if _, dup := registry[name]; dup {
		panic("transform: duplicate parser: " + name)
	}
	registry[name] = p
}

// Get 按名取 parser。
func Get(name string) (Parser, bool) {
	p, ok := registry[name]
	return p, ok
}

// Names 返回已注册的格式名（排序）。
func Names() []string {
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// DefaultRule 构造一条默认类型规则。
func DefaultRule() option.HeadlessRule {
	return option.HeadlessRule{Type: "default"}
}
