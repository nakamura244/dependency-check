package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/findcall"
)

var testdata string

func init() {
	findcall.Analyzer.Flags.Bool("ignoreTests", true, "ignoreTests bool")
	testdata = analysistest.TestData()

	findcall.Analyzer.Flags.StringVar(&config, "config", "./testdata/src/valid-pattern/.dependency-check/config.yml", "config path")
}

func TestValid(t *testing.T) {
	analysistest.Run(t, testdata, findcall.Analyzer, "valid-pattern/config", "valid-pattern/domain", "valid-pattern/infrastructures", "valid-pattern/interfaces", "valid-pattern/tasks", "valid-pattern/usecase")
}

func TestInvalidDependecyPackage(t *testing.T) {
	findcall.Analyzer.Flags.Set("config", "./testdata/src/invalid-dependency-package/.dependency-check/config.yml")
	analysistest.Run(t, testdata, findcall.Analyzer, "invalid-dependency-package/layer1", "invalid-dependency-package/layer2")
}

func TestInvalidDependencyOutside(t *testing.T) {
	findcall.Analyzer.Flags.Set("config", "./testdata/src/invalid-dependency-outside/.dependency-check/config.yml")
	analysistest.Run(t, testdata, findcall.Analyzer, "invalid-dependency-package/layer1", "invalid-dependency-package/layer2")
}

func TestInvalidDependencyBuildin(t *testing.T) {
	findcall.Analyzer.Flags.Set("config", "./testdata/src/invalid-dependency-buildin/.dependency-check/config.yml")
	analysistest.Run(t, testdata, findcall.Analyzer, "invalid-dependency-package/layer1", "invalid-dependency-package/layer2")
}
