# kafka-examples

Apache Kafka examples

## environment

You should create .env file with IP address to run kafka

example .env

```bash
IP_ADDR=192.168.0.180
```

You can get your ip address using command:

```bash
ifconfig | grep "inet " | grep -v 127.0.0.1 | cut -d\ -f2
```

## setup

To setup cluster run:

```bash
docker-compose up -d
```

This command will create 3 Zookeeper instances and 3 Kafka instances.

## management

To create `example` topic run:

```bash
docker exec -t <KafkaContainerID> \
    kafka-topics --create \
    --topic example \
    --partitions 4 \
    --replication-factor 2 \
    --if-not-exists \
    --zookeeper 192.168.0.180:32181
```

To list topics run:

```bash
docker exec -t <KafkaContainerID> kafka-topics --list --zookeeper 192.168.0.180:32181
```

You can use kafkacat tool to use kafka CLI:

```bash
brew install kafkacat  #macos

apt install kafkacat   #apt
```

To listen to partiotions run:

```bash
kafkacat -C -b localhost:19092,localhost:29092,localhost:39092 -t example -p 0
```

To list consumer groups:

```bash
kafka-consumer-groups --list --bootstrap-server 192.168.0.180:19092
```

## app

- Application uses golang gin for HTTP server creation and routing
- default port: `8080`
- default topic: `example`
- default partition: `0`
- endpoints

  - `GET /messages` - get all messages for default consumer group `example-consumer-group`
  - `GET /messages/:partition/:offset` - get message on desired partition and offset
  - `POST /messages` - create message

    - body:

    ```json
    {
      "Key": "key",
      "Value": "value"
    }
    ```
