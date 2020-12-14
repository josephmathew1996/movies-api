package middleware

import (
	"fmt"
	"movies-api/config"
	"movies-api/controllers"
	"movies-api/models"
	"net/http"
	"strings"

	"github.com/gorilla/context"

	"github.com/dgrijalva/jwt-go"
)

//AuthMiddleware authenticates user
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				tknStr := bearerToken[1]
				claims := &models.AuthTokenClaim{}
				token, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Error verifying token signing algorithm")
					}
					return []byte(config.ENV.JWTSecret), nil
				})
				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						controllers.ProcessResponse(w, "Invalid token signature", http.StatusUnauthorized, nil)
						return
					}
					controllers.ProcessResponse(w, "Invalid authorization token", http.StatusUnauthorized, nil)
					return
				}
				if !token.Valid {
					controllers.ProcessResponse(w, "Invalid authorization token", http.StatusUnauthorized, nil)
					return
				}
				context.Set(r, "Session", claims)
				next(w, r)
			} else {
				controllers.ProcessResponse(w, "Invalid authorization token", http.StatusUnauthorized, nil)
			}
		} else {
			controllers.ProcessResponse(w, "An authorization header is required", http.StatusUnauthorized, nil)
		}
	})
}
