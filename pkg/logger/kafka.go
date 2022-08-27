package logger

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

const (
	Topic          = "test"
	MyKafkaAddress = "localhost:9092"
)

func NewKafka() (*kafka.Conn, error) {
	p, err := kafka.DefaultDialer.LookupPartition(context.TODO(), "tcp", MyKafkaAddress, Topic, 0)
	if err != nil {
		return nil, fmt.Errorf("can not lookup partition: %v", err)
	}
	//I had problems assigning a host, so I change it manually
	p.Leader.Host = "localhost"
	conn, err := kafka.DialPartition(context.Background(), "tcp", MyKafkaAddress, p)
	if err != nil {
		return nil, fmt.Errorf("partition error: %v", err)
	}
	err = conn.SetWriteDeadline(time.Now().Add(100 * time.Second))
	if err != nil {
		return nil, fmt.Errorf("set deadline error: %v", err)
	}

	return conn, nil
}
