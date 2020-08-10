package main

import (
	"fmt"

	"github.com/morzhanov/kafka-examples/connection"
)

func main() {
	// create kafka connection
	conn := connection.CreateKafkaConnection()
	fmt.Printf("%v\n", conn)
}
