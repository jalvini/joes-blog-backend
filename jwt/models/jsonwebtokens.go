package models

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

var JwtKey = []byte(os.Getenv("JSONWebTokenPassword"))

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserToken struct {
	Username  string
	TimeStamp int64
}
