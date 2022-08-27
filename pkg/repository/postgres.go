package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"grpcProject/pkg/config"
)

func NewPostgresDataBase(config config.Config) (*sql.DB, error) {
	connStr := config.GetSqlPath()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("database is not responding: %v", err)
	}

	return db, nil
}
