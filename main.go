package main

import (
	"github.com/nakamura244/dependency-check/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var (
	config string
)

func init() {
	analyzer.Analyzer.Flags.Bool("ignoreTests", false, "ignoreTests bool")
	analyzer.Analyzer.Flags.StringVar(&config, "config", "", "config path")
}

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
