package ogen

import (
	"net/http"

	"github.com/go-chi/cors"
)

func SetCors() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost",
			"http://194.87.226.28",
			"http://localhost:5173",
			"http://uoffer.ru",
			"http://admin.localhost",
			"http://admin.localhost:5173",
			"http://admin.uoffer.ru",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}
