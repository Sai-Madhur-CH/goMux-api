package routers

import (
	"goMux-api/controllers"

	"github.com/gorilla/mux"
)

// To create a new mux router
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.Use(controllers.LoggingMiddleware)
	router.Use(controllers.AuthMiddleware)
	return router
}
