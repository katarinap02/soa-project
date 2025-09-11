FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 8084
ENTRYPOINT ["go", "run", "main.go"]