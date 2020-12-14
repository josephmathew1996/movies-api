package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"movies-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/assert.v1"
)

func TestMovieGetHandler(t *testing.T) {
	expectedData := models.GetMovieResponse{
		Total:        2,
		NoOfPages:    1,
		ItemsPerPage: 10,
		Movies:       []models.Movie{{ID: 1, Name: "The Wizard of Oz", Director: "Victor Fleming", Genre: []string{"Adventure", "Family", "Fantasy", "Musical"}, Popularity: 83, IMDBScore: 8.3}, {ID: 2, Name: "Star Wars", Director: "George Lucas", Genre: []string{"Action", "Adventure", "Fantasy", "Sci-Fi"}, Popularity: 83, IMDBScore: 8.3}},
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/movies?page=1&count=10", nil)
	movie.Get(w, r)
	resp, _ := ioutil.ReadAll(w.Body)
	actual := struct {
		Status     string                  `json:"status"`
		StatusCode int                     `json:"statusCode"`
		Message    string                  `json:"message"`
		Data       models.GetMovieResponse `json:"data"`
	}{}
	json.Unmarshal(resp, &actual)

	movie1 := actual.Data.Movies[0]
	movie2 := actual.Data.Movies[1]

	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, expectedData.Total, actual.Data.Total)
	assert.Equal(t, expectedData.NoOfPages, actual.Data.NoOfPages)
	assert.Equal(t, expectedData.ItemsPerPage, actual.Data.ItemsPerPage)
	assert.Equal(t, expectedData.Movies[0], movie1)
	assert.Equal(t, expectedData.Movies[1], movie2)
	// assert.Equal(t, expectedData.Movies[0].ID, movie1.ID)
	// assert.Equal(t, expectedData.Movies[0].Name, movie1.Name)
	// assert.Equal(t, expectedData.Movies[0].Genre, movie1.Genre)
	// assert.Equal(t, expectedData.Movies[0].Popularity, movie1.Popularity)
	// assert.Equal(t, expectedData.Movies[0].IMDBScore, movie1.IMDBScore)
	// assert.Equal(t, expectedData.Movies[1].ID, movie2.ID)
	// assert.Equal(t, expectedData.Movies[1].Name, movie2.Name)
	// assert.Equal(t, expectedData.Movies[1].Genre, movie2.Genre)
	// assert.Equal(t, expectedData.Movies[1].Popularity, movie2.Popularity)
	// assert.Equal(t, expectedData.Movies[1].IMDBScore, movie2.IMDBScore)
}

func TestMovieCreateHandler(t *testing.T) {
	expectedData := models.Movie{ID: 253, Name: "King Kong", Director: "Merian C. Cooper", Genre: []string{"Adventure", "Fantasy", "Horror"}, Popularity: 80, IMDBScore: 8}
	b, _ := json.Marshal(expectedData)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/api/v1/movies", bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")
	movie.Create(w, r)
	resp, _ := ioutil.ReadAll(w.Body)
	actual := struct {
		Status     string       `json:"status"`
		StatusCode int          `json:"statusCode"`
		Message    string       `json:"message"`
		Data       models.Movie `json:"data"`
	}{}
	json.Unmarshal(resp, &actual)
	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, expectedData.ID, actual.Data.ID)
	assert.Equal(t, expectedData.Name, actual.Data.Name)
	assert.Equal(t, expectedData.Director, actual.Data.Director)
	assert.Equal(t, expectedData.Genre, actual.Data.Genre)
	assert.Equal(t, expectedData.Popularity, actual.Data.Popularity)
	assert.Equal(t, expectedData.IMDBScore, actual.Data.IMDBScore)
}

func TestMovieUpdateHandler(t *testing.T) {
	requestBody := models.Movie{ID: 253, Name: "King Kong", Director: "Merian C. Cooper", Genre: []string{"Adventure", "Fantasy", "Horror"}, Popularity: 80, IMDBScore: 8}
	b, _ := json.Marshal(requestBody)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/api/v1/movies/", bytes.NewReader(b))
	vars := map[string]string{
		"id": "253",
	}
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "application/json")
	movie.Update(w, r)
	resp, _ := ioutil.ReadAll(w.Body)
	actual := struct {
		Status     string      `json:"status"`
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
	}{}
	json.Unmarshal(resp, &actual)
	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, 200, actual.StatusCode)
}

func TestMovieDeleteHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/api/v1/movies/", nil)
	vars := map[string]string{
		"id": "253",
	}
	r = mux.SetURLVars(r, vars)
	r.Header.Set("Content-Type", "application/json")
	movie.Delete(w, r)
	resp, _ := ioutil.ReadAll(w.Body)
	actual := struct {
		Status     string      `json:"status"`
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
	}{}
	json.Unmarshal(resp, &actual)
	assert.Equal(t, 200, w.Result().StatusCode)
	assert.Equal(t, 200, actual.StatusCode)
	assert.Equal(t, "SUCCESS", actual.Status)
	assert.Equal(t, "Movie deleted successfully", actual.Message)
	assert.Equal(t, nil, actual.Data)
}
