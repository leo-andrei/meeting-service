# Stage 1: Build
FROM golang:1.19 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o meeting-service ./cmd/meeting-service

# Stage 2: Run
FROM alpine:latest

# Add certificates for SSL
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the built binary from the builder
COPY --from=builder /app/meeting-service .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./meeting-service"]
