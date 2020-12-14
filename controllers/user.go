package controllers

import (
	"encoding/json"
	"log"
	"movies-api/config"
	"movies-api/models"
	"movies-api/util"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//User holds the interface of User entity
type User struct {
	User models.UserService
}

//NewUser initialises User
func NewUser(user models.UserService) *User {
	return &User{
		User: user,
	}
}

//Auth handler
func (u *User) Auth(w http.ResponseWriter, r *http.Request) {
	var (
		request  models.User
		response models.AuthToken
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error parsing request body : ", err)
		ProcessResponse(w, "Invalid request body", http.StatusBadRequest, nil)
		return
	}
	if request.Email == "" || request.Password == "" {
		log.Println("Either email or password is missing")
		ProcessResponse(w, "Please provide valid details", http.StatusBadRequest, nil)
		return
	}
	//Validation success
	user, err := u.User.GetUserByEmail(request.Email)
	if err != nil {
		log.Println("Error fetching user by email : ", err)
		ProcessResponse(w, "Failed", http.StatusInternalServerError, nil)
		return
	}
	if (user == models.User{}) {
		log.Println("User not found with provided email")
		ProcessResponse(w, "User not found", http.StatusNotFound, nil)
		return
	}
	isPasswordMatched := util.CheckPasswordHash(request.Password, user.Password)
	if user.Email != request.Email || !isPasswordMatched {
		log.Println("Invalid email or password")
		ProcessResponse(w, "Invalid email or password", http.StatusUnauthorized, nil)
		return
	}
	user.Password = ""
	token, expiresAt, err := createToken(user)
	if err != nil {
		log.Println("Error generating JWT Token : ", err)
		ProcessResponse(w, "Failed", http.StatusInternalServerError, nil)
		return
	}
	response.TokenType = "Bearer"
	response.Token = token
	response.ExpiresAt = expiresAt
	response.UserDetail = user
	ProcessResponse(w, "Successfully logged in", 200, response)
	return
}

func createToken(user models.User) (string, int64, error) {
	expiresAt := time.Now().Add(time.Second * time.Duration(config.ENV.TokenExpiry)).Unix()
	claims := &models.AuthTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.ENV.JWTSecret))
	if err != nil {
		return "", 0, err
	}
	return tokenString, expiresAt, nil
}
