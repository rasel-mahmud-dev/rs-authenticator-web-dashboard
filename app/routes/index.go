package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/app/handlers"
	//"rs/auth/internal/app/handlers"
)

func Init(router *mux.Router) {
	router.HandleFunc("/api/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/api/home", handlers.HomeHandler).Methods("GET")
	//router.HandleFunc("/api/configuration", handlers.ConfigurationHandler).Methods("GET")
	//router.HandleFunc("/api/health", handlers.HealthHandler).Methods("GET")

	AuthRoutes(router)
}
