package models

import "time"

type Experience struct {
	Title       string     `json:"title"`
	Company     string     `json:"company"`
	Location    string     `json:"location"`
	StartDate   *time.Time `json:"start_date"`
	Active      bool       `json:"active"`
	EndDate     *time.Time `json:"end_date"`
	Tags        []string   `json:"tags"`
	Description string     `json:"description"`
}
