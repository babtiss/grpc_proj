package repository

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"grpcProject/pkg/config"
)

func NewClickhouse(config config.Config) (*sql.DB, error) {
	connect, err := sql.Open("clickhouse", config.GetDsnAddDB())
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
	return connect, nil
}
