package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"fyi/internal/models"
)

func Experience(w http.ResponseWriter, r *http.Request) {
	// Helper function to parse dates
	parseDate := func(dateStr string) *time.Time {
		if dateStr == "" {
			return nil
		}
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return nil
		}
		return &t
	}

	experiences := []models.Experience{
		{
			Title:       "Senior Remote Software Engineer",
			Company:     "Zoomph",
			Location:    "Remote, Virginia",
			StartDate:   parseDate("2015-06-01"),
			Active:      true,
			EndDate:     nil,
			Tags:        []string{"salaried"},
			Description: "I am a senior remote software engineer at Zoomph, where I work with the frontend team to build and maintain their flagship product. I also work with integrating frontend features into the backend, so I work full-stack.",
		},
		{
			Title:       "Hot Yoga Instructor",
			Company:     "Corepower Yoga",
			Location:    "Northern Virginia",
			StartDate:   nil, // TODO: Add start date
			Active:      true,
			EndDate:     nil,
			Tags:        []string{"hourly", "social"},
			Description: "I am an internationally certified hot yoga instructor, and I teach twice-weekly classes in 100+ degree rooms.",
		},
		{
			Title:       "Social Outreach Attendant",
			Company:     "CDCN",
			Location:    "Northern Virginia",
			StartDate:   parseDate("2008-01-01"),
			Active:      false,
			EndDate:     parseDate("2012-01-01"),
			Tags:        []string{"hourly", "social"},
			Description: "",
		},
		{
			Title:       "Barista",
			Company:     "Borjo Coffee",
			Location:    "Norfolk, Virginia",
			StartDate:   nil, // TODO: Add start date
			Active:      false,
			EndDate:     parseDate("2015-05-01"),
			Tags:        []string{"hourly", "social", "service"},
			Description: "I made coffee, made food, and helped organize concerts. The coffee shop was a popular spot for locals so I have fond memories of working there.",
		},
		{
			Title:       "Sales Associate",
			Company:     "Autozone",
			Location:    "Warrenton, Virginia",
			StartDate:   parseDate("2011-01-01"),
			Active:      false,
			EndDate:     parseDate("2011-04-01"),
			Tags:        []string{"hourly", "social", "retail", "commissioned"},
			Description: "",
		},
		{
			Title:       "Sales Associate",
			Company:     "Play N' Trade Videogames",
			Location:    "Warrenton, Virginia",
			StartDate:   parseDate("2010-06-01"),
			Active:      false,
			EndDate:     parseDate("2010-11-01"),
			Tags:        []string{"hourly", "social", "retail", "commissioned"},
			Description: "",
		},
		{
			Title:       "Barista",
			Company:     "Starbucks Coffee",
			Location:    "Warrenton, Virginia",
			StartDate:   parseDate("2007-01-01"),
			Active:      false,
			EndDate:     parseDate("2008-01-01"),
			Tags:        []string{"hourly", "service"},
			Description: "",
		},
		{
			Title:       "Associate",
			Company:     "Panera Bread",
			Location:    "Warrenton, Virginia",
			StartDate:   parseDate("2006-01-01"),
			Active:      false,
			EndDate:     parseDate("2008-01-01"),
			Tags:        []string{"hourly", "service"},
			Description: "",
		},
	}
	json.NewEncoder(w).Encode(experiences)
}
