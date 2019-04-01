package jwtToken

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"joes-blog-backend/helpers"
	"joes-blog-backend/jwt/models"
	"net/http"
)

func CheckUserInfoToken(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "GET" {
		tokenString := r.Header.Get("Authorization")
		token, valid := CheckToken(tokenString)
		if valid {
			w.Header().Set("Content-Type", "application/json")
			user := helpers.ReadUser(token.Username)

			json.Marshal(user)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

func CheckToken(tokenString string) (models.UserToken, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return models.JwtKey, nil
	})

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return models.UserToken{
			claims.Username,
			claims.StandardClaims.ExpiresAt,
		}, true
	}

	return models.UserToken{}, false
}
