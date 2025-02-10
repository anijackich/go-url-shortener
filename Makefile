.PHONY: help tidy swag test build run

help: ## Display this help
	@echo "Usage: make [target] STORAGE=storage"
	@echo
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'
	@echo
	@echo "Storages:"
	@echo "  STORAGE=memory      In-memory storage (default)"
	@echo "  STORAGE=postgres    PostgreSQL storage"

STORAGE ?= memory

tidy: ## Run go mod tidy
	go mod tidy

fmt: ## Run formatter
	go fmt ./...

lint: ## Run linter
	golangci-lint run

test: ## Run tests & Write coverage report
	go test -cover ./... > coverage.out

swag: ## Init Swagger docs
	swag fmt -d .\internal\handlers\ -d .\cmd\
	swag init --parseDependency --parseInternal -d .\internal\handlers\ -g ..\..\cmd\main.go

build: ## Build the application
	go build -o go-url-shortener ./cmd/main.go

run: ## Run the application
	go run ./cmd/main.go --storage $(STORAGE)
