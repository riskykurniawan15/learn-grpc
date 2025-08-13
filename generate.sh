#!/bin/bash

# Generate Go code from protobuf
echo "Generating Go code from protobuf..."

# Create proto directory if it doesn't exist
mkdir -p proto

# Generate Go code
protoc --go_out=. \
       --go_opt=paths=source_relative \
       --go-grpc_out=. \
       --go-grpc_opt=paths=source_relative \
       proto/user.proto

echo "Protobuf code generated successfully!"
echo "Now you can run: go mod tidy"

