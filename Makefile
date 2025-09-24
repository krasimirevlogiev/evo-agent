.PHONY: run build tidy clean

# Go parameters
BINARY_NAME=financial-agent.exe
CMD_PATH=./cmd/financial-agent

run:
	@go run $(CMD_PATH)

build:
	@echo "Building for production..."
	@go build -o $(BINARY_NAME) $(CMD_PATH)
	@echo "$(BINARY_NAME) built"

tidy:
	@go mod tidy

clean:
	@echo "Cleaning up..."
	@if exist $(BINARY_NAME) (del $(BINARY_NAME))
	@echo "Cleanup complete"
