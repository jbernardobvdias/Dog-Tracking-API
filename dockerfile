# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Create a minimal image
FROM debian:bookworm-slim

# Set environment variables
ENV PORT=8080

# Create a non-root user (optional but recommended)
RUN useradd -m appuser

# Copy binary from builder
COPY --from=builder /app/main /usr/local/bin/main

# Expose port
EXPOSE 8080

# Switch to non-root user
USER appuser

# Command to run the executable
ENTRYPOINT ["/usr/local/bin/main"]