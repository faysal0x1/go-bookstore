# Variables
APP_NAME=go-bookstore
BUILD_DIR=bin
ENTRY_POINT=cmd/app/main.go

# Commands
.PHONY: all build run test clean tidy

all: build

build: swagger
	@echo "Building..."
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(ENTRY_POINT)

swagger:
	@echo "Generating Swagger documentation..."
	@go run github.com/swaggo/swag/cmd/swag init -g $(ENTRY_POINT) --output docs

migrate-fresh: build
	@echo "Migrating fresh..."
	@./$(BUILD_DIR)/$(APP_NAME) --migrate=fresh

migrate-fresh-seed: build
	@echo "Migrating fresh with seed..."
	@./$(BUILD_DIR)/$(APP_NAME) --migrate=fresh --seed

run: build
	@echo "Running..."
	@./$(BUILD_DIR)/$(APP_NAME)

test:
	@echo "Testing..."
	@go test -v ./...

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

tidy:
	@echo "Tidying dependencies..."
	@go mod tidy
