package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/app/handlers/auth/login"
	"rs/auth/app/handlers/auth/loginWithAuthenticator"
	"rs/auth/app/handlers/auth/registration"
	"rs/auth/app/handlers/auth/verify"
	"rs/auth/app/utils"
)

func AuthRoutes(router *mux.Router) {
	utils.LoggerInstance.Info("Initializing routes")
	router.HandleFunc("/api/v1/login", login.LoginHandler).Methods("POST")
	router.HandleFunc("/api/v1/registration", registration.RegistrationHandler).Methods("POST")
	router.HandleFunc("/api/v1/verify", verify.AuthVerifyHandler).Methods("GET")
	router.HandleFunc("/api/v1/login-with-authenticator", loginWithAuthenticator.LoginWithAuthenticator).Methods("POST")
}
