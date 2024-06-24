.PHONY: test
test:
	go test ./...

.PHONY: install-dev-dependencies
install-dev-dependencies:
	go install github.com/matryer/moq@latest

.PHONY: generate
generate: install-dev-dependencies
	go generate ./...

.PHONY: help
help: ## Shows help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'