package config

import (
	"fmt"
)

const (
	configPath      = "configs/config.yml"
	configLocalPath = "configs/config_local.yml"
)

type Config interface {
	GetDsnDB() string
	GetDsnAddDB() string
	GetDsnLogger() string
	GetHostLogger() string
	GetDsnCache() string
}

type MainConfig struct {
	AdditionalDB AdditionalDB
	DB           DB
	App          App
	Logger       Logger
	Cache        Cache
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
	dsn        string
}

type Logger struct {
	DriverName string
	dsn        string
	host       string
}

type Cache struct {
	DriverName string
	dsn        string
}

func (c *MainConfig) GetDsnDB() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		c.DB.host, c.DB.port, c.DB.user, c.DB.password, c.DB.dbname, c.DB.sslmode)
}

func (c *MainConfig) GetDsnAddDB() string {
	return c.AdditionalDB.dsn
}

func (c *MainConfig) GetDsnLogger() string {
	return c.Logger.dsn
}

func (c *MainConfig) GetHostLogger() string {
	return c.Logger.host
}

func (c *MainConfig) GetDsnCache() string {
	return c.Cache.dsn
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
			dsn:        cfg["additional_db"]["dsn"],
		},
		Logger: Logger{
			DriverName: cfg["logger"]["drivername"],
			dsn:        cfg["logger"]["dsn"],
			host:       cfg["logger"]["host"],
		},
		Cache: Cache{
			DriverName: cfg["cache"]["drivername"],
			dsn:        cfg["cache"]["dsn"],
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
