SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} ## Compile the code, build Executable File
	GOARCH=amd64 $(GOCMD) build -v -o $(BINARY_DIR)/api-linux-amd64 ./cmd/api

run: ## Start application
	$(GOCMD) run ./cmd/api

test: ## Run tests
	$(GOCMD) test ./... -cover

test-coverage: ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out

deps: ## Install dependencies
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

deps-cleancache: ## Clear cache in Go module
	$(GOCMD) clean -modcache

wire: ## Generate wire_gen.go
	cd pkg/di && wire

swag: ## Generate swagger docs
	swag init -g pkg/api/server.go -o ./cmd/api/docs

mockgen: # Generate mock files for the test
	mockgen -source=pkg/repository/interface/user.go -destination=pkg/repository/mockrepo/user_mock.go -package=mockrepo
	mockgen -source=pkg/usecase/interface/user.go -destination=pkg/usecase/mockusecase/user_mock.go -package=mockusecase

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
