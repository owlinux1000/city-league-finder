help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run tests
	@go test ./...

run: ## Run the program
	@go run main.go

build:
	@go build -o bin/city-league-detector

.PHONY: help test run