package models

import "database/sql"

//Services holds all the services
type Services struct {
	User  UserService
	Movie MovieService
}

//NewServices initialises the services
func NewServices(db *sql.DB) Services {
	return Services{
		User:  UserStorer{db},
		Movie: MovieStorer{db},
	}
}
