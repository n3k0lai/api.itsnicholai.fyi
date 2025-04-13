package handlers

import (
	"encoding/json"
	"net/http"

	"fyi/internal/models"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	info := models.ContactInfo{
		Email:     "nicholai@comfy.sh",
		Github:    "https://github.com/n3k0lai",
		Twitter:   "https://x.com/n3k0lai",
		Twitch:    "https://twitch.tv/n3k0lai",
		Spotify:   "https://open.spotify.com/user/n3k0lai",
		LinkedIn:  "https://linkedin.com/in/nicholai",
		Instagram: "https://instagram.com/n3k0lai",
	}
	json.NewEncoder(w).Encode(info)
}
