package routes

import (
	"github.com/gorilla/mux"
	authHandler "rs/auth/internal/app/handlers/auth"
	"rs/auth/internal/utils"
)

func AuthRoutes(router *mux.Router) {
	utils.LoggerInstance.Info("Initializing routes")
	router.HandleFunc("/login", authHandler.LoginHandler).Methods("POST")
}
