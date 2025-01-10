package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rs/auth/app/configs"
	"rs/auth/app/db/repositories"
	"rs/auth/app/routes"
)

func main() {
	port := configs.Config.Port
	router := mux.NewRouter()
	repositories.NewUserRepository()

	routes.Init(router)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
