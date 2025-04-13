package models

type ContactInfo struct {
	Email     string `json:"email"`
	Github    string `json:"github"`
	Twitter   string `json:"twitter"`
	Twitch    string `json:"twitch"`
	Spotify   string `json:"spotify"`
	LinkedIn  string `json:"linkedin"`
	Instagram string `json:"instagram"`
}
