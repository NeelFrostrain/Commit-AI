# Commit-AI Makefile

.PHONY: build test clean install run lint fmt help

# Variables
BINARY_NAME=commit-ai
INSTALL_PATH=$(GOPATH)/bin
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

help: ## Show this help message
	@echo "Commit-AI - Development Commands"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_NAME).exe .

build-all: ## Build for all platforms
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe .
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 .
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 .

install: build ## Install the binary to GOPATH/bin
	@echo "Installing to $(INSTALL_PATH)..."
	cp $(BINARY_NAME).exe $(INSTALL_PATH)/

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run linter
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	golangci-lint run

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...
	gofmt -s -w .

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -f $(BINARY_NAME).exe
	rm -f install-$(BINARY_NAME).exe
	rm -rf dist/
	rm -f coverage.out coverage.html

run: build ## Build and run
	./$(BINARY_NAME).exe

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

update-deps: ## Update dependencies
	@echo "Updating dependencies..."
	go get -u ./...
	go mod tidy

dev: ## Run in development mode with verbose output
	go run . -v

check: fmt lint test ## Run all checks (format, lint, test)

release: clean test build-all ## Prepare release builds
	@echo "Release builds ready in dist/"
