APP_NAME := go-social-project
BINARY_NAME := $(APP_NAME)

dev:
	@echo "Running dev server..."
	air

test:
	@echo "Running tests..."
	go test -v ./...

lint:
	@echo "Running linter..."
	golangci-lint run

tidy:
	@echo "Tidying module dependencies..."
	go mod tidy