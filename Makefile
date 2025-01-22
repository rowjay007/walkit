.PHONY: all build run dev test clean docker-build docker-run lint generate-swagger

# Variables
BINARY_NAME=walkit
DOCKER_IMAGE=walkit-api

all: build

# Build the application
build:
	go build -o bin/$(BINARY_NAME) ./cmd/server/

# Run the application
run:
	./bin/$(BINARY_NAME)

# Start development with live reloading
dev:
	air -c air.toml

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	go clean
	rm -rf bin/

# Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run Docker container
docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE)

# Run linter
lint:
	golangci-lint run

# Generate Swagger documentation
generate-swagger:
	swag init -g cmd/server/main.go