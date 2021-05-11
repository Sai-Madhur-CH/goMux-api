package utils

import (
	"encoding/json"
	"net/http"
)

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
