package handlers

import (
	"encoding/json"
	"net/http"

	"fyi/internal/models"
)

func Projects(w http.ResponseWriter, r *http.Request) {
	projects := []models.Project{
		{
			Name:         "My API",
			Description:  "the API powering this website",
			Technologies: []string{"go", "docker"},
			URL:          "https://api.itsnicholai.fyi",
			Github:       "https://github.com/n3k0lai/fyi",
		},
		{
			Name:         "My Website",
			Description:  "the website you're currently viewing",
			Technologies: []string{"typescript"},
			URL:          "https://itsnicholai.fyi",
			Github:       "https://github.com/nicholai/itsnicholai.fyi",
		},
	}
	json.NewEncoder(w).Encode(projects)
}
