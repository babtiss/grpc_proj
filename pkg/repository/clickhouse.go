package repository

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"grpcProject/pkg/logger"
)

const ClickHouseTCPConnect = "tcp://127.0.0.1:9000?debug=true"

func NewClickhouse() (*sql.DB, error) {
	connect, err := sql.Open("clickhouse", ClickHouseTCPConnect)
	if err != nil {
		return nil, fmt.Errorf("open clickhouse error: %v", err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return nil, fmt.Errorf("clickhouse database is not responding: %v", err)
	}
	begin, err := connect.Begin()
	if err != nil {
		fmt.Printf("begin error: %v", err)
		return nil, fmt.Errorf("clickhouse begin error: %v", err)
	}
	_, err = begin.Exec("CREATE TABLE IF NOT EXISTS clientbase (Name String) ENGINE = Kafka($1, $2, 'test', 'JSONEachRow');",
		logger.MyKafkaAddress, logger.Topic)
	if err != nil {
		return nil, fmt.Errorf("clickhouse create clientbase table error: %v", err)
	}
	err = begin.Commit()
	if err != nil {
		fmt.Printf("comit error: %v", err)
		return nil, fmt.Errorf("clickhouse commit error: %v", err)
	}
	return connect, nil
}
