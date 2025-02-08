package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/app/handlers"
)

func UserProfileRoutes(router *mux.Router) {
	//router.HandleFunc("/api/configuration", handlers.ConfigurationHandler).Methods("GET")
	router.HandleFunc("/api/v1/profile", handlers.InsertOrUpdateUserProfileHandler).Methods("POST")
	router.HandleFunc("/api/v1/profile", handlers.GetUserProfileHandler).Methods("GET")
}
