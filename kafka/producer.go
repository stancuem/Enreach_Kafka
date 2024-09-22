package kafka

import (
	"context"
	"encoding/json"
	"log"

	"Go_REST_Kafka/models"
	"github.com/segmentio/kafka-go"
)

func ProduceMessage(brokers []string, topic string, user *models.User) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	userData, _ := json.Marshal(user)
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: userData,
		},
	)
	if err != nil {
		log.Fatal("Failed to write message:", err)
	}
}
