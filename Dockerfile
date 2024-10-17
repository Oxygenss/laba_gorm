FROM golang:latest

WORKDIR /app

# Копируем зависимости
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# Копируем исходный код
COPY ./ ./

# Собираем приложение
RUN go build -o ./bin/app cmd/app/main.go

# Указываем, что нужно запускать
CMD [ "./bin/app" ]