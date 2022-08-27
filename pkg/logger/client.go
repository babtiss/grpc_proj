package logger

import (
	"fmt"
	"github.com/segmentio/kafka-go"
)

type ClientKafka struct {
	conn *kafka.Conn
}

type ClientEntity struct {
	Name string
}

func NewClientKafka(logger *kafka.Conn) *ClientKafka {
	return &ClientKafka{conn: logger}
}

func (c *ClientKafka) AddLog(entity ClientEntity) error {
	message := fmt.Sprintf("Add new client with name %v", entity.Name)
	_, err := c.conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)
	if err != nil {
		return err
	}
	return nil
}
