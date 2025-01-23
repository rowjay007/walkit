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

.PHONY: run build test docs clean

run: docs
	air

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v ./...

docs:
	swag init -g ./cmd/server/main.go -d . -o docs --parseInternal --parseDependency

clean:
	rm -rf bin/
	rm -rf tmp/
	rm -rf docs/

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE)

lint:
	golangci-lint run

generate-swagger:
	swag fmt  # Format swagger comments
	swag init --parseDependency --parseInternal -g cmd/server/main.go -o docs