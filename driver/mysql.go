package driver

import (
	"database/sql"
	"fmt"
	"movies-api/config"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	Host     string
	User     string
	Password string
	Port     int
	Db       string
}

// ConnectToMySQL takes mysql config, forms the connection string and connects to mysql.
func ConnectToMySQL() (*sql.DB, error) {
	// get the mysql configs from env:
	conf := MySQLConfig{
		Host:     config.ENV.Mysql.Host,
		User:     config.ENV.Mysql.User,
		Password: config.ENV.Mysql.Password,
		Port:     config.ENV.Mysql.Port,
		Db:       config.ENV.Mysql.DB,
	}
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.Db)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
