# Variables
GO = go
MAIN = cmd/main.go
BIN = myapp
PORT = 8080

# Default target: build and run the application
.PHONY: run
run:
	$(GO) run $(MAIN)

# To Build the Go application
.PHONY: build
build:
	$(GO) build -o $(BIN) $(MAIN)

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BIN)

# Start the application with a specific port
.PHONY: start
start:
	@echo "Starting the application on port $(PORT)..."
	./$(BIN)

# Test the Go code
.PHONY: test
test:
	$(GO) test ./...
