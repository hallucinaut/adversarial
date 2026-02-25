.PHONY: all build test clean lint vet fmt fmt-check vet-check coverage build-linux build-darwin build-windows

# Variables
APP_NAME := adversarial
VERSION := 1.0.0
GO := go
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
BINARY := $(APP_NAME)
OUTPUT_DIR := bin

# Go flags
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

# Default target
all: lint test build

# Build the application
build: $(OUTPUT_DIR)/$(BINARY)

$(OUTPUT_DIR)/$(BINARY):
	@echo "Building $(BINARY) $(VERSION) for $(GOOS)/$(GOARCH)..."
	@mkdir -p $(OUTPUT_DIR)
	$(GO) build -o $(OUTPUT_DIR)/$(BINARY) $(LDFLAGS) ./cmd/$(APP_NAME)

# Cross-compilation targets
build-linux:
	@echo "Building $(BINARY) for Linux..."
	GOOS=linux GOARCH=amd64 $(GO) build -o $(OUTPUT_DIR)/$(BINARY)_linux_amd64 ./cmd/$(APP_NAME)
	GOOS=linux GOARCH=arm64 $(GO) build -o $(OUTPUT_DIR)/$(BINARY)_linux_arm64 ./cmd/$(APP_NAME)

build-darwin:
	@echo "Building $(BINARY) for macOS..."
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(OUTPUT_DIR)/$(BINARY)_darwin_amd64 ./cmd/$(APP_NAME)
	GOOS=darwin GOARCH=arm64 $(GO) build -o $(OUTPUT_DIR)/$(BINARY)_darwin_arm64 ./cmd/$(APP_NAME)

build-windows:
	@echo "Building $(BINARY) for Windows..."
	GOOS=windows GOARCH=amd64 $(GO) build -o $(OUTPUT_DIR)/$(BINARY)_windows_amd64.exe ./cmd/$(APP_NAME)

# Test the application
test:
	@echo "Running tests..."
	$(GO) test -v -race -coverprofile=coverage.out ./...

# Run tests with coverage report
coverage: test
	@echo "Generating coverage report..."
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Run specific tests
test-run:
	$(GO) test -v -run $(TEST_PATTERN) ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(OUTPUT_DIR)
	rm -f coverage.out coverage.html

# Lint the code
lint:
	@echo "Running linter..."
	@golangci-lint run ./...

# Run go vet
vet:
	@echo "Running go vet..."
	$(GO) vet ./...

# Format the code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...
	$(GO) imports -w ./...

# Check if code is formatted
fmt-check:
	@echo "Checking code formatting..."
	@diff=$$($(GO) diff -r .); \
	if [ -n "$$diff" ]; then \
		echo "Code is not formatted. Run 'make fmt' to fix."; \
		echo "$$diff"; \
		exit 1; \
	fi

# Security scan
security:
	@echo "Running security scan..."
	@gosec -quiet ./...

# Integration tests
integration:
	@echo "Running integration tests..."
	$(GO) test -tags=integration -v ./...

# Install dependencies
install:
	@echo "Installing dependencies..."
	$(GO) mod download
	$(GO) install ./cmd/$(APP_NAME)

# Generate mocks (if using mocking)
mocks:
	@echo "Generating mocks..."
	mockgen -source=pkg/detect/detect.go -destination=pkg/detect/mock_detect.go -package=detect

# Help command
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build        - Build the application"
	@echo "  test         - Run all tests"
	@echo "  coverage     - Run tests with coverage report"
	@echo "  lint         - Run linter"
	@echo "  vet          - Run go vet"
	@echo "  fmt          - Format code"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install the application"
	@echo "  security     - Run security scan"
	@echo "  help         - Show this help message"
	@echo ""
	@echo "Cross-compilation:"
	@echo "  build-linux    - Build for Linux"
	@echo "  build-darwin   - Build for macOS"
	@echo "  build-windows  - Build for Windows"