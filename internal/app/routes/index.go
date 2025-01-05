package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/internal/app/handlers"
)

func Init(router *mux.Router) {
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	AuthRoutes(router)
}
