package jwtToken

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"joes-blog-backend/helpers"
	"joes-blog-backend/jwt/models"
	"net/http"
	"time"
)

type Exception struct {
	Message string `json:"message"`
}

// Create the Signin handler
func Signin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	if (*r).Method == "POST" {
		var t models.Credentials
		err := decoder.Decode(&t)

		if err != nil {
			panic(err)
		}

		var password = "Make00"
		hash := helpers.HashPassword(password)

		getUser := helpers.ReadUser("jalvini")

		checkHash := helpers.ComparePassword(hash, t.Password)

		if t.Username == getUser.Username && checkHash == true {
			fmt.Println(t.Username)
			// Declare the expiration time of the token
			// here, we have kept it as 5 minutes
			expirationTime := time.Now().Add(48 * time.Hour)
			// Create the JWT claims, which includes the username and expiry time
			claims := &models.Claims{
				Username: t.Username,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

			// Declare the token with the algorithm used for signing, and the claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			// Create the JWT string
			tokenString, err := token.SignedString(models.JwtKey)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			tokenByte, err := json.Marshal(tokenString)
			w.Write(tokenByte)
		} else {

			in := `{ error: { message: 'Username or password is incorrect' } }`

			rawIn := json.RawMessage(in)

			bytes, err := rawIn.MarshalJSON()
			if err != nil {
				panic(err)
			}

			w.Write(bytes)
		}
	}
}
