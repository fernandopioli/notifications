package infra

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"notifications/application"
	"notifications/domain"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader         *kafka.Reader
	OrderProcessor *application.OrderProcessor
}

var _ application.Consumer = (*KafkaConsumer)(nil)

func NewKafkaConsumer(brokers []string, topic string, groupID string, processor *application.OrderProcessor) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return &KafkaConsumer{
		Reader:         reader,
		OrderProcessor: processor,
	}
}
func (c *KafkaConsumer) Consume(ctx context.Context) error {
	defer func() {
		log.Println("Closing Kafka consumer...")
		c.Reader.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("Context canceled")
			return nil
		default:
			msg, err := c.Reader.ReadMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return nil
				}
				log.Printf("Error reading message: %v", err)
				continue
			}

			log.Printf("Received message: %s", string(msg.Value))

			var order domain.Order
			if err = json.Unmarshal(msg.Value, &order); err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			if err = c.OrderProcessor.ProcessOrder(order); err != nil {
				log.Printf("Error processing order: %v", err)
				continue
			}

			log.Printf("Order processed successfully. Order ID: %s", order.ID)
		}
	}
}