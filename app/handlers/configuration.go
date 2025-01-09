package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type ConfigurationCategory struct {
	Label       string
	Description string
	Checked     bool
	Options     []ConfigurationOption
}

type Setting struct {
	Name    string
	Enabled bool
}

type ConfigurationOption struct {
	Label       string
	Description string
	Settings    []Setting
}

type ConfigurationData struct {
	Title       string
	Description string
	Items       []string
	Categories  []ConfigurationCategory
}

func ConfigurationHandler(w http.ResponseWriter, r *http.Request) {

	categories := []ConfigurationCategory{
		{
			Label:       "Contact information",
			Description: "Specify wheth er your users should have email addresses or phone numbers",
			Checked:     true,
			Options: []ConfigurationOption{
				{
					Label:       "Email address",
					Description: "Users can add email addresses to their account",
					Settings: []Setting{
						{Name: "Required", Enabled: true},
						{Name: "Used for sign-in", Enabled: true},
						{Name: "Verify at sign-up", Enabled: true},
						{Name: "Email verification link", Enabled: true},
						{Name: "Email verification code", Enabled: true},
					},
				},
				{
					Label:       "Phone number",
					Description: "Users can add phone numbers to their account",
					Settings:    nil,
				},
			},
		},
		{

			Label:       "Username",
			Description: "Specify whether your users have a unique username",
			Checked:     true,
			Options: []ConfigurationOption{
				{
					Label:       "Email address",
					Description: "Users can add email addresses to their account",
					Settings: []Setting{
						{Name: "Required", Enabled: true},
						{Name: "Used for sign-in", Enabled: true},
						{Name: "Verify at sign-up", Enabled: true},
						{Name: "Email verification link", Enabled: true},
						{Name: "Email verification code", Enabled: true},
					},
				},
				{
					Label:       "Phone number",
					Description: "Users can add phone numbers to their account",
					Settings:    nil,
				},
			},
		},
	}

	tmpl, err := template.New("base").ParseFiles(
		"templates/header.gohtml",
		"templates/head.gohtml",
		"templates/feat-card.gohtml",
		"templates/configuration.gohtml",
	)

	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := ConfigurationData{
		Title:       "Railway Reservation System",
		Description: "Manage your railway bookings easily.",
		Items:       []string{""},
		Categories:  categories,
	}

	err = tmpl.ExecuteTemplate(w, "configuration", data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
