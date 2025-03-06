# Notification Microservice (Go)

This microservice is responsible for handling notifications based on events from Apache Kafka, particularly focused on listening to the order creation topic to send simulated email notifications.

## Features
- Listens to Kafka topic (`orders`) for new order events.
- Simulates sending emails by logging event details (name, email, and total) to the console.

## Kafka Integration
- Currently listens to the `orders` Kafka topic. (Topic creation is manual.)

## Build and Run
```shell
go build -o bin/notifications ./cmd/main.go
./bin/notifications
```

### Run locally
```shell
go run ./cmd/main.go 
```
### Tests
```shell
go test ./...

```

### Build the Docker image
```shell
docker build -t notifications-service .
```

### Run the container
```shell
docker run -e KAFKA_BROKERS=your-kafka-host:port notifications-service
```
