.PHONY: build install uninstall clean test run

# Build the binary
build:
	@echo "Building focus..."
	@go build -o focus

# Install to /usr/local/bin
install: build
	@echo "Installing focus to /usr/local/bin..."
	@sudo mv focus /usr/local/bin/
	@echo "✓ Installation complete! Run 'focus --help' to get started."

# Uninstall from /usr/local/bin
uninstall:
	@echo "Uninstalling focus..."
	@sudo rm -f /usr/local/bin/focus
	@echo "✓ Uninstall complete."

# Remove binary from current directory
clean:
	@echo "Cleaning..."
	@rm -f focus
	@echo "✓ Clean complete."

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Run without installing
run: build
	@./focus

# Development build (with race detector)
dev:
	@go build -race -o focus