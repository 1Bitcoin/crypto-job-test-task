# Установка переменных
APP_NAME=crypto-job-test-task
DOCKER_IMAGE_NAME=rate_service_image

# Команды
.PHONY: build test docker-build run lint

# Сборка приложения
build:
	go build -o $(APP_NAME) ./cmd/service/main.go

# Запуск unit-тестов
test:
	go test ./... -v

# Сборка Docker-образа
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# Запуск приложения
run:
	go run ./cmd/service/main.go

# Запуск линтера
lint:
	golangci-lint run