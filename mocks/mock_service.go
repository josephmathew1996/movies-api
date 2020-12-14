package mocks

import (
	"database/sql"
	"movies-api/models"
)

type MockServices struct {
	User  models.UserService
	Movie models.MovieService
}

func NewMockServices(db *sql.DB) MockServices {
	return MockServices{
		User:  UserService{db: db},
		Movie: MovieService{db: db},
	}
}
