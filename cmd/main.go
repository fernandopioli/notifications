package main

import (
	"context"
	"log"
	"notifications/application"
	"notifications/config"
	"notifications/infra"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println("Microservice starting...")

	emailSender := infra.NewLoggerEmailSender()
	uc := application.NewProcessOrderUseCase(emailSender)

	kafkaConsumer := infra.NewKafkaConsumer(
		cfg.Brokers,
		cfg.Topic,
		cfg.GroupID,
		uc,
	)

	log.Println("Kafka consumer created. About to consume messages.")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		if err := kafkaConsumer.Consume(ctx); err != nil {
			log.Fatalf("Error starting consumer: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan

	log.Printf("Received signal: %v. Shutting down gracefully...", sig)
	cancel()
	time.Sleep(2 * time.Second)
}
