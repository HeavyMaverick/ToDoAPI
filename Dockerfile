#1) этап сборки (builder)
FROM golang:1.25-alpine AS builder

# Указываем рабочую директорию
WORKDIR /app

# Копируем go-модули и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем бинарник
RUN go build -o todoapp ./cmd/api

#2) итоговый образ
FROM alpine:latest

WORKDIR /app

# Копируем из builder собранный бинарник
COPY --from=builder /app/todoapp .
#COPY .env .

# Порт, который слушает API
EXPOSE 8080

# Команда запуска
CMD ["./todoapp"]
