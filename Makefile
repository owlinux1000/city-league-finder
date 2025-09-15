help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run tests
	@go test ./...

test/v : ## Run test with verbose mode
	@go test -v ./...

run: ## Run the program
	@go run main.go

build: ## Build the program
	@go build -o bin/city-league-finder main.go

release: ## Build the program with version

	@go build -ldflags "-X github.com/owlinux1000/city-league-finder/cmd.Version=0.0.1" -o bin/city-league-finder

fmt: ## Format the code
	@golangci-lint fmt

lint: ## Lint the code
	@golangci-lint run

lint/fix: ## Lint the code if it can be fixed by linters
	@golangci-lint run --fix

.PHONY: help test test/v run fmt lint lint/fix