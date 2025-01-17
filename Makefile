
.PHONY: all build run test clean docker-build docker-run

# Variables
BINARY_NAME=walkit
DOCKER_IMAGE=walkit-api

all: build

build:
	go build -o bin/$(BINARY_NAME) cmd/server/main.go

run:
	go run cmd/server/main.go

dev:
	air

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
	swag init -g cmd/server/main.go

