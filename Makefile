SRC = $(shell find . -type f -name '*.go')
VERSION = $(shell godzil show-version)

.PHONY:setup
setup: ## Setup some tools
	GO111MODULE=off go get -u golang.org/x/lint/golint
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get -u github.com/Songmu/godzil/cmd/godzil

.PHONY:goimports
goimports: ## Run the goimports in all directories
	@goimports -w ${SRC}

.PHONY:test
test: ## Run test
	go test -race -v -cover ./...

.PHONY:lint
lint: ## Run the Golint in all directories
	golint -min_confidence 0.6 -set_exit_status ./...

.PHONY: vet
vet: ## Run vet
	go vet ./...

.PHONY: go-build
go-build:
	go build -ldflags="-X main.version=$(VERSION)"

.PHONY: install
install:
	mv dependency-check "$(shell go env GOPATH)/bin/"

