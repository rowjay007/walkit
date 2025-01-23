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

test:
	go test -v ./...

clean:
	go clean
	rm -rf bin/

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE)

lint:
	golangci-lint run

generate-swagger:
	swag fmt  # Format swagger comments
	swag init --parseDependency --parseInternal -g cmd/server/main.go -o docs