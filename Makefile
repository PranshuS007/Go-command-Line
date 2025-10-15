
.PHONY: build clean test install run help

# Binary name
BINARY_NAME=filer

# Build variables
BUILD_DIR=build
VERSION=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the binary
	@echo "Building ${BINARY_NAME}..."
	@go build ${LDFLAGS} -o ${BINARY_NAME} .
	@echo "Build complete: ${BINARY_NAME}"

build-all: ## Build for all platforms
	@echo "Building for multiple platforms..."
	@mkdir -p ${BUILD_DIR}
	@GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME}-linux-amd64 .
	@GOOS=linux GOARCH=arm64 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME}-linux-arm64 .
	@GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME}-darwin-amd64 .
	@GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME}-darwin-arm64 .
	@GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${BINARY_NAME}-windows-amd64.exe .
	@echo "Multi-platform build complete"

clean: ## Remove build artifacts
	@echo "Cleaning..."
	@rm -f ${BINARY_NAME}
	@rm -rf ${BUILD_DIR}
	@rm -f coverage.txt coverage.html
	@echo "Clean complete"

test: ## Run tests
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	@echo "Tests complete"

test-coverage: test ## Run tests with coverage report
	@go tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage report generated: coverage.html"

install: build ## Install the binary to GOPATH/bin
	@echo "Installing ${BINARY_NAME}..."
	@go install ${LDFLAGS}
	@echo "Installation complete"

run: build ## Build and run the binary
	@./${BINARY_NAME}

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete"

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...
	@echo "Vet complete"

lint: ## Run linters
	@echo "Running linters..."
	@golangci-lint run
	@echo "Lint complete"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod verify
	@echo "Dependencies downloaded"

tidy: ## Tidy go.mod
	@echo "Tidying go.mod..."
	@go mod tidy
	@echo "Tidy complete"

demo: build ## Run the demo script
	@echo "Running demo..."
	@chmod +x demo.sh
	@./demo.sh

.DEFAULT_GOAL := help
