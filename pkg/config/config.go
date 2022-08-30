package config

import (
	"fmt"
)

const (
	configPath = "configs/config.yml"
)

type Config interface {
	GetSqlPath() string
}

type MainConfig struct {
	AdditionalDB AdditionalDB
	DB           DB
	App          App
}

type App struct {
	Addressforlisten string
}

type DB struct {
	DriverName string
	host       string
	port       string
	user       string
	password   string
	dbname     string
	sslmode    string
}

type AdditionalDB struct {
	DriverName string
}

func (c *MainConfig) GetSqlPath() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		c.DB.host, c.DB.port, c.DB.user, c.DB.password, c.DB.dbname, c.DB.sslmode)
}

func NewConfig() (*MainConfig, error) {

	cfg, err := ParseConfigFile()
	if err != nil {
		return nil, err
	}

	return &MainConfig{
		App: App{
			Addressforlisten: cfg["app"]["addressforlisten"],
		},
		AdditionalDB: AdditionalDB{
			DriverName: cfg["additional_db"]["drivername"],
		},
		DB: DB{
			DriverName: cfg["db"]["drivername"],
			host:       cfg["db"]["host"],
			port:       cfg["db"]["port"],
			user:       cfg["db"]["user"],
			password:   cfg["db"]["password"],
			dbname:     cfg["db"]["dbname"],
			sslmode:    cfg["db"]["sslmode"],
		},
	}, nil
}
