package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/app/handlers"
	"rs/auth/app/middlewares"
)

func StatisticsRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/slats/registration", middlewares.Auth(handlers.RegistrationSlatsHandler)).Methods("GET")
	router.HandleFunc("/api/v1/slats/authentication", middlewares.Auth(handlers.AuthenticationSlatsHandler)).Methods("GET")
	router.HandleFunc("/api/v1/slats/auth-attempts", middlewares.Auth(handlers.LoginAttemptSlatsHandler)).Methods("GET")
	router.HandleFunc("/api/v1/slats/traffic", middlewares.Auth(handlers.FetchTrafficStats)).Methods("GET")
	router.HandleFunc("/api/v1/slats/api-latency", middlewares.Auth(handlers.GetApiLatencyStats)).Methods("GET")
	router.HandleFunc("/api/v1/slats/users", middlewares.Auth(handlers.FetchUsers)).Methods("GET")
}
