.PHONY: run
run:
	@echo "Running the program..."
	@go run cmd/main.go

.PHONY: build
build:
	@echo "Building the program..."
	@go build -o bin/main cmd/main.go


.PHONY: install
install:
	@echo "Installing the program..."
	@go mod download
	@go install github.com/a-h/templ/cmd/templ@latest

.PHONY: generate-templ
generate-templ:
	@echo "Generating the templ..."
	@templ generate static/templates