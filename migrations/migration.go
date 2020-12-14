package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"movies-api/config"
	"movies-api/util"
	"strings"
)

//MigrateDBData populates DB with data from json file
func MigrateDBData(db *sql.DB) {
	log.Println("No of movies = ", len(config.Movies))
	log.Println("No of users = ", len(config.Users))
	//INSERT TO DB...
	// migrateMovies(db)
	// migrateUsers(db)
}

func migrateMovies(db *sql.DB) {
	var args []interface{}
	var controlString []string
	for _, movie := range config.Movies {
		controlString = append(controlString, "(?,?,?,?,?)")
		for i := range movie.Genre { //removing white spaces
			movie.Genre[i] = strings.TrimSpace(movie.Genre[i])
		}
		args = append(args, movie.Name, movie.Director, strings.Join(movie.Genre, ","), movie.Popularity, movie.IMDBScore)
	}
	query := fmt.Sprintf("INSERT INTO movies(name,director,genre,popularity,imdb_score) VALUES %s", strings.Join(controlString, ","))
	stmnt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln("Failed to prepare SQL statement, err : ", err)
	}
	defer stmnt.Close()
	_, err = stmnt.Exec(args...)
	if err != nil {
		log.Fatalln("Error migrating json data to DB from imdb.json file, err : ", err)
	}
}

func migrateUsers(db *sql.DB) {
	var args []interface{}
	var controlString []string
	for _, user := range config.Users {
		controlString = append(controlString, "(?,?,?,?)")
		passwordHash, err := util.HashPassword(user.Password)
		if err != nil {
			log.Fatalln("Error hashing password : ", err)
		}
		args = append(args, user.Name, user.Email, passwordHash, user.Role)
	}
	query := fmt.Sprintf("INSERT INTO users(name, email, password, role) VALUES %s", strings.Join(controlString, ","))
	stmnt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln("Failed to prepare SQL statement, err : ", err)
	}
	defer stmnt.Close()
	_, err = stmnt.Exec(args...)
	if err != nil {
		log.Fatalln("Error migrating json data to DB from user.json file, err : ", err)
	}
}
