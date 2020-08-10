package connection

import "github.com/morzhanov/kafka-examples/config"

var (
	// KafkaIPAddress contains Kafka IP address
	KafkaIPAddress = config.GetEnvVar("KAFKA_IP_ADDR")
	// KafkaPort contains Kafka IP address
	KafkaPort = config.GetEnvVar("KAFKA_PORT")
	// KafkaConnectionURI contains Kafka connection uri
	KafkaConnectionURI = KafkaIPAddress + ":" + KafkaPort
)
