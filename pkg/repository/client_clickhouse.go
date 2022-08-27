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
	begin, err := c.db.Begin()
	if err != nil {
		fmt.Printf("begin error: %v", err)
		return fmt.Errorf("clickhouse begin error: %v", err)
	}
	_, err = begin.Exec("insert into queue (Name) values ($1)", entity.Name)
	if err != nil {
		return fmt.Errorf("clickhouse add data error: %v", err)
	}
	err = begin.Commit()
	if err != nil {
		fmt.Printf("comit error: %v", err)
		return fmt.Errorf("clickhouse commit error: %v", err)
	}
	return err
}

func (c *ClientClickhouse) Delete(entity ClientEntity) error { return nil }

func (c *ClientClickhouse) GetClients() ([]ClientEntity, error) { return nil, nil }
