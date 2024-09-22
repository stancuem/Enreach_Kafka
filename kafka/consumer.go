package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"Go_REST_Kafka/models"
	"github.com/segmentio/kafka-go"
)

func ConsumeMessages(brokers []string, topic string, msgChan chan *models.User) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     strconv.Itoa(rand.Intn(500)),
		StartOffset: kafka.FirstOffset,
	})

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var user models.User
		if err := json.Unmarshal(m.Value, &user); err == nil {
			fmt.Println(user.Id)
			msgChan <- &user
		}
	}
}
