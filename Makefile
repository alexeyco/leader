GOPATH:=$(shell go env GOPATH)

all: help

.PHONY: install
install: ## Install dependencies
	@if [ ! -f go.mod ]; then go mod init; fi
	@go mod vendor

.PHONY: lint
lint: ## Lint source code
	@golangci-lint --exclude-use-default=false run ./...

.PHONY: test
test: ## Run tests
	@if [ -f coverage.out ]; then rm coverage.out; fi
	@go test -v -covermode=count -coverprofile=coverage.out \
		github.com/alexeyco/leader \
		github.com/alexeyco/leader/etcd

	@go tool cover -func=coverage.out
	@rm coverage.out

.PHONY: etcd
etcd: ## Run etcd example
	@go run examples/etcd/main.go

help: ## Help
	@echo "Available actions:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo
