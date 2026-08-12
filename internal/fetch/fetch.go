package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"sing-box-rules/internal/confparse"
)

// ghAccelPrefix 环境变量 GH_ACCEL 指定的 GitHub raw 下载加速前缀（临时方案，不入仓库）。
// 设置后 https://raw.githubusercontent.com/... 重写为 <GH_ACCEL>/<原URL>；缺省为空 = 直连。
var ghAccelPrefix = strings.TrimSuffix(os.Getenv("GH_ACCEL"), "/")

// Source 抓取结果。
type Source struct {
	Name  string // 用于错误报告
	Body  []byte
	Err   error
}

// FetchAll 并发抓取所有 source。
// useExamples 时优先从 confDir/examples/ 读取本地文件（按 URL 或 path 的文件名）。
// retries 为每个 URL 的重试次数。
func FetchAll(sources []confparse.Source, confDir string, useExamples bool, timeout time.Duration, retries int) []Source {
	results := make([]Source, len(sources))
	var wg sync.WaitGroup
	sem := make(chan struct{}, 8)
	for i, src := range sources {
		wg.Add(1)
		go func(i int, src confparse.Source) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			results[i] = fetchOne(src, confDir, useExamples, timeout, retries)
		}(i, src)
	}
	wg.Wait()
	return results
}

func fetchOne(src confparse.Source, confDir string, useExamples bool, timeout time.Duration, retries int) Source {
	name := src.URL
	if src.Path != "" {
		name = src.Path
	}
	if !useExamples {
		if src.Path != "" {
			return readLocal(filepath.Join(confDir, src.Path))
		}
		return fetchURL(src.URL, timeout, retries)
	}
	// examples 模式：按文件名在 examples/ 下找本地文件
	base := filepath.Base(name)
	candidates := []string{
		filepath.Join(confDir, "examples", base),
		filepath.Join(confDir, "examples", base+".txt"),
	}
	for _, p := range candidates {
		if res := readLocal(p); res.Err == nil {
			return res
		}
	}
	return Source{Name: name, Err: fmt.Errorf("examples file not found: %v", candidates)}
}

func readLocal(path string) Source {
	content, err := os.ReadFile(path)
	return Source{Name: path, Body: content, Err: err}
}

// maybeAccelURL 应用 GH_ACCEL 前缀重写（仅 GitHub raw 直链，其余 URL 不变）。
func maybeAccelURL(url string) string {
	if ghAccelPrefix == "" || !strings.HasPrefix(url, "https://raw.githubusercontent.com/") {
		return url
	}
	return ghAccelPrefix + "/" + url
}

func fetchURL(url string, timeout time.Duration, retries int) Source {
	url = maybeAccelURL(url)
	client := &http.Client{Timeout: timeout}
	var lastErr error
	for attempt := 0; attempt <= retries; attempt++ {
		resp, err := client.Get(url)
		if err != nil {
			lastErr = err
			if attempt < retries {
				time.Sleep(time.Duration(attempt+1) * time.Second)
			}
			continue
		}
		body, readErr := io.ReadAll(io.LimitReader(resp.Body, 64<<20))
		resp.Body.Close()
		if readErr == nil && resp.StatusCode == http.StatusOK {
			return Source{Name: url, Body: body}
		}
		if readErr != nil {
			lastErr = readErr
		} else {
			lastErr = fmt.Errorf("status %d", resp.StatusCode)
		}
		if attempt < retries {
			time.Sleep(time.Duration(attempt+1) * time.Second)
		}
	}
	return Source{Name: url, Err: lastErr}
}
