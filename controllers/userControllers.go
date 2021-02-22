package controllers

import (
	"encoding/json"
	"goMux-api/models"
	"goMux-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// To create the user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	utils.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// To get all the users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	utils.DB.Find(&users)
	response := map[string]interface{}{"status": "success"}
	response["rows"] = users
	json.NewEncoder(w).Encode(response)
}

// To get the selected user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user models.User
	utils.DB.First(&user, params["id"])
	response := map[string]interface{}{"status": "success"}
	response["user"] = user
	json.NewEncoder(w).Encode(response)

}
