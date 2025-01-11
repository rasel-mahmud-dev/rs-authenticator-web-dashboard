package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/app/handlers"
	"rs/auth/app/handlers/twoFactor"
	"rs/auth/app/middlewares"
	//"rs/auth/internal/app/handlers"
)

func Init(router *mux.Router) {
	router.HandleFunc("/api/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/api/home", handlers.HomeHandler).Methods("GET")
	//router.HandleFunc("/api/configuration", handlers.ConfigurationHandler).Methods("GET")
	router.HandleFunc("/api/health", handlers.HealthHandler).Methods("GET")
	router.HandleFunc("/api/v1/generate-2fa-secret", middlewares.Auth(twoFactor.Generate2FASecret)).Methods("GET")
	router.HandleFunc("/api/v1/generate-2fa-secret", middlewares.Auth(twoFactor.Finalize2FASecret)).Methods("POST")

	AuthRoutes(router)
}
