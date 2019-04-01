package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"joes-blog-backend/helpers"
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
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	helpers.ReadUser("jalvini")
	r := router.Routes()

	log.Fatal(http.ListenAndServe(":8001", c.Handler(r)))
}
