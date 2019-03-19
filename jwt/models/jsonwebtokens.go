package models

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

var JwtKey = []byte(os.Getenv("JSONWebTokenPassword"))

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserToken struct {
	Username  string
	TimeStamp int64
}
