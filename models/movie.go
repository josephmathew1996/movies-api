package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//MovieStorer struct implements the IMovie interface
type MovieStorer struct {
	db *sql.DB
}

//GetAllMovies gets all movies
func (ms MovieStorer) GetAllMovies(filter MovieFilter) ([]Movie, int, error) {
	var (
		total           int
		whereClause     string
		sortByClause    = "ORDER BY id" //default sorting
		keyword         string
		query           string
		finalQuery      string
		whereConditions = make([]string, 0)
		whereParams     = make([]interface{}, 0)
		movies          = []Movie{}
		err             error
	)
	if filter.SortBy != "" {
		sortByClause = "ORDER BY " + filter.SortBy
	}
	if filter.Asc != nil {
		isAsc, err := strconv.ParseBool(filter.Asc.(string))
		if err != nil {
			return []Movie{}, 0, err
		}
		if isAsc {
			sortByClause += " ASC"
		} else {
			sortByClause += " DESC"
		}
	}
	if filter.Movie != "" {
		whereConditions = append(whereConditions, "name LIKE ?")
		keyword = filter.Movie + "%"
		whereParams = append(whereParams, keyword)
	}
	if filter.Director != "" {
		whereConditions = append(whereConditions, "director LIKE ?")
		keyword = filter.Director + "%"
		whereParams = append(whereParams, keyword)
	}
	if len(whereConditions) > 0 {
		whereClause = "WHERE " + strings.Join(whereConditions, " AND ")
	}
	query = "SELECT id, name, director, genre, popularity, imdb_score FROM movies %s %s LIMIT ? OFFSET ?"
	finalQuery = fmt.Sprintf(query, whereClause, sortByClause)
	whereParams = append(whereParams, filter.Count, (filter.Page.(int)-1)*filter.Count.(int))
	stmnt, err := ms.db.Prepare(finalQuery)
	if err != nil {
		return []Movie{}, 0, err
	}
	rows, err := stmnt.Query(whereParams...)
	if err != nil {
		return []Movie{}, 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var movie Movie
		var genre string
		err = rows.Scan(&movie.ID, &movie.Name, &movie.Director, &genre, &movie.Popularity, &movie.IMDBScore)
		if err != nil {
			return []Movie{}, 0, err
		}
		movie.Genre = strings.Split(genre, ",")
		movies = append(movies, movie)
	}
	query = "SELECT COUNT(*) FROM movies %s %s"
	finalQuery = fmt.Sprintf(query, whereClause, sortByClause)
	whereParams = whereParams[:len(whereParams)-2] //removing the LIMIT AND OFFSET params
	stmnt, err = ms.db.Prepare(finalQuery)
	if err != nil {
		return []Movie{}, 0, err
	}
	row := stmnt.QueryRow(whereParams...)
	err = row.Scan(&total)
	if err != nil {
		return []Movie{}, 0, err
	}
	return movies, total, nil
}

//GetMovieByID returns specific movie details by id
func (ms MovieStorer) GetMovieByID(id int) (Movie, error) {
	var movie Movie
	var genre string
	query := "SELECT id, name, director, genre, popularity, imdb_score FROM movies WHERE id=?"
	stmnt, err := ms.db.Prepare(query)
	if err != nil {
		return Movie{}, err
	}
	row := stmnt.QueryRow(id)
	err = row.Scan(&movie.ID, &movie.Name, &movie.Director, &genre, &movie.Popularity, &movie.IMDBScore)
	if err != nil {
		if err == sql.ErrNoRows {
			return Movie{}, nil
		}
		return Movie{}, err
	}
	movie.Genre = strings.Split(genre, ",")
	return movie, nil
}

//Create creates a new movie
func (ms MovieStorer) Create(movie Movie) (Movie, error) {
	query := "INSERT INTO movies(name, director, genre, popularity, imdb_score) VALUES(?,?,?,?,?)"
	stmnt, err := ms.db.Prepare(query)
	if err != nil {
		return Movie{}, err
	}
	res, err := stmnt.Exec(movie.Name, movie.Director, strings.Join(movie.Genre, ","), movie.Popularity, movie.IMDBScore)
	if err != nil {
		return Movie{}, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return Movie{}, err
	}
	movie.ID = int(lastInsertID)
	return movie, nil
}

//Update updates a  movie
func (ms MovieStorer) Update(movie Movie) error {
	query := "UPDATE movies SET name=?, director=?, genre=?, popularity=?, imdb_score=? WHERE id=?"
	stmnt, err := ms.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmnt.Exec(movie.Name, movie.Director, strings.Join(movie.Genre, ","), movie.Popularity, movie.IMDBScore, movie.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("No of rows affected after movie updation = ", rowsAffected)
	return nil
}

//Delete deletes a  movie
func (ms MovieStorer) Delete(id int) error {
	query := "DELETE FROM movies WHERE id=?"
	stmnt, err := ms.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmnt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println("No of rows affected after movie deletion = ", rowsAffected)
	return nil
}
