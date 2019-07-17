package main

import (
	"fmt"
	"github.com/nakamura244/dependency-check/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var (
	config string
	revision = "HEAD"
)
const version = "0.0.1"


func init() {

}

func main() {
	doc := fmt.Sprintf("DependencyChecker is checking dependency from imports (v%s rev:%s)", version, revision)
	analyzer.Analyzer.Doc = doc
	analyzer.Analyzer.Flags.Bool("ignoreTests", false, "ignoreTests bool")
	analyzer.Analyzer.Flags.StringVar(&config, "config", "", "config path")

	singlechecker.Main(analyzer.Analyzer)
}
