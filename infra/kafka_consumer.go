package infra

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"notifications/application"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader  *kafka.Reader
	useCase *application.ProcessOrderUseCase
}

var _ application.Consumer = (*KafkaConsumer)(nil)

func NewKafkaConsumer(brokers []string, topic string, groupID string, uc *application.ProcessOrderUseCase) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return &KafkaConsumer{
		Reader:  reader,
		useCase: uc,
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

			var orderData struct {
				Id       string  `json:"id"`
				Total    float64 `json:"total"`
				Customer struct {
					Id    string `json:"id"`
					Name  string `json:"name"`
					Email string `json:"email"`
				} `json:"customer"`
			}
			if err := json.Unmarshal(msg.Value, &orderData); err != nil {
				log.Printf("Error unmarshaling Kafka message: %v", err)
				continue
			}

			if err = c.useCase.Execute(orderData.Id, orderData.Total, orderData.Customer.Id, orderData.Customer.Name, orderData.Customer.Email); err != nil {
				log.Printf("Error processing order: %v", err)
				continue
			}

			log.Printf("Order processed successfully. Order ID: %s", orderData.Id)
		}
	}
}
