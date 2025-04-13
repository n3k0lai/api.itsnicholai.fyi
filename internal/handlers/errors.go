package handlers

import (
	"encoding/json"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":               "Not found",
		"message":             "The requested URL was not found on this server.",
		"available_endpoints": []string{"/", "/about", "/projects", "/contact"},
	})
}

func ServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   "Server error",
		"message": "An internal server error occurred.",
	})
}
