.DEFAULT_GOAL := help

# AutoDoc
# -------------------------------------------------------------------------
.PHONY: help
help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: tidy
tidy: ## Run golangci-lint, goimports and gofmt
	golangci-lint run ./... && goimports -w  . && gofmt -s -w -e -d .

.PHONY: tests
tests: ## Run integration and unit tests
	go test ./... -cover -coverpkg=./... --tags=unit,integration
