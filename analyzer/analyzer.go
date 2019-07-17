package analyzer

import (
	"errors"
	"fmt"
	"go/ast"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Config is yaml config
type Config struct {
	Base   base       `yaml:"base"`
	Layer1 configBody `yaml:"layer1"`
	Layer2 configBody `yaml:"layer2"`
	Layer3 configBody `yaml:"layer3"`
	Layer4 configBody `yaml:"layer4"`
}

type configBody struct {
	InnerPath           string   `yaml:"innerPath"`
	LayerName           string   `yaml:"layerName"`
	PackageNames        []string `yaml:"packageNames"`
	AllowDependPackages []string `yaml:"allowDependPackages"`
	AllowDependBuildIn  string   `yaml:"allowDependBuildIn"`
	AllowDependOutside  string   `yaml:"allowDependOutside"`
}

type base struct {
	InnerPath string `yaml:"innerPath"`
}

// Analyzer is analysis analyzer struct
var Analyzer = &analysis.Analyzer{
	Name:             "DependencyChecker",
	RunDespiteErrors: true,
	Run:              run,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
}

var (
	goRoot = os.Getenv("GOROOT")
)

func run(pass *analysis.Pass) (interface{}, error) {
	var err error
	ignoreTests := pass.Analyzer.Flags.Lookup("ignoreTests")
	config := pass.Analyzer.Flags.Lookup("config")
	buf, err := ioutil.ReadFile(config.Value.String())
	if err != nil {
		err = errors.New("do not read config")
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		err = errors.New("do not unmarshal config")
		return nil, err
	}

	fset := pass.Fset
	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}
	ins.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:

			filename := fset.File(n.Package).Name()

			if !strings.HasSuffix(filename, ".go") {
				break
			}

			if ignoreTests.Value.String() == "true" && strings.HasSuffix(filename, "_test.go") {
				break
			}

			configBody := fileMetadata(n.Name.Name, &c)

			if configBody.LayerName == "" {
				fmt.Printf("skip: %s is not belong to anywhere layer.\n", filename)
				break
			}

			buildInImports := makeBuildInInports(n.Imports)
			innerImports := makeInnerImports(n.Imports, configBody.InnerPath)
			outsideImports := makeOutSideImports(n.Imports, buildInImports, innerImports)

			// allowDependBuildIn check
			if !checkBuildIn(buildInImports, configBody) {
				err = fmt.Errorf(
					"found a violation of terms : %s is not allow buildin package",
					filename,
				)
			}
			// allowDependPackage check
			if !checkDependPackage(innerImports, configBody) {
				err = fmt.Errorf(
					"found a violation of terms: %s is not allow inner package",
					filename,
				)
			}

			// allowDependOutside check
			if !checkOutsidePackage(outsideImports, configBody) {
				err = fmt.Errorf(
					"found a violation of terms: %s is not allow outside package",
					filename,
				)
			}
			if err == nil {
				fmt.Printf("checking[layer:%s] file %s is ok \n", configBody.LayerName, filename)
			}
		}
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func fileMetadata(pName string, c *Config) *configBody {
	if find(pName, c.Layer1.PackageNames) {
		return &c.Layer1
	} else if find(pName, c.Layer2.PackageNames) {
		return &c.Layer2
	} else if find(pName, c.Layer3.PackageNames) {
		return &c.Layer3
	} else if find(pName, c.Layer4.PackageNames) {
		return &c.Layer4
	}
	return &configBody{}
}

func find(p string, ls []string) bool {
	for _, l := range ls {
		if strings.HasSuffix(l, "/"+p) {
			return true
		}
	}
	return false
}

func checkBuildIn(list []string, c *configBody) bool {
	if c.AllowDependBuildIn == "yes" {
		return true
	}
	for _, l := range list {
		bp := goRoot + "/src/" + l
		if isExist(bp) {
			return false
		}
	}
	return true
}

func checkDependPackage(list []string, c *configBody) bool {
	if len(list) == 0 {
		return true
	}

	if len(c.AllowDependPackages) > 0 && c.AllowDependPackages[0] == "all" {
		return true
	}

	if len(c.AllowDependPackages) == 0 && len(list) != 0 {
		return false
	}

	var f bool
	for _, l := range list {
		for _, p := range c.AllowDependPackages {
			if l == p {
				f = true
			}
		}
		if f == false {
			return false
		}
	}
	return true
}

func checkOutsidePackage(list []string, c *configBody) bool {
	if c.AllowDependOutside == "yes" {
		return true
	}
	if len(list) > 0 {
		return false
	}
	return true
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func makeBuildInInports(imports []*ast.ImportSpec) []string {
	var buidInImports []string
	for _, i := range imports {
		importName := trim(i.Path.Value)
		bp := goRoot + "/src/" + importName
		if isExist(bp) {
			buidInImports = append(buidInImports, importName)
		}
	}
	return buidInImports
}

func makeInnerImports(imports []*ast.ImportSpec, base string) []string {
	var innerImports []string
	for _, i := range imports {
		importName := trim(i.Path.Value)
		if strings.HasPrefix(importName, base) {
			innerImports = append(innerImports, importName)
		}
	}
	return innerImports
}

func makeOutSideImports(imports []*ast.ImportSpec, bi, ii []string) []string {
	var outSideImports []string
	ms := append(bi, ii...)
	for _, i := range imports {
		var match bool
		importName := trim(i.Path.Value)
		for _, m := range ms {
			if importName == m {
				match = true
			}
		}
		if !match {
			outSideImports = append(outSideImports, importName)
		}
	}
	return outSideImports
}

func trim(s string) string {
	s = strings.TrimSuffix(s, `"`)
	s = strings.TrimPrefix(s, `"`)
	return s
}
