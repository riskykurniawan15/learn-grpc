.PHONY: help generate tidy build server client clean

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

generate: ## Generate protobuf Go code
	@echo "Generating protobuf code..."
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user.proto
	@echo "Protobuf code generated!"

tidy: ## Run go mod tidy
	@echo "Running go mod tidy..."
	go mod tidy
	@echo "Dependencies updated!"

build: generate tidy ## Build the project
	@echo "Building project..."
	go build -o bin/server server/server.go
	go build -o bin/client client/client.go
	@echo "Build completed!"

server: generate tidy ## Run the gRPC server
	@echo "Starting gRPC server..."
	go run server/server.go

client: ## Run the gRPC client
	@echo "Starting gRPC client..."
	go run client/client.go

clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -f users.db
	@echo "Clean completed!"

install-deps: ## Install required Go dependencies
	@echo "Installing Go dependencies..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "Dependencies installed!"

setup: install-deps generate tidy ## Complete setup
	@echo "Setup completed! You can now run 'make server' to start the server."

