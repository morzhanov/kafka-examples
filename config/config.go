package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVar function returns environment variable by key
func GetEnvVar(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var (
	// KafkaIPAddress contains Kafka IP address
	KafkaIPAddress = GetEnvVar("KAFKA_IP_ADDR")
	// KafkaPort contains Kafka IP address
	KafkaPort = GetEnvVar("KAFKA_PORT")
	// KafkaConnectionURI contains Kafka connection uri
	KafkaConnectionURI = KafkaIPAddress + ":" + KafkaPort
)
