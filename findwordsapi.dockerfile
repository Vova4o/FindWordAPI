# Этап сборки
FROM golang:alpine as builder

WORKDIR /app

# Копируем исходный код в образ
COPY . .

# Собираем статически связанное приложение для уменьшения зависимостей
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o findwords .

# Этап запуска
FROM alpine:latest  

WORKDIR /app

# Копируем только необходимые файлы из этапа сборки
COPY --from=builder /app/findwords .
COPY templates /templates
COPY static /static
COPY russian_nouns.txt .

CMD ["/app/findwords"]