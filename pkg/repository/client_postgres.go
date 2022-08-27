package repository

import (
	"database/sql"
	"fmt"
)

type ClientPostgres struct {
	db *sql.DB
}

func NewClientPostgres(db *sql.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (c *ClientPostgres) Add(entity ClientEntity) error {
	_, err := c.db.Exec("insert into Client (name) values ($1)", entity.Name)
	return err
}

func (c *ClientPostgres) Delete(entity ClientEntity) error {
	_, err := c.db.Exec("delete from Client where name = ($1)", entity.Name)
	return err
}

func (c *ClientPostgres) GetClients() ([]ClientEntity, error) {
	var clients []ClientEntity
	rows, err := c.db.Query("select * from Client")
	if err != nil {
		return nil, fmt.Errorf("postgres can't make request: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		client := ClientEntity{}
		err = rows.Scan(&client.Name)
		if err != nil {
			return nil, fmt.Errorf("trouble scan data to client entity: %v", err)
		}
		clients = append(clients, client)
	}
	return clients, nil
}
