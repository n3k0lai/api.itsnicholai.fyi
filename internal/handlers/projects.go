package handlers

import (
	"encoding/json"
	"net/http"

	"fyi/internal/models"
)

func Projects(w http.ResponseWriter, r *http.Request) {
	projects := []models.Project{
		{
			Name:         "Project 1",
			Description:  "Description of Project 1",
			Technologies: []string{"Tech 1", "Tech 2"},
			URL:          "https://project1-url.com",
			Github:       "https://github.com/nicholai/project1",
		},
		{
			Name:         "Project 2",
			Description:  "Description of Project 2",
			Technologies: []string{"Tech 3", "Tech 4"},
			URL:          "https://project2-url.com",
			Github:       "https://github.com/nicholai/project2",
		},
	}
	json.NewEncoder(w).Encode(projects)
}
