package models

import "github.com/dgrijalva/jwt-go"

//User holds the user entity
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role" `
}

//AuthToken holds the bearer token
type AuthToken struct {
	UserDetail User   `json:"userDetails"`
	TokenType  string `json:"token_type"`
	Token      string `json:"access_token"`
	ExpiresAt  int64  `json:"expires_in"`
}

//AuthTokenClaim holds the claim object parsed from the authorization header
type AuthTokenClaim struct {
	*jwt.StandardClaims
	User
}

//Movie holds the movie entity
type Movie struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Director   string   `json:"director"`
	Genre      []string `json:"genre"`
	Popularity float64  `json:"99popularity"`
	IMDBScore  float64  `json:"imdb_score"`
}

//MovieFilter holds the movie filters
type MovieFilter struct {
	Page     interface{} `json:"page"`
	Count    interface{} `json:"count"`
	SortBy   string      `json:"sortby"`
	Asc      interface{} `json:"asc"`
	Movie    string      `json:"name"`
	Director string      `json:"director"`
}

//GetMovieResponse holds get all movies resposne
type GetMovieResponse struct {
	Total        int     `json:"total"`
	NoOfPages    int     `json:"noOfPages"`
	ItemsPerPage int     `json:"itemsPerPage"`
	Movies       []Movie `json:"movies"`
}
