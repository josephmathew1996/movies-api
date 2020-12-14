package models

import (
	"database/sql"
)

//UserStorer struct implements the IUser interface
type UserStorer struct {
	db *sql.DB
}

//GetUserByEmail gets user by id
func (a UserStorer) GetUserByEmail(email string) (User, error) {
	var user User
	stmnt, err := a.db.Prepare("SELECT id, name, email, password, role FROM users WHERE email=?")
	if err != nil {
		return User{}, err
	}
	row := stmnt.QueryRow(email)
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		}
		return User{}, err
	}
	return user, nil
}
