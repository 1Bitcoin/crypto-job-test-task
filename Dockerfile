# Используем официальный образ Golang как базовый
FROM golang:1.23.1 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код проекта
COPY . .

# Собираем исполняемый файл и называем его 'app'
RUN go build -o app ./cmd/service/main.go

# Используем более легкий образ для запуска приложения
FROM gcr.io/distroless/base

# Копируем скомпилированный исполняемый файл из первого этапа
COPY --from=builder /app/app .

# Указываем команду для запуска
CMD ["./app"]