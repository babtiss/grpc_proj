package repository

import (
	"database/sql"
	"fmt"
)

type ClientClickhouse struct {
	db *sql.DB
}

func NewClientClickhouse(db *sql.DB) *ClientClickhouse {
	return &ClientClickhouse{db: db}
}

func (c *ClientClickhouse) Add(entity ClientEntity) error {
	tx, err := c.db.Begin()
	if err != nil {
		return fmt.Errorf("clickhouse begin error: %v", err)
	}
	stmt, err := tx.Prepare("INSERT INTO Client (Name) values ($1)")
	if err != nil {
		return fmt.Errorf("clickhouse prepare insert error: %v", err)
	}
	_, err = stmt.Exec(entity.Name)
	if err != nil {
		return fmt.Errorf("clickhouse exec insert error: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("clickhouse commit error: %v", err)
	}
	if err := c.db.Ping(); err != nil {
		return err
	}
	return nil
}

func (c *ClientClickhouse) Delete(entity ClientEntity) error { return nil }

func (c *ClientClickhouse) GetClients() ([]ClientEntity, error) { return nil, nil }
