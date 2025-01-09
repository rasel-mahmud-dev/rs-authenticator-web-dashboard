package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type HomeData struct {
	Title       string
	Description string
	Items       []string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/home.gohtml",
	)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := HomeData{
		Title:       "Railway Reservation System",
		Description: "Manage your railway bookings easily.",
		Items:       []string{"Book Tickets", "Cancel Tickets", "View Train Schedules"},
	}

	err = tmpl.ExecuteTemplate(w, "home", data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
