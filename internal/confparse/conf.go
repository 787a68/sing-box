package confparse

import (
	"fmt"
	"os"

	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing/common/json"
)

// Source 描述一个上游规则源。
type Source struct {
	// URL 或 Path 二选一。
	URL string `json:"url,omitempty"`
	// 本地文件路径（相对 conf 文件所在目录）。
	Path string `json:"path,omitempty"`
	// 行格式：clash / host / qx，缺省 clash。
	Format string `json:"format,omitempty"`
}

// OutputSpec 定义一个输出变体（按字段类型拆分规则集）。
type OutputSpec struct {
	// Tag：输出规则集名（缺省 = conf 文件名）。
	Tag string `json:"tag,omitempty"`
	// Fields：白名单字段，只保留列出的字段；空 = 全部字段。
	Fields []string `json:"fields,omitempty"`
}

// Conf 是一个 sb 风格规则集定义（JSONC，文件名即规则集名）。
type Conf struct {
	Sources []Source `json:"sources"`
	// 头部规则：可归并字段（domain 系 / ip 系）参与语义去重；
	// 独立字段（query_type / logical 等）保持独立、输出在最前。
	HeadRules []option.HeadlessRule `json:"head_rules,omitempty"`
	// 排除：值级语义过滤，命中即从最终结果删除。
	Excludes []option.HeadlessRule `json:"exclude,omitempty"`
	// 输出变体：缺省 = 单输出（tag = conf 文件名，全部字段）。
	Outputs []OutputSpec `json:"outputs,omitempty"`
}

// Parse 读取并解析一个 .conf 文件（支持 // 注释）。
func Parse(path string) (*Conf, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseBytes(content)
}

// ParseBytes 解析 JSONC 内容。
func ParseBytes(content []byte) (*Conf, error) {
	conf, err := json.UnmarshalExtended[Conf](content)
	if err != nil {
		return nil, err
	}
	confPtr := conf
	if len(confPtr.Sources) == 0 {
		return nil, fmt.Errorf("confparse: at least one source is required")
	}
	for i, src := range confPtr.Sources {
		if src.URL == "" && src.Path == "" {
			return nil, fmt.Errorf("confparse: source[%d]: url or path is required", i)
		}
		if src.URL != "" && src.Path != "" {
			return nil, fmt.Errorf("confparse: source[%d]: url and path are mutually exclusive", i)
		}
		if src.Format == "" {
			confPtr.Sources[i].Format = "clash"
		}
	}
	for i, out := range confPtr.Outputs {
		if out.Tag == "" {
			return nil, fmt.Errorf("confparse: outputs[%d]: tag is required", i)
		}
	}
	return &confPtr, nil
}
