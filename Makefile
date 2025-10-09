# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: help dev dev-api dev-worker build build-api build-worker clean install-air migrate-create migrate-up migrate-down migrate-force migrate-version

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

install-air: ## Install Air for hot-reloading
	@which air > /dev/null || (echo "Installing air..." && go install github.com/air-verse/air@latest && echo "Air installed to $(shell go env GOPATH)/bin/air")

dev-api: install-air ## Run API server with hot-reloading
	$(shell go env GOPATH)/bin/air -c .air.toml

dev-worker: install-air ## Run worker with hot-reloading
	$(shell go env GOPATH)/bin/air -c .air.worker.toml

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

migrate-create: ## Create a new migration file (usage: make migrate-create)
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir database/migrations -seq $$name

migrate-up: ## Run all pending migrations
	@migrate -path database/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" up

migrate-down: ## Rollback the last migration
	@migrate -path database/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" down 1

migrate-down-all: ## Rollback all migrations
	@migrate -path database/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" down -all

migrate-force: ## Force set migration version (usage: make migrate-force)
	@read -p "Enter version to force: " version; \
	migrate -path database/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" force $$version

migrate-version: ## Show current migration version
	@migrate -path database/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" version

migrate-drop: ## Drop everything in the database (BE CAREFUL!)
	@echo "WARNING: This will drop all tables in the database!"
	@read -p "Are you sure? (yes/no): " confirm; \
	if [ "$$confirm" = "yes" ]; then \
		migrate -path database/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" drop -f; \
	else \
		echo "Operation cancelled."; \
	fi
