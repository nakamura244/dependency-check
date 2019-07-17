package analyzer

import (
	"go/ast"
	"reflect"
	"testing"
)

func Test_fileMetadata(t *testing.T) {
	tests := []struct {
		p string
		c *Config
		r *configBody
	}{
		{
			p: "ccc",
			c: &Config{Layer1: configBody{PackageNames: []string{"aaa/bbb/ccc"}}},
			r: &configBody{PackageNames: []string{"aaa/bbb/ccc"}},
		},
		{
			p: "ccc",
			c: &Config{Layer2: configBody{PackageNames: []string{"aaa/bbb/ccc"}}},
			r: &configBody{PackageNames: []string{"aaa/bbb/ccc"}},
		},
		{
			p: "ccc",
			c: &Config{Layer3: configBody{PackageNames: []string{"aaa/bbb/ccc"}}},
			r: &configBody{PackageNames: []string{"aaa/bbb/ccc"}},
		},
		{
			p: "ccc",
			c: &Config{Layer4: configBody{PackageNames: []string{"aaa/bbb/ccc"}}},
			r: &configBody{PackageNames: []string{"aaa/bbb/ccc"}},
		},
		{
			p: "ddd",
			c: &Config{Layer2: configBody{PackageNames: []string{"aaa/bbb/ccc"}}},
			r: &configBody{},
		},
	}
	for i, test := range tests {
		a := fileMetadata(test.p, test.c)
		if reflect.DeepEqual(a, test.r) == false {
			t.Errorf("%d, expected %+v, got %+v", i, test.r, a)
		}
	}
}

func Test_find(t *testing.T) {
	tests := []struct {
		list []string
		f    string
		r    bool
	}{
		{
			list: []string{"/aaa/bbb/ccc", "/aaa/bbb/ddd"},
			f:    "eee",
			r:    false,
		},
		{
			list: []string{"/aaa/bbb/ccc", "/aaa/bbb/ddd"},
			f:    "ddd",
			r:    true,
		},
	}
	for i, test := range tests {
		b := find(test.f, test.list)
		if b != test.r {
			t.Errorf("%d, expected  %v, actual %v", i, test.r, b)
		}
	}
}

func Test_checkBuildIn(t *testing.T) {
	tests := []struct {
		list []string
		c    *configBody
		r    bool
	}{
		{
			list: []string{},
			c:    &configBody{AllowDependBuildIn: "yes"},
			r:    true,
		},
		{
			list: []string{"os"},
			c:    &configBody{AllowDependBuildIn: "no"},
			r:    false,
		},
		{
			list: []string{"ossss"},
			c:    &configBody{AllowDependBuildIn: "no"},
			r:    true,
		},
	}

	for i, test := range tests {
		b := checkBuildIn(test.list, test.c)
		if b != test.r {
			t.Errorf("%d, expected  %v, actual %v", i, test.r, b)
		}
	}
}

func Test_checkDependPackage(t *testing.T) {
	tests := []struct {
		list []string
		c    *configBody
		r    bool
	}{
		{
			list: []string{},
			c:    &configBody{AllowDependPackages: []string{"aaa", "bbb"}},
			r:    true,
		},
		{
			list: []string{"aaa", "bbb"},
			c:    &configBody{AllowDependPackages: []string{"all"}},
			r:    true,
		},
		{
			list: []string{"aaa"},
			c:    &configBody{AllowDependPackages: []string{}},
			r:    false,
		},
		{
			list: []string{"ccc"},
			c:    &configBody{AllowDependPackages: []string{"aaa", "bbb"}},
			r:    false,
		},
		{
			list: []string{"bbb"},
			c:    &configBody{AllowDependPackages: []string{"aaa", "bbb"}},
			r:    true,
		},
	}
	for i, test := range tests {
		b := checkDependPackage(test.list, test.c)
		if b != test.r {
			t.Errorf("%d, expected  %v, actual %v", i, test.r, b)
		}
	}
}

