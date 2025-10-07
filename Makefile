.PHONY: help dev dev-api dev-worker build build-api build-worker clean install-air

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

install-air: ## Install Air for hot-reloading
	@which air > /dev/null || (echo "Installing air..." && go install github.com/air-verse/air@latest)

dev-api: install-air ## Run API server with hot-reloading
	air -c .air.toml

dev-worker: install-air ## Run worker with hot-reloading
	air -c .air.worker.toml

dev: ## Run both API and worker with hot-reloading (in parallel)
	@echo "Starting API and Worker in development mode..."
	@make -j2 dev-api dev-worker

build-api: ## Build API server binary
	@echo "Building API server..."
	@go build -o bin/api ./cmd/api

build-worker: ## Build worker binary
	@echo "Building worker..."
	@go build -o bin/worker ./cmd/worker

build: ## Build all binaries
	@echo "Building all binaries..."
	@mkdir -p bin
	@make build-api
	@make build-worker
	@echo "Build complete!"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin tmp
	@echo "Clean complete!"

run-api: build-api ## Build and run API server
	@./bin/api

run-worker: build-worker ## Build and run worker
	@./bin/worker
