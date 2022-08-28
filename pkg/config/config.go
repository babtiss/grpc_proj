package config

import (
	"fmt"
	"os"
)

type Config interface {
	GetSqlPath() string
}

type PostgresConfig struct {
	host         string
	port         string
	user         string
	password     string
	dbname       string
	sslmode      string
	AdditionalDB string
	DB           string
}

func (c *PostgresConfig) GetSqlPath() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		c.host, c.port, c.user, c.password, c.dbname, c.sslmode)
}

func NewConfig() (*PostgresConfig, error) {
	// можно за настройками ходить курлом в 'cекретницу', но я по-простому сделал
	return &PostgresConfig{
		host:         os.Getenv("host"),
		port:         os.Getenv("port"),
		user:         os.Getenv("user"),
		password:     os.Getenv("password"),
		dbname:       os.Getenv("dbname"),
		sslmode:      os.Getenv("sslmode"),
		AdditionalDB: "clickhouse",
		DB:           "postgres",
	}, nil
}
