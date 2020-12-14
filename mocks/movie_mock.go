package mocks

import (
	"database/sql"
	"movies-api/models"
)

type MovieService struct {
	db *sql.DB
}

//GetAllMovies gets all movies
func (a MovieService) GetAllMovies(filter models.MovieFilter) ([]models.Movie, int, error) {
	return []models.Movie{{ID: 1, Name: "The Wizard of Oz", Director: "Victor Fleming", Genre: []string{"Adventure", "Family", "Fantasy", "Musical"}, Popularity: 83, IMDBScore: 8.3}, {ID: 2, Name: "Star Wars", Director: "George Lucas", Genre: []string{"Action", "Adventure", "Fantasy", "Sci-Fi"}, Popularity: 83, IMDBScore: 8.3}}, 2, nil

}

//GetMovieByID returns specific movie details by id
func (a MovieService) GetMovieByID(id int) (models.Movie, error) {
	return models.Movie{ID: 253, Name: "King Kong", Director: "Merian C. Cooper", Genre: []string{"Adventure", "Fantasy", "Horror"}, Popularity: 80, IMDBScore: 8}, nil
}

//Create creates a new movie
func (a MovieService) Create(movie models.Movie) (models.Movie, error) {
	return models.Movie{ID: 253, Name: "King Kong", Director: "Merian C. Cooper", Genre: []string{"Adventure", "Fantasy", "Horror"}, Popularity: 80, IMDBScore: 8}, nil
}

//Update updates a movie
func (a MovieService) Update(movie models.Movie) error {
	return nil
}

//Delete deletes a movie
func (a MovieService) Delete(id int) error {
	return nil
}
