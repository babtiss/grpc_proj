package logger

import "github.com/segmentio/kafka-go"

type Client interface {
	AddLog(entity ClientEntity) error
}

type Logger struct {
	Client
}

func NewLogger(logger *kafka.Conn) *Logger {
	return &Logger{
		Client: NewClientKafka(logger),
	}
}
