package routes

import (
	"movies-api/controllers"
	"movies-api/middleware"
	"movies-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

//Set sets all the routes
func Set(services models.Services) http.Handler {

	router := mux.NewRouter()

	userController := controllers.NewUser(services.User)
	movieController := controllers.NewMovie(services.Movie)

	r := router.PathPrefix("/api/v1").Subrouter()

	r.HandleFunc("/auth", userController.Auth).Methods("POST")

	r.HandleFunc("/movies", movieController.Get).Methods("GET")
	r.HandleFunc("/movies", middleware.AuthMiddleware(movieController.Create)).Methods("POST")
	r.HandleFunc("/movies/{id}", middleware.AuthMiddleware(movieController.Update)).Methods("PUT")
	r.HandleFunc("/movies/{id}", middleware.AuthMiddleware(movieController.Delete)).Methods("DELETE")
	return router
}
