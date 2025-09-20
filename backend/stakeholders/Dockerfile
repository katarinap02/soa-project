FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["go", "run", "main.go"]
