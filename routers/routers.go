package routers

import (
	"goMux-api/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// To create a new mux router
func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.Use(controllers.LoggingMiddleware)
	router.Use(controllers.AuthMiddleware)

	// CORS LOGIC
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	return handler
}
