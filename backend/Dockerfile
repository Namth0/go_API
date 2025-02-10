# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download

# Ajoutez ces variables d'environnement
ENV GIN_MODE=debug
ENV ALLOWED_ORIGINS="http://localhost:3000"

EXPOSE 8080
CMD ["go", "run", "cmd/api/main.go"]