package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/app/handlers"
)

func StatisticsRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/slats/registration", handlers.RegistrationSlatsHandler).Methods("GET")
	router.HandleFunc("/api/v1/slats/authentication", handlers.AuthenticationSlatsHandler).Methods("GET")
	router.HandleFunc("/api/v1/slats/auth-attempts", handlers.LoginAttemptSlatsHandler).Methods("GET")
}
