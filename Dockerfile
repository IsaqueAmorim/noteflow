FROM golang:1.24.1-alpine AS builder

# Set environment variables for Go
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install necessary dependencies
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first for dependency caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o /app/bin/noteflow ./cmd/main.go

# Use a minimal base image for production
FROM alpine:latest

# Set environment variables for the application
ENV APP_ENV=production \
    PORT=8080

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/noteflow /app/noteflow

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["/app/noteflow"]