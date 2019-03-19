package main

import (
	"github.com/rs/cors"
	"joes-blog-backend/router"
	"log"
	"net/http"
)

var c = cors.New(cors.Options{
	AllowedOrigins:   []string{"*"},
	AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
	AllowCredentials: true,
	AllowedHeaders:   []string{"*"},
	Debug:            false,
})

func main() {

	r := router.Routes()

	log.Fatal(http.ListenAndServe(":8001", c.Handler(r)))
}
