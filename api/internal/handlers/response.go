package handlers

import (
	"encoding/json"
	"net/http"
)

func respondError(w http.ResponseWriter, err string, statusCode int) {
	respondJSON(w, map[string]string{"error": err}, statusCode)
}

func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
