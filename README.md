# kafka-examples

Apache Kafka examples

You should create .env file with IP address to run kafka

example .env

```bash
IP_ADDR=192.168.0.180
```

You can get your ip address using command: `ifconfig | grep "inet " | grep -v 127.0.0.1 | cut -d\  -f2`

You can create `example` topic using this command:
`docker exec -t <KafkaContainerID> kafka-topics --create --topic example --partitions 4 --replication-factor 2 --if-not-exists --zookeeper 192.168.0.180:32181`

You can use kafkacat tool to use kafka CLI: `brew install kafkacat` `apt install kafkacat`

To listen to partiotions: `kafkacat -C -b localhost:19092,localhost:29092,localhost:39092 -t example -p 0`
