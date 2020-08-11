package connection

import (
	"context"
	"time"

	"github.com/morzhanov/kafka-examples/config"
	"github.com/segmentio/kafka-go"
)

// CreateKafkaConnection function creates connection to topic for Kafka
func CreateKafkaConnection(topic string, partition int) *kafka.Conn {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", config.KafkaConnectionURI, topic, partition)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	return conn
}
