FROM golang:alpine AS builder
WORKDIR /app
COPY . .
EXPOSE 8080
EXPOSE 9091
ENTRYPOINT ["go", "run", "main.go"]
