FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 8082
ENTRYPOINT ["go", "run", "main.go"]