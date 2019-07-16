SRC = $(shell find . -type f -name '*.go')

.PHONY:setup
setup: ## Setup some tools
	GO111MODULE=off go get -u golang.org/x/lint/golint
	GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports

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

.PHONY: go-linux-build
go-linux-build: ## Run go linux build
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.buildVersion=$(VERSION)" -a -installsuffix cgo -o main .
