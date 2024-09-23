package main

import (
	"Go_REST_Kafka/enrich"
	"Go_REST_Kafka/internal/config"
	"Go_REST_Kafka/kafka"
	"Go_REST_Kafka/models"
	"Go_REST_Kafka/postgres"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()

	postgres.InitDB(cfg.PostgresDSN)

	msgChan := make(chan *models.User)

	go kafka.ConsumeMessages(cfg.KafkaBrokers, cfg.InputTopic, msgChan)

	for user := range msgChan {
		enrichedUser, err := enrich.EnrichUser(user)
		if err != nil {
			fmt.Println("Error enriching user:", err)
			continue
		}

		kafka.ProduceMessage(cfg.KafkaBrokers, cfg.OutputTopic, enrichedUser)
	}
	//brokers := []string{"localhost:9092"} 
	//topic := "InputTopic"                
	//
	//writer := kafka.NewWriter(kafka.WriterConfig{
	//	Brokers: brokers,
	//	Topic:   topic,
	//})
	//
	//users := []map[string]interface{}{
	//	{"id": 1, "first_name": "madara"},
	//	{"id": 2, "first_name": "valadimir"},
	//	{"id": 3, "first_name": "roma"},
	//}
	//
	//for _, user := range users {
	//	userData, err := json.Marshal(user)
	//	if err != nil {
	//		log.Fatalf("failed to marshal user data: %v", err)
	//	}
	//
	//	err = writer.WriteMessages(context.Background(),
	//		kafka.Message{
	//			Value: userData,
	//		},
	//	)
	//	if err != nil {
	//		log.Fatalf("failed to write message: %v", err)
	//	}
	//	log.Printf("Message written to Kafka: %s", userData)
	//}
	//
	//if err := writer.Close(); err != nil {
	//	log.Fatalf("failed to close writer: %v", err)
	//}
}
