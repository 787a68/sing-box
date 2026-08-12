package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"sing-box-rules/internal/pipeline"
	"sing-box-rules/internal/render"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	switch os.Args[1] {
	case "build":
		runBuild(os.Args[2:])
	case "check":
		runCheck(os.Args[2:])
	case "diff":
		runDiff(os.Args[2:])
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Print(`sing-box rules generator

Usage:
  sbrules build [flags]      build rule sets from .conf files
  sbrules check <file.srs>   verify a binary rule set
  sbrules diff [--summary] <old.json> <new.json>   compare two report.json (upstream activity & quality)

build flags:
  --conf-dir <dir>   directory containing .conf files (default ".")
  --out-dir <dir>    output directory (default "rules")
  --examples         use local files under <conf-dir>/examples instead of fetching URLs
  --strict           fail on fetch errors
  --timeout <dur>    HTTP timeout (default 30s)
  --retries <n>      HTTP retries per URL (default 2)
`)
}

func runBuild(args []string) {
	fs := flag.NewFlagSet("build", flag.ExitOnError)
	confDir := fs.String("conf-dir", ".", "directory containing .conf files")
	outDir := fs.String("out-dir", "rules", "output directory")
	useExamples := fs.Bool("examples", false, "use local files under conf-dir/examples")
	strict := fs.Bool("strict", false, "fail on fetch errors")
	timeout := fs.Duration("timeout", 30*time.Second, "HTTP timeout")
	retries := fs.Int("retries", 2, "HTTP retries per URL")
	fs.Parse(args)

	confPaths, err := filepath.Glob(filepath.Join(*confDir, "*.conf"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "glob: %v\n", err)
		os.Exit(1)
	}
	if len(confPaths) == 0 {
		fmt.Fprintf(os.Stderr, "no .conf files found in %s\n", *confDir)
		os.Exit(1)
	}

	cfg := pipeline.Config{
		ConfDir:     *confDir,
		OutDir:      *outDir,
		UseExamples: *useExamples,
		Strict:      *strict,
		Timeout:     *timeout,
		Retries:     *retries,
	}

	failed := false
	var reports []*pipeline.Report
	results := make([]buildResult, len(confPaths))
	var wg sync.WaitGroup
	for i, confPath := range confPaths {
		wg.Add(1)
		go func(i int, confPath string) {
			defer wg.Done()
			fmt.Printf("building %s\n", confPath)
			confReports, err := pipeline.Build(confPath, cfg)
			results[i] = buildResult{confReports: confReports, err: err}
		}(i, confPath)
	}
	wg.Wait()
	for _, res := range results {
		if res.err != nil {
			fmt.Fprintf(os.Stderr, "  error: %v\n", res.err)
			failed = true
			continue
		}
		reports = append(reports, res.confReports...)
		for _, r := range res.confReports {
			fmt.Printf("  %s\n", r.Summary())
		}
	}
	if len(reports) > 0 {
		if err := writeReportJSON(*outDir, reports); err != nil {
			fmt.Fprintf(os.Stderr, "write report: %v\n", err)
			failed = true
		} else {
			fmt.Printf("wrote %s\n", filepath.Join(*outDir, "report.json"))
		}
	}
	if failed {
		os.Exit(1)
	}
}

// buildResult 单个 conf 的构建结果（按输入顺序收集，保证 report.json 顺序确定）。
type buildResult struct {
	confReports []*pipeline.Report
	err         error
}

func runCheck(args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: sbrules check <file.srs>")
		os.Exit(2)
	}
	if err := render.CheckSRS(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "check failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("OK: %s\n", strings.TrimSpace(args[0]))
}

func runDiff(args []string) {
	summary := false
	if len(args) == 3 && args[0] == "--summary" {
		summary = true
		args = args[1:]
	}
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: sbrules diff [--summary] <old.json> <new.json>")
		os.Exit(2)
	}
	if summary {
		out, err := diffSummary(args[0], args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "diff failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(out)
		return
	}
	if err := diffReports(args[0], args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "diff failed: %v\n", err)
		os.Exit(1)
	}
}
