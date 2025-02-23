package config

import (
	"os"
	"strings"
)

type AppConfig struct {
	Brokers  []string
	Topic    string
	GroupID  string
}

func LoadConfig() (*AppConfig, error) {
	brokersEnv := os.Getenv("KAFKA_BROKERS")
	if brokersEnv == "" {
		brokersEnv = "localhost:19092"
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "orders"
	}

	groupID := os.Getenv("KAFKA_GROUP_ID")
	if groupID == "" {
		groupID = "notifications-group"
	}

	return &AppConfig{
		Brokers: strings.Split(brokersEnv, ","),
		Topic:   topic,
		GroupID: groupID,
	}, nil
} 