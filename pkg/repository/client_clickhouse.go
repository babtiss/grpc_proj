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
	stmt, err := tx.Prepare("INSERT INTO Client (name) values ($1)")
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
	return err
}

func (c *ClientClickhouse) Delete(entity ClientEntity) error { return nil }

func (c *ClientClickhouse) GetClients() ([]ClientEntity, error) { return nil, nil }

//if err != nil {
//	return fmt.Errorf("clickhouse begin error: %v", err)
//}
//_, err = begin.Exec("CREATE TABLE IF NOT EXISTS Client (Name String) ENGINE = Kafka('localhost:9092', 'test', 'test', 'JSONEachRow');",
//	logger.MyKafkaAddress, logger.Topic)
//if err != nil {
//	return fmt.Errorf("clickhouse create clientbase table error: %v", err)
//}
//err = begin.Commit()
//
//begin, err = c.db.Begin()
