package connection

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

const topic = "my-topic"
const partition = 0

// CreateKafkaConnection function creates connection to topic for Kafka
func CreateKafkaConnection() *kafka.Conn {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	return conn
}
