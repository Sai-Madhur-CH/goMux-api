package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request is made for URL:", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := w.Header().Get("token")
		if token == "" {
			log.Println("Please provide the token")
			UnAuthorisedStatus(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func UnAuthorisedStatus(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{"status": "UnAuthorized"}
	JSONError(w, response, 401)
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
