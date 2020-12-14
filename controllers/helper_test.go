package controllers

import (
	"database/sql"
	"movies-api/mocks"
)

var (
	db       *sql.DB
	services = mocks.NewMockServices(db)
	movie    = NewMovie(services.Movie)
	user     = NewUser(services.User)
)
