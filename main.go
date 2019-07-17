package main

import (
	"fmt"
	"github.com/nakamura244/dependency-check/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var (
	config string
)

func init() {
	doc := fmt.Sprintf("DependencyChecker is checking dependency from imports (v%s rev:%s)", version, revision)
	analyzer.Analyzer.Doc = doc
	analyzer.Analyzer.Flags.Bool("ignoreTests", false, "ignoreTests bool")
	analyzer.Analyzer.Flags.StringVar(&config, "config", "", "config path")
}

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
