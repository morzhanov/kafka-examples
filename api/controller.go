package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morzhanov/kafka-examples/config"
	"github.com/segmentio/kafka-go"
)

// CreateMessageBody struct
type CreateMessageBody struct {
	Key   string
	Value string
}

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Alive",
	})
}

func getMessages(c *gin.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{config.KafkaConnectionURI},
		GroupID: "example-consumer-group",
		Topic:   "example",
		MaxWait: 1 * time.Second,
	})
	defer r.Close()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 7*time.Second)
	defer cancel()

	m, err := r.ReadMessage(ctx)

	if err != nil && err.Error() != "context deadline exceeded" {
		fmt.Printf("\ngetMessages error: %v\n", err)
		c.String(http.StatusNotFound, err.Error())
		return
	}
	if m.Value == nil {
		fmt.Printf("\nNo new messages\n")
		c.String(http.StatusNotFound, "[EMPTY]: No new messages")
		return
	}

	fmt.Printf("\nmessage at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	c.String(http.StatusOK, string(m.Value))
}

func getMessageByOffset(c *gin.Context) {
	offsetParam := c.Param("offset")
	partitionParam := c.Param("partition")
	partition, err := strconv.Atoi(partitionParam)
	if err != nil {
		message := fmt.Sprintf("partition param should be valid integer, provided is: %v", partitionParam)
		fmt.Printf("\n%v\n", message)
		c.String(http.StatusBadRequest, message)
		return
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.KafkaConnectionURI},
		Topic:     "example",
		Partition: partition,
		MaxWait:   1 * time.Second,
	})
	defer r.Close()

	offset, err := strconv.ParseInt(offsetParam, 10, 64)
	if err != nil {
		message := fmt.Sprintf("offset param should be valid integer, provided is: %v", offsetParam)
		fmt.Printf("\n%v\n", message)
		c.String(http.StatusBadRequest, message)
		return
	}
	r.SetOffset(offset)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	m, err := r.ReadMessage(ctx)

	if err != nil && err.Error() != "context deadline exceeded" {
		fmt.Printf("\ngetMessages error: %v\n", err)
		c.String(http.StatusNotFound, err.Error())
		return
	}
	if m.Value == nil {
		message := fmt.Sprintf("No message found on the %v offset", offset)
		fmt.Printf("\n%v\n", message)
		c.String(http.StatusNotFound, message)
		return
	}

	fmt.Printf("\nmessage at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	c.String(http.StatusOK, string(m.Value))
}

func createMessage(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var data CreateMessageBody
	json.Unmarshal(body, &data)
	fmt.Printf("createMessage body is: %s \n", data)

	if data.Key == "" || data.Value == "" {
		message := "Key and Value body params are required"
		fmt.Printf("\n%v\n", message)
		c.String(http.StatusBadRequest, message)
		return
	}

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{config.KafkaConnectionURI},
		Topic:    "example",
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(data.Key),
			Value: []byte(data.Value),
		},
	)
}

// CreateRouter function creates main app router
func CreateRouter(kafkaConn *kafka.Conn) {
	r := gin.Default()
	r.GET("/health", health)
	r.GET("/messages", getMessages)
	r.GET("/messages/:partition/:offset", getMessageByOffset)
	r.POST("/messages", createMessage)
	r.Run()
}
