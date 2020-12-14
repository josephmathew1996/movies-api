package main

import (
	"log"
	"movies-api/config"
	"movies-api/driver"
	"movies-api/models"
	"movies-api/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	services := models.NewServices(db)
	router := routes.Set(services)

	go func() {
		log.Fatal(http.ListenAndServe(":"+config.ENV.ServerPort, handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
	}()

	log.Printf("Started HTTP server on port %s\n", config.ENV.ServerPort)
	exitCode := waitForStop()
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func waitForStop() int {
	okSig := make(chan os.Signal, 2)
	signal.Notify(okSig, syscall.SIGTERM)
	failSig := make(chan os.Signal, 2)
	signal.Notify(failSig, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	for {
		select {
		case sig := <-okSig:
			log.Printf("Exit OK on Signal  %s", sig)
			return 0
		case sig := <-failSig:
			log.Printf("Exit FAIL on Signal  %s", sig)
			return 2
		}
	}
}
