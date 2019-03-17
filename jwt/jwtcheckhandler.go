package jwtToken

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"joes-blog-backend/jwt/models"
	"net/http"
)

func CheckToken(tokenString string) (models.UserToken, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		if validate(claims.Username) == true {
			return models.UserToken{
				claims.Username,
				claims.StandardClaims.ExpiresAt,
			}, true
		}
	} else {
		fmt.Println(err)
	}
	return models.UserToken{}, false
}

func validate(username string) bool {
	if username == "Joe" { // TODO && unix timestamp == Positive NumberZz
		return true
	}
	return false
}

func CheckUserInfoToken(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "GET" {
		tokenString := r.Header.Get("Authorization")
		token, valid := CheckToken(tokenString)
		if valid {
			w.Header().Set("Content-Type", "application/json")
			fmt.Println(token.Username)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
