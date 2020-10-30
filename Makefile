
build: test all ## Test and build binaries for local architecture
first: tools deps

.PHONY: tools
tools: ## Download and install all dev/code tools
	@echo "==> Installing dev tools"
	go get -u honnef.co/go/tools/cmd/staticcheck

.PHONY: deps
deps: ## Update dependencies to latest version
	go mod verify

.PHONY: test
test: ## Ensure that code matches best practices
	go test ./...
	staticcheck ./...
	go vet

.PHONY: help
help: ## Display this help message
	@echo "GNU make(1) targets:"
	@grep -E '^[a-zA-Z_.-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'


