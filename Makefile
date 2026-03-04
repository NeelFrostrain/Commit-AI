# Commit-AI Makefile

.PHONY: build build-prod build-all test clean install run lint fmt help installer version

# Variables
BINARY_NAME=commit-ai
BIN_DIR=bin
DIST_DIR=dist
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "v1.2.0")
BUILD_DATE=$(shell date -u '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LDFLAGS=-ldflags "-X 'main.Version=$(VERSION)' -X 'main.BuildDate=$(BUILD_DATE)' -X 'main.GitCommit=$(GIT_COMMIT)' -s -w"
GOFLAGS=-trimpath

help: ## Show this help message
	@echo "Commit-AI - Development Commands"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the binary for development
	@echo "Building $(BINARY_NAME) for development..."
	@go build $(LDFLAGS) -o $(BINARY_NAME).exe .
	@echo "✓ Built: $(BINARY_NAME).exe"

build-prod: ## Build production binary to bin/
	@echo "Building $(BINARY_NAME) for production..."
	@echo "  Version: $(VERSION)"
	@mkdir -p $(BIN_DIR)
	@go build $(GOFLAGS) $(LDFLAGS) -o $(BIN_DIR)/$(BINARY_NAME).exe .
	@echo "[OK] Built: $(BIN_DIR)/$(BINARY_NAME).exe"

installer: ## Build the installer
	@echo "Building installer..."
	@echo "  Version: $(VERSION)"
	@mkdir -p $(BIN_DIR)
	@cd installer && go build $(GOFLAGS) $(LDFLAGS) -o ../$(BIN_DIR)/install-$(BINARY_NAME).exe .
	@echo "[OK] Built: $(BIN_DIR)/install-$(BINARY_NAME).exe"

build-all: ## Build for all platforms
	@echo "Building for all platforms..."
	@echo "  Version: $(VERSION)"
	@mkdir -p $(DIST_DIR)
	@echo "  → Windows (amd64)..."
	@GOOS=windows GOARCH=amd64 go build $(GOFLAGS) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe .
	@echo "  → Linux (amd64)..."
	@GOOS=linux GOARCH=amd64 go build $(GOFLAGS) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64 .
	@echo "  → Linux (arm64)..."
	@GOOS=linux GOARCH=arm64 go build $(GOFLAGS) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm64 .
	@echo "  → macOS (amd64)..."
	@GOOS=darwin GOARCH=amd64 go build $(GOFLAGS) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64 .
	@echo "  → macOS (arm64)..."
	@GOOS=darwin GOARCH=arm64 go build $(GOFLAGS) $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64 .
	@echo "[OK] All builds complete in $(DIST_DIR)/"

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✓ Coverage report: coverage.html"

lint: ## Run linter
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Install: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	@golangci-lint run

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@gofmt -s -w .
	@echo "✓ Code formatted"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME).exe
	@rm -rf $(DIST_DIR)
	@rm -f coverage.out coverage.html
	@echo "[OK] Cleaned"

clean-all: clean ## Clean everything including bin/
	@rm -rf $(BIN_DIR)
	@echo "[OK] Cleaned all"

run: build ## Build and run
	@./$(BINARY_NAME).exe

run-verbose: build ## Build and run with verbose output
	@./$(BINARY_NAME).exe -v

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "✓ Dependencies ready"

update-deps: ## Update dependencies
	@echo "Updating dependencies..."
	@go get -u ./...
	@go mod tidy
	@echo "✓ Dependencies updated"

check: fmt test ## Run format and tests
	@echo "✓ All checks passed"

release: clean test build-prod installer build-all ## Prepare full release
	@echo ""
	@echo "[OK] Release ready!"
	@echo "  Version:    $(VERSION)"
	@echo "  Production: $(BIN_DIR)/"
	@echo "  Installer:  $(BIN_DIR)/install-$(BINARY_NAME).exe"
	@echo "  Platforms:  $(DIST_DIR)/"

install-local: build-prod ## Install to local bin
	@echo "Installing locally..."
	@cp $(BIN_DIR)/$(BINARY_NAME).exe $(GOPATH)/bin/ 2>/dev/null || cp $(BIN_DIR)/$(BINARY_NAME).exe ~/go/bin/
	@echo "✓ Installed to Go bin"

.DEFAULT_GOAL := help
