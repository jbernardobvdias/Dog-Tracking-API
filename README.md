# Dog Tracking API

A simple and scalable RESTful API for tracking dogs, built with Go (Golang) and containerized using Docker.

## Features

- Add, update, delete, and retrieve dog records
- RESTful API with JSON responses
- Dockerized for easy deployment

## Technologies

- Go (Golang)
- Docker
- net/http (standard go library)
- gin (third party go library)
- JSON for API communication

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started) installed

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

```bash
docker build -t dogs .
docker run dogs
```