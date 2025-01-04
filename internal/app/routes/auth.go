package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/internal/utils"
)

func AuthRoutes(router *mux.Router) {
	utils.LoggerInstance.Info("Initializing routes")
	//router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
}
