package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rs/auth/configs"
	"rs/auth/internal/app/routes"
	"rs/auth/internal/db/repositories"
)

func main() {
	port := configs.ConfigInstance().Port
	router := mux.NewRouter()
	repositories.NewUserRepository()

	routes.Init(router)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
