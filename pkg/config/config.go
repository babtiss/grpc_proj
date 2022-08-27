package config

import (
	"fmt"
	"os"
)

type Config interface {
	GetSqlPath() string
}

type PostgresConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	Dbname       string
	SSLmode      string
	AdditionalDB string
	DB           string
}

func (c *PostgresConfig) GetSqlPath() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		c.Host, c.Port, c.User, c.Password, c.Dbname, c.SSLmode)
}

func NewConfig() (*PostgresConfig, error) {
	// можно за настройками ходить курлом в 'cекретницу', но я по-простому сделал
	return &PostgresConfig{
		Host:         os.Getenv("host"),
		Port:         os.Getenv("port"),
		User:         os.Getenv("user"),
		Password:     os.Getenv("password"),
		Dbname:       os.Getenv("dbname"),
		SSLmode:      os.Getenv("sslmode"),
		AdditionalDB: "postgres",
		DB:           "clickhouse",
	}, nil
}
