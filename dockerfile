# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

ENV CGO_ENABLED=1
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Stage 2: Create a minimal image
FROM debian:stable-slim

ENV PORT=8080

RUN useradd -m appuser

COPY --from=builder /app/main /usr/local/bin/main

RUN mkdir -p /app/data && chown appuser /app/data

ENV DBPATH=/app/data/database.db

EXPOSE 8080

USER appuser

ENTRYPOINT ["/usr/local/bin/main"]