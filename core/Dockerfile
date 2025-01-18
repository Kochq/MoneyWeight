FROM golang:1.21-alpine

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o moneyweight-api ./cmd/api/main.go

# Expose port
EXPOSE 8080

# Run the application
CMD ["./moneyweight-api"]
