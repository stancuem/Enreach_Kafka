package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Подключение к Kafka
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "users",
		GroupID: "data_enrichment_group",
	})

	// Подключение к PostgreSQL
	db, err := ConnectDatabase(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		// Десериализация входящих данных
		var user User
		if err := json.Unmarshal(m.Value, &user); err != nil {
			log.Printf("Не удалось десериализовать сообщение: %v", err)
			continue
		}

		// Обогащение данных через REST API
		enrichedUser := enricData(user)
		db.Save(enrichedUser)

		fmt.Printf("Сохранен пользователь: %+v\n", enrichedUser)
	}

	if err := r.Close(); err != nil {
		log.Fatal("Не удалось закрыть ридер:", err)
	}
}
