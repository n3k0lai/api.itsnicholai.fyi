package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIInfo struct {
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	Description string     `json:"description"`
	Endpoints   []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Path        string `json:"path"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

type AboutInfo struct {
	Name      string   `json:"name"`
	Bio       string   `json:"bio"`
	Skills    []string `json:"skills"`
	Interests []string `json:"interests"`
}

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

type Project struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Technologies []string `json:"technologies"`
	URL          string   `json:"url"`
	Github       string   `json:"github"`
}

type ContactInfo struct {
	Email     string `json:"email"`
	Github    string `json:"github"`
	Twitter   string `json:"twitter"`
	Twitch    string `json:"twitch"`
	Spotify   string `json:"spotify"`
	LinkedIn  string `json:"linkedin"`
	Instagram string `json:"instagram"`
}

func index(w http.ResponseWriter, r *http.Request) {
	info := APIInfo{
		Name:        "Nicholai's Personal API",
		Version:     "1.0.0",
		Description: "A personal API for Nicholai",
		Endpoints: []Endpoint{
			{Path: "/", Method: "GET", Description: "API information"},
			{Path: "/about", Method: "GET", Description: "About Nicholai"},
			{Path: "/projects", Method: "GET", Description: "List of projects"},
			{Path: "/contact", Method: "GET", Description: "Contact information"},
		},
	}
	json.NewEncoder(w).Encode(info)
}

func about(w http.ResponseWriter, r *http.Request) {
	info := AboutInfo{
		Name:      "Nicholai",
		Bio:       "Senior Remote Software Engineer",
		Skills:    []string{"Skill 1", "Skill 2", "Skill 3"},
		Interests: []string{"Yoga", "Livestreaming", "Social Outreach"},
	}
	json.NewEncoder(w).Encode(info)
}

func experience(w http.ResponseWriter, r *http.Request) {
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

	experiences := []Experience{
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
			StartDate:   nil, // No start date provided
			Active:      true,
			EndDate:     nil,
			Tags:        []string{"hourly"},
			Description: "I am an internationally certified hot yoga instructor, and I teach twice-weekly classes in 100+ degree rooms.",
		},
		{
			Title:       "Social Outreach Coordinator",
			Company:     "CDCN",
			Location:    "Northern Virginia",
			StartDate:   parseDate("2008-01-01"),
			Active:      false,
			EndDate:     parseDate("2012-01-01"),
			Tags:        []string{"hourly"},
			Description: "",
		},
		{
			Title:       "Barista",
			Company:     "Borjo Coffee",
			Location:    "Norfolk, Virginia",
			StartDate:   nil, // No start date provided
			Active:      false,
			EndDate:     parseDate("2015-05-01"),
			Tags:        []string{"hourly"},
			Description: "I made coffee, made food, and helped organize concerts. The coffee shop was a popular spot for locals so I have fond memories of working there.",
		},
		{
			Title:       "Sales Associate",
			Company:     "Autozone",
			Location:    "Warrenton, Virginia",
			StartDate:   parseDate("2011-01-01"),
			Active:      false,
			EndDate:     parseDate("2011-04-01"),
			Tags:        []string{"hourly"},
			Description: "",
		},
		{
			Title:       "Sales Associate",
			Company:     "Play N' Trade Videogames",
			Location:    "Warrenton, Virginia",
			StartDate:   parseDate("2010-06-01"),
			Active:      false,
			EndDate:     parseDate("2010-11-01"),
			Tags:        []string{"hourly"},
			Description: "",
		},
		{
			Title:       "Barista",
			Company:     "Starbucks Coffee",
			Location:    "Warrenton, Virginia",
			StartDate:   parseDate("2007-01-01"),
			Active:      false,
			EndDate:     parseDate("2008-01-01"),
			Tags:        []string{"hourly"},
			Description: "",
		},
		{
			Title:       "Associate",
			Company:     "Panera Bread",
			Location:    "Warrenton, Virginia",
			StartDate:   parseDate("2006-01-01"),
			Active:      false,
			EndDate:     parseDate("2008-01-01"),
			Tags:        []string{"hourly"},
			Description: "",
		},
	}
	json.NewEncoder(w).Encode(experiences)
}

func projects(w http.ResponseWriter, r *http.Request) {
	projects := []Project{
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

func contact(w http.ResponseWriter, r *http.Request) {
	info := ContactInfo{
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

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error":               "Not found",
		"message":             "The requested URL was not found on this server.",
		"available_endpoints": []string{"/", "/about", "/projects", "/contact"},
	})
}

func serverError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   "Server error",
		"message": "An internal server error occurred.",
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/about", about).Methods("GET")
	r.HandleFunc("/experience", experience).Methods("GET")
	r.HandleFunc("/projects", projects).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(notFound)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "6000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}
