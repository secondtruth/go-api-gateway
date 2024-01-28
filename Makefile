.DEFAULT_GOAL := help
SHELL := /bin/bash

help: ## Prints this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-z0-9A-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build Go app for current OS/Arch
	go build ./...

test: ## Run all unit tests
	go test ./...

vet: ## Examines the source code and reports suspicious constructs
	go list ./... | grep -v /vendor/ | xargs -L1 go vet

format: ## Format the source code
	@find ./ -type f -name "*.go" -exec gofmt -w {} \;

lint: ## Lint all go source files
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

fmtcheck: ## Check if the source code has been formatted
	@mkdir -p target
	@find ./ -type f -name "*.go" -exec gofmt -d {} \; | tee target/format.diff
	@test ! -s target/format.diff || { echo "ERROR: the source code has not been formatted - please use 'make format' or 'gofmt'"; exit 1; }