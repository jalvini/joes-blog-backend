package router

import (
	"github.com/gorilla/mux"
	"joes-blog-backend/jwt"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/welcome", jwtToken.CheckUserInfoToken)
	r.HandleFunc("/signin", jwtToken.Signin)
	r.HandleFunc("/refresh", jwtToken.Refresh)

	return r
}
