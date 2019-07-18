[![CircleCI](https://circleci.com/gh/nakamura244/dependency-check.svg?style=svg)](https://circleci.com/gh/nakamura244/dependency-check)
[![license](https://img.shields.io/github/license/srvc/wraperr.svg)](./LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/nakamura244/dependency-check)](https://goreportcard.com/report/github.com/nakamura244/dependency-check)

# Dependency check
This program is for dependency check (import check).

For example.
DIP(Dependency Inversion Principle) is an important factor in clean architecture etc.

This checker is a tool to check dependencies by paying attention to the import part

Dependent packages in go are limited to build-in packages, internal packages or third-party packages

Tool to check the dependencies of the above three packages


Developed based on `golang.org/x/tools/go/analysis`

# Installation
```console
go get github.com/nakamura244/dependency-check
```

# Getting Started
## Add Config
Add .config.yml file to your repository. like this

```yaml
.dependency-check/config.yml

base:
  innerPath: &innerPath github.com/nakamura244/dependency-check/testdata/src/valid

layer1:
  layerName: infrastructures
  packageNames:
    - github.com/nakamura244/dependency-check/testdata/src/valid/infrastructures
  innerPath: *innerPath
  allowDependPackages:
    - github.com/nakamura244/dependency-check/testdata/src/valid/domain
    - github.com/nakamura244/dependency-check/testdata/src/valid/interfaces
    - github.com/nakamura244/dependency-check/testdata/src/valid/config
    - github.com/nakamura244/dependency-check/testdata/src/valid/infrastructures/iface
  allowDependBuildIn: yes
  allowDependOutside: yes

layer2:
  layerName: interfaces
  packageNames:
    - github.com/nakamura244/dependency-check/testdata/src/valid/interfaces/database
    - github.com/nakamura244/dependency-check/testdata/src/valid/interfaces/api
    - github.com/nakamura244/dependency-check/testdata/src/valid/interfaces/log
  innerPath: *innerPath
  allowDependPackages:
    - github.com/nakamura244/dependency-check/testdata/src/valid/domain
    - github.com/nakamura244/dependency-check/testdata/src/valid/usecase
  allowDependBuildIn: yes
  allowDependOutside: no

layer3:
  layerName: usecase
  packageNames:
    - github.com/nakamura244/dependency-check/testdata/src/valid/usecase
  innerPath: *innerPath
  allowDependPackages:
    - github.com/nakamura244/dependency-check/testdata/src/valid/domain
  allowDependBuildIn: yes
  allowDependOutside: yes

layer4:
  layerName: domain
  packageNames:
    - github.com/nakamura244/dependency-check/testdata/src/valid/domain
  innerPath: *innerPath
  allowDependPackages:
  allowDependBuildIn: yes
  allowDependOutside: yes
```

### yml detail
You can set up to 4 layers or less. (layer1, layer2, layer3, layer4)


- base
  - innerPath ... Define internal repository path | required
- layerName ... Define layer name (single line can be set) | required
- packageNames ... Define package name belonging to layer (Multiple lines can be set) | required
- innerPath ...  Definer internal repository path belong ing to layer | required
- allowDependPackages ... Allow internal import package  (Multiple lines can be set)
  - If you allow all. write `all`
  - If you want to ban everything. `null`
- allowDependBuildIn ... Allow import of build-in packages | required (`yes` or `no`)
- allowDependOutside ... Allow import of third party package | required (`yes` or `no`)

## usages
```console
dependency-check 
DependencyChecker: DependencyChecker is checking dependency from imports (v0.0.1 rev:HEAD)

Usage: DependencyChecker [-flag] [package]


Flags:  -V      print version and exit
  -all
        no effect (deprecated)
  -c int
        display offending line with this many lines of context (default -1)
  -config string
        config path
  -cpuprofile string
        write CPU profile to this file
  -debug string
        debug flags, any subset of "fpstv"
  -flags
        print analyzer flags in JSON
  -ignoreTests
        ignoreTests bool
  -json
        emit JSON output
  -memprofile string
        write memory profile to this file
  -source
        no effect (deprecated)
  -tags string
        no effect (deprecated)
  -trace string
        write trace log to this file
  -v    no effect (deprecated)

```

## Run

### case valid
```console
dependency-check -ignoreTests=true -config=./testdata/src/valid/.dependency-check/config.yml ./testdata/src/valid/...
skip: /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/infrastructures/iface/http.go is not belong to anywhere layer.
skip: /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/infrastructures/iface/sql.go is not belong to anywhere layer.
skip: /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/config/config.go is not belong to anywhere layer.
checking[layer:domain] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/domain/a.go is ok 
skip: /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/tasks/g.go is not belong to anywhere layer.
checking[layer:domain] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/domain/b.go is ok 
checking[layer:interfaces] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/interfaces/database/c.go is ok 
checking[layer:domain] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/domain/c.go is ok 
checking[layer:interfaces] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/interfaces/log/e.go is ok 
checking[layer:infrastructures] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/infrastructures/a.go is ok 
checking[layer:interfaces] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/interfaces/log/f.go is ok 
checking[layer:infrastructures] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/infrastructures/b.go is ok 
checking[layer:usecase] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/usecase/h.go is ok 
checking[layer:interfaces] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/interfaces/database/d.go is ok 
checking[layer:infrastructures] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/infrastructures/c.go is ok 
checking[layer:interfaces] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/interfaces/api/a.go is ok 
checking[layer:usecase] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/usecase/i.go is ok 
checking[layer:interfaces] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/interfaces/api/b.go is ok 
checking[layer:infrastructures] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/valid/infrastructures/d.go is ok 
```
- Packages not registered in any layer will be skipped

### case invalid
```console
dependency-check -ignoreTests=true -config=./testdata/src/invalid-dependency-buildin/.dependency-check/config.yml ./testdata/src/invalid-dependency-buildin/...
checking[layer:layer1] file /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/invalid-dependency-buildin/layer1/a.go is ok 
DependencyChecker: found a violation of terms : /Users/a12091/go/src/github.com/nakamura244/dependency-check/testdata/src/invalid-dependency-buildin/layer2/b.go is not allow buildin package
```
- You will be warned if there is a violation


# Caution
If you are using CircleCI, you may need to set environment variables
like this (https://github.com/nakamura244/dependency-check/blob/master/.circleci/config.yml#L63)

Need to set GOROOT's environment variable
