APP_NAME=crypto-job-test-task
DOCKER_IMAGE_NAME=rate_service_image

.PHONY: build test docker-build run lint

build:
	go build -o $(APP_NAME) ./cmd/service/main.go

test:
	go test ./... -v

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

run:
	go run ./cmd/service/main.go

lint:
	golangci-lint run