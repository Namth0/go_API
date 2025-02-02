# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 8080
CMD ["go", "run", "cmd/api/main.go"]