#!/bin/bash
set -e

echo "Running Go linters..."

# Run `go fmt` to format the code
echo "Formatting code with go fmt..."
go fmt ./...

# Run `go vet` to check for potential issues
echo "Running go vet..."
go vet ./...

# Run `golangci-lint` if installed
if command -v golangci-lint &> /dev/null; then
    echo "Running golangci-lint..."
    golangci-lint run ./...
else
    echo "golangci-lint not installed. Skipping..."
fi

echo "✅ All lint checks passed!"