package routes

import (
	"github.com/gorilla/mux"
	"rs/auth/app/handlers"
	"rs/auth/app/middlewares"
)

func UserProfileRoutes(router *mux.Router) {
	//router.HandleFunc("/api/configuration", handlers.ConfigurationHandler).Methods("GET")
	router.HandleFunc("/api/v1/profile", middlewares.Auth(handlers.InsertOrUpdateUserProfileHandler)).Methods("PUT")
	router.HandleFunc("/api/v1/profile", middlewares.Auth(handlers.GetUserProfileHandler)).Methods("GET")
	router.HandleFunc("/api/v1/profile/avatar", middlewares.Auth(handlers.UpdateProfileAvatarHandler)).Methods("POST")
}
