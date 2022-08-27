package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type ClientEntity struct {
	Name string
}

type Client interface {
	Add(ClientEntity) error
	Delete(ClientEntity) error
	GetClients() ([]ClientEntity, error)
}

type Repository struct {
	Client
}

func NewRepository(db *sql.DB, typeDB string) *Repository {
	var rep *Repository
	switch typeDB {
	case "postgres":
		rep = &Repository{
			Client: NewClientPostgres(db),
		}
	case "clickhouse":
		rep = &Repository{
			Client: NewClientClickhouse(db),
		}
	}
	return rep
}
