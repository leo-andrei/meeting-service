# Variables
GO = go
GOFMT = gofmt
GOFILES = $(shell find . -name '*.go')
APP_NAME = meeting-service
DOCKER_IMAGE_NAME = meeting-service-image
DOCKER_CONTAINER_NAME = meeting-service-container

# Default target
.PHONY: all
all: build

# Build the Go application
.PHONY: build
build:
	$(GO) build -o $(APP_NAME) ./cmd/meeting-service/main.go

# Run the Go application directly
.PHONY: run
run:
	$(GO) run ./cmd/meeting-service/main.go

# Format Go code
.PHONY: fmt
fmt:
	$(GOFMT) -s -w $(GOFILES)

# Test Go application
.PHONY: test
test:
	$(GO) test ./...

# Build Docker image
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# Run Docker container
.PHONY: docker-run
docker-run:
	docker-compose up -d

# Stop Docker container
.PHONY: docker-stop
docker-stop:
	docker-compose down

# Clean Go build artifacts
.PHONY: clean
clean:
	rm -f $(APP_NAME)

# Rebuild the Docker image and run the container
.PHONY: docker-restart
docker-restart: docker-stop docker-build docker-run

# Show Docker logs
.PHONY: docker-logs
docker-logs:
	docker logs -f $(DOCKER_CONTAINER_NAME)

# Clean up Go build and Docker images
.PHONY: clean-all
clean-all: clean
	docker rmi $(DOCKER_IMAGE_NAME)

