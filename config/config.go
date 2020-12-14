package config

import (
	"log"

	"github.com/hifx/envconfig"
)

var ENV APIConfig

// APIConfig holds the configuration for the api
type APIConfig struct {
	ServerPort   string
	ItemsPerPage int
	JWTSecret    string
	TokenExpiry  int64
	Mysql        MysqlConfig
}

// MysqlConfig holds configuration values for mysql
type MysqlConfig struct {
	Host               string
	Port               int
	User               string
	Password           string
	DB                 string
	MaxOpenConnections int `envconfig:"default=10"`
	MaxIdleConnections int `envconfig:"default=5"`
}

//Load loads data from json file
func Load() {
	err := envconfig.InitWithPrefix(&ENV, "MOVIES_API")
	if err != nil {
		log.Fatalln("Error while loading configuration : ", err)
	}
}
