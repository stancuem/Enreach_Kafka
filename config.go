package main

import (
	"os"
)

type Config struct {
	KafkaBroker string
	KafkaTopic  string
	DatabaseURL string
}

func LoadConfig() Config {
	return Config{
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		KafkaTopic:  os.Getenv("KAFKA_TOPIC"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
