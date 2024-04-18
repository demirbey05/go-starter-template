


.PHONY: run
run:
	@echo "Running the program..."
	@go run cmd/main.go

.PHONY: build
build:
	@echo "Building the program..."
	@go build -o bin/main cmd/main.go