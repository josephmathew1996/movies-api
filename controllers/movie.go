package controllers

import (
	"encoding/json"
	"log"
	"math"
	"movies-api/config"
	"movies-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//Movie holds the interface of Movie entity
type Movie struct {
	Movie models.MovieService
}

//NewMovie initialises Movie
func NewMovie(movie models.MovieService) *Movie {
	return &Movie{
		Movie: movie,
	}
}

//Get handler
func (m *Movie) Get(w http.ResponseWriter, r *http.Request) {
	filter := models.MovieFilter{}
	params := make(map[string]interface{})
	for k, v := range r.URL.Query() { //only considering one value of each parameter
		params[k] = v[0]
	}
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		log.Println("Error marshalling the movie query params : ", err)
		ProcessResponse(w, "Invalid request query params", http.StatusBadRequest, nil)
		return
	}
	err = json.Unmarshal(paramsJSON, &filter)
	if err != nil {
		log.Println("Error parsing the movie query params : ", err)
		ProcessResponse(w, "Invalid request query params", http.StatusBadRequest, nil)
		return
	}
	if filter.Page == nil || filter.Page == "0" {
		filter.Page = 1
	} else {
		filter.Page, err = strconv.Atoi(filter.Page.(string))
		if err != nil {
			ProcessResponse(w, "Failed", http.StatusInternalServerError, nil)
			return
		}
	}
	if filter.Count == nil || filter.Count == "0" {
		filter.Count = config.ENV.ItemsPerPage
	} else {
		filter.Count, err = strconv.Atoi(filter.Count.(string))
		if err != nil {
			ProcessResponse(w, "Failed", http.StatusInternalServerError, nil)
			return
		}
	}
	movies, total, err := m.Movie.GetAllMovies(filter)
	if err != nil {
		log.Println("Error fetching movies : ", err)
		ProcessResponse(w, "Failed", http.StatusInternalServerError, nil)
		return
	}
	response := models.GetMovieResponse{}
	response.Total = total
	response.NoOfPages = int(math.Ceil(float64(total) / float64(filter.Count.(int))))
	response.ItemsPerPage = filter.Count.(int)
	response.Movies = movies
	ProcessResponse(w, "Movies fetch success", http.StatusOK, response)
	return
}

//Create handler
func (m *Movie) Create(w http.ResponseWriter, r *http.Request) {
	if session, ok := context.Get(r, "Session").(*models.AuthTokenClaim); ok {
		if session.Role != "admin" {
			log.Println("User not authorized")
			ProcessResponse(w, "Unauthorized user", http.StatusUnauthorized, nil)
			return
		}
	}
	movie := models.Movie{}
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Println("Error marshalling the movie request body : ", err)
		ProcessResponse(w, "Invalid request body", http.StatusBadRequest, nil)
		return
	}
	if movie.Name == "" || movie.Director == "" || len(movie.Genre) == 0 || movie.Popularity == 0 || movie.IMDBScore == 0 {
		log.Println("Validation failed")
		ProcessResponse(w, "Please provide valid details", http.StatusBadRequest, nil)
		return
	}
	movie, err = m.Movie.Create(movie)
	if err != nil {
		log.Println("Error creating new movie : ", err)
		ProcessResponse(w, "Movie creation failed", http.StatusInternalServerError, nil)
		return
	}
	ProcessResponse(w, "Movie created successfully", http.StatusOK, movie)
}

//Update handler
func (m *Movie) Update(w http.ResponseWriter, r *http.Request) {
	if session, ok := context.Get(r, "Session").(*models.AuthTokenClaim); ok {
		if session.Role != "admin" {
			log.Println("User not authorized")
			ProcessResponse(w, "Unauthorized user", http.StatusUnauthorized, nil)
			return
		}
	}
	vars := mux.Vars(r)
	var movie models.Movie
	id := vars["id"]
	movieID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid path parameter")
		ProcessResponse(w, "Please provide a valid movie ID", http.StatusBadRequest, nil)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Println("Error marshalling the movie request body : ", err)
		ProcessResponse(w, "Invalid request body", http.StatusBadRequest, nil)
		return
	}
	if movieID < 1 || (movieID != movie.ID) {
		log.Printf("Invalid movie ID. Path parameter movie ID = %d, Request body movie ID = %d\n", movieID, movie.ID)
		ProcessResponse(w, "Please provide a valid movie ID", http.StatusBadRequest, nil)
		return
	}
	if movie.Name == "" || movie.Director == "" || len(movie.Genre) == 0 || movie.Popularity == 0 || movie.IMDBScore == 0 {
		log.Println("Validation failed")
		ProcessResponse(w, "Please provide valid details", http.StatusBadRequest, nil)
		return
	}
	movieByID, err := m.Movie.GetMovieByID(movieID)
	if err != nil {
		log.Println("Error fetching movie by ID : ", err)
		ProcessResponse(w, "Movie updation failed", http.StatusInternalServerError, nil)
		return
	}
	if movieByID.ID == 0 {
		log.Println("Movie not found")
		ProcessResponse(w, "Movie not found", http.StatusNotFound, nil)
		return
	}
	err = m.Movie.Update(movie)
	if err != nil {
		log.Println("Error updating movie : ", err)
		ProcessResponse(w, "Movie updation failed", http.StatusInternalServerError, nil)
		return
	}
	ProcessResponse(w, "Movie updated successfully", http.StatusOK, nil)
	return
}

//Delete handler
func (m *Movie) Delete(w http.ResponseWriter, r *http.Request) {
	if session, ok := context.Get(r, "Session").(*models.AuthTokenClaim); ok {
		if session.Role != "admin" {
			log.Println("User not authorized")
			ProcessResponse(w, "Unauthorized user", http.StatusUnauthorized, nil)
			return
		}
	}
	vars := mux.Vars(r)
	id := vars["id"]
	movieID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid path parameter")
		ProcessResponse(w, "Please provide a valid movie ID", http.StatusBadRequest, nil)
		return
	}
	if movieID < 1 {
		log.Printf("Invalid movie ID. Path parameter movie ID = %d\n", movieID)
		ProcessResponse(w, "Please provide a valid movie ID", http.StatusBadRequest, nil)
		return
	}
	movieByID, err := m.Movie.GetMovieByID(movieID)
	if err != nil {
		log.Println("Error fetching movie by ID : ", err)
		ProcessResponse(w, "Movie deletion failed", http.StatusInternalServerError, nil)
		return
	}
	if movieByID.ID == 0 {
		log.Println("Movie not found")
		ProcessResponse(w, "Movie not found", http.StatusNotFound, nil)
		return
	}
	err = m.Movie.Delete(movieID)
	if err != nil {
		log.Println("Error deleting movie : ", err)
		ProcessResponse(w, "Movie deletion failed", http.StatusInternalServerError, nil)
		return
	}
	ProcessResponse(w, "Movie deleted successfully", http.StatusOK, nil)
	return
}