func Test_checkOutsidePackage(t *testing.T) {
	tests := []struct {
		list []string
		c    *configBody
		r    bool
	}{
		{
			list: []string{"aaa", "bbb"},
			c:    &configBody{AllowDependOutside: "yes"},
			r:    true,
		},
		{
			list: []string{"aaa", "bbb"},
			c:    &configBody{AllowDependOutside: "yes"},
			r:    true,
		},
		{
			list: []string{"aaa", "bbb"},
			c:    &configBody{AllowDependOutside: "no"},
			r:    false,
		},
		{
			list: []string{},
			c:    &configBody{AllowDependOutside: "no"},
			r:    true,
		},
	}

	for i, test := range tests {
		b := checkOutsidePackage(test.list, test.c)
		if b != test.r {
			t.Errorf("%d, expected  %v, actual %v", i, test.r, b)
		}
	}
}

func Test_isExist(t *testing.T) {
	pass := goRoot + "/src/"
	tests := []struct {
		f string
		b bool
	}{
		{
			f: pass + "/database/sql",
			b: true,
		},
		{
			f: "/aaa/sql",
			b: false,
		},
	}

	for i, test := range tests {
		bool := isExist(test.f)
		if test.b != bool {
			t.Errorf("%d, expected  %v, actual %v", i, test.b, bool)
		}
	}
}

func Test_makeBuildInInports(t *testing.T) {
	var list []*ast.ImportSpec

	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "/aaa/bbb/"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "github.com/nakamura244/dependency-check/testdata/src/example1/config"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "go.uber.org/zap"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "/database/sql"}})

	b := makeBuildInInports(list)

	expected := []string{"/database/sql"}
	if !reflect.DeepEqual(b, expected) {
		t.Errorf("expected  %v, actual %v", expected, b)
	}
}

func Test_makeInnerImports(t *testing.T) {
	var list []*ast.ImportSpec

	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "/aaa/bbb/"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "github.com/nakamura244/dependency-check/testdata/src/example1/config"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "github.com/nakamura244/dependency-check/testdata/src/example1/domain"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "go.uber.org/zap"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "/database/sql"}})

	i := makeInnerImports(list, "github.com/nakamura244/dependency-check/testdata/src/example1/")
	expected := []string{"github.com/nakamura244/dependency-check/testdata/src/example1/config", "github.com/nakamura244/dependency-check/testdata/src/example1/domain"}

	if !reflect.DeepEqual(i, expected) {
		t.Errorf("expected  %v, actual %v", expected, i)
	}
}

func Test_makeOutSideImports(t *testing.T) {
	var list []*ast.ImportSpec

	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "/aaa/bbb/"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "github.com/nakamura244/dependency-check/testdata/src/example1/config"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "github.com/nakamura244/dependency-check/testdata/src/example1/domain"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "go.uber.org/zap"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "gopkg.in/go-playground/validator.v9"}})
	list = append(list, &ast.ImportSpec{Path: &ast.BasicLit{Value: "/database/sql"}})

	list1 := []string{"/aaa/bbb/", "/database/sql"}
	list2 := []string{"github.com/nakamura244/dependency-check/testdata/src/example1/config", "github.com/nakamura244/dependency-check/testdata/src/example1/domain"}

	o := makeOutSideImports(list, list1, list2)
	expected := []string{"go.uber.org/zap", "gopkg.in/go-playground/validator.v9"}
	if !reflect.DeepEqual(o, expected) {
		t.Errorf("expected  %v, actual %v", expected, o)
	}
}

func Test_trim(t *testing.T) {
	tests := []struct {
		s string
		r string
	}{
		{
			s: "\"/aaa/bb/ss/\"",
			r: "/aaa/bb/ss/",
		},
		{
			s: "\"/aaa/\"bb/ss/\"",
			r: "/aaa/\"bb/ss/",
		},
	}
	for i, test := range tests {
		actual := trim(test.s)
		if test.r != actual {
			t.Errorf("%d, expected  %v, actual %v", i, test.r, actual)
		}
	}
}
