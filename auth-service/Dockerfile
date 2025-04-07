FROM golang:1.21-alpine

WORKDIR /app

# Copy go.mod and go.sum files first for better caching
COPY go.mod go.sum* ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the application
RUN go build -o auth-service .

# Expose the port
EXPOSE 8081

# Run the service
CMD ["./auth-service"]
