# Variables
APP_NAME = strategy-pattern
BIN_DIR = bin
BUILD_TARGET = $(BIN_DIR)/$(APP_NAME)
MAIN_FILE = ./cmd/main.go
TEST_DIR = ./internal/test
GO_CMD = go

# Default target
.DEFAULT_GOAL := help

# Build the application
build: $(BUILD_TARGET) ## Build the application binary
$(BUILD_TARGET): $(MAIN_FILE)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	$(GO_CMD) build -o $(BUILD_TARGET) $(MAIN_FILE)
	@echo "Build completed: $(BUILD_TARGET)"

# Run the application
run: build ## Build and run the application
	@echo "Running $(APP_NAME)..."
	./$(BUILD_TARGET)

# Run all tests
test: ## Run all unit tests
	@echo "Running tests..."
	$(GO_CMD) test ./...

# Run specific tests
test-internal: ## Run tests located in internal/test
	@echo "Running internal tests..."
	$(GO_CMD) test $(TEST_DIR)/...

# Clean build artifacts
clean: ## Remove build artifacts
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)
	@echo "Cleanup completed!"

# Format code
fmt: ## Format Go code
	@echo "Formatting code..."
	$(GO_CMD) fmt ./...

# Display help
help: ## Display this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'
