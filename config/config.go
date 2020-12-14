package config

import (
	"log"
	"movies-api/models"

	"github.com/hifx/envconfig"
)

var (
	Movies []models.Movie
	Users  []models.User
	ENV    APIConfig
)

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
	// res, err := ioutil.ReadFile("config/imdb.json")
	// if err != nil {
	// 	log.Fatalln("Error reading data from imdb.json file, err : ", err)
	// }
	// err = json.Unmarshal(res, &Movies)
	// if err != nil {
	// 	log.Fatalln("Error parsing json data from imdb.json file, err : ", err)
	// }
	// res, err = ioutil.ReadFile("config/user.json")
	// if err != nil {
	// 	log.Fatalln("Error reading data from user.json file, err : ", err)
	// }
	// err = json.Unmarshal(res, &Users)
	// if err != nil {
	// 	log.Fatalln("Error parsing json data from user.json file, err : ", err)
	// }
}
