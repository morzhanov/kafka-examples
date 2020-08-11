package main

import (
	"fmt"

	"github.com/morzhanov/kafka-examples/api"
	"github.com/morzhanov/kafka-examples/connection"
)

func main() {
	const topic = "example"
	const partition = 0

	// create kafka connection
	conn := connection.CreateKafkaConnection(topic, partition)
	fmt.Printf("Kafka connection local address: %v\n", conn.LocalAddr().String())

	// create api router
	api.CreateRouter(conn)
}
