package main

import (
	"log"
	"movies-api/config"
	"movies-api/driver"
	"movies-api/models"
	"movies-api/routes"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	log.Println("-------Movie API Server-------")
	config.Load()

	db, err := driver.ConnectToMySQL()
	if err != nil {
		log.Fatalln("Could not connect to sql, err : ", err)
	}
	log.Println("MySQL connection successfull...")

	// migrations.MigrateDBData(db)

	services := models.NewServices(db)
	router := routes.Set(services)

	log.Fatal(http.ListenAndServe(":"+config.ENV.ServerPort, handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

	if err != nil {
		log.Fatalf("HTTP Server stopped with error %s", err)
	}
}
