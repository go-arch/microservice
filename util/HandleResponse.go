package util

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func  HandleServiceError(errors []error, w http.ResponseWriter){
	if len(errors) > 0 {
		switch errors[0] {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, errors[0].Error())
		}
		return
	}
}
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
