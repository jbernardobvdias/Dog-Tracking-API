# Dog Tracking API

A simple and scalable RESTful API for tracking dogs, built with Go (Golang) and containerized using Docker.

## Features

- Add, update, delete, and retrieve dog records
- Dockerized for easy deployment

## Technologies

- Go (Golang)
- Docker
- net/http (standard go library)
- gin (third party go library)

## Getting Started

### Prerequisites

- [Golang](https://go.dev)
- [Docker](https://www.docker.com)

### Clone the Repository

```bash
git clone https://github.com/jbernardobvdias/dog-tracking-api.git
cd dog-tracking-api
```

### Running locally

```bash
go run main.go
```

### Running on Docker

You can use this with regular docker.

```bash
docker build -t dogs .
docker run -p 8080:8080 dogs
```

Or you can use it with docker compose.

```bash
docker-compose up --build
```