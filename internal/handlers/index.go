package handlers

import (
	"encoding/json"
	"net/http"

	"fyi/internal/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	info := models.APIInfo{
		Name:        "Nicholai's Personal API",
		Version:     "1.0.0",
		Description: "A personal API for Nicholai",
		Endpoints: []models.Endpoint{
			{Path: "/", Method: "GET", Description: "API information"},
			{Path: "/about", Method: "GET", Description: "About Nicholai"},
			{Path: "/projects", Method: "GET", Description: "List of projects"},
			{Path: "/contact", Method: "GET", Description: "Contact information"},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
