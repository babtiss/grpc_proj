package repository

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
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
	return connect, nil
}
