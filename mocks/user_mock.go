package mocks

import (
	"database/sql"
	"movies-api/models"
)

type UserService struct {
	db *sql.DB
}

//GetUserByEmail gets user by id
func (a UserService) GetUserByEmail(email string) (models.User, error) {
	if email == "" {
		return models.User{}, nil
	}
	return models.User{ID: 1, Name: "Admin", Email: "admin@mailsac.com", Password: "$2a$04$F8KPSzu0962d88ZFKcyf8ua3CvD5SaWOKHcTZZ0DNC8OP8VhYHka2", Role: "admin"}, nil
}
