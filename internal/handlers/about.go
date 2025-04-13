package handlers

import (
	"encoding/json"
	"net/http"

	"fyi/internal/models"
)

func About(w http.ResponseWriter, r *http.Request) {
	info := models.AboutInfo{
		Name:      "Nicholai",
		Bio:       "Senior Remote Software Engineer",
		Skills:    []string{"Skill 1", "Skill 2", "Skill 3"},
		Interests: []string{"Yoga", "Livestreaming", "Social Outreach"},
	}
	json.NewEncoder(w).Encode(info)
}
