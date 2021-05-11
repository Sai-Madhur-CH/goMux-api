package controllers

import (
	"goMux-api/utils"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request is made for URL:", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			log.Println("Please provide the token")
			utils.UnAuthorisedStatus(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
