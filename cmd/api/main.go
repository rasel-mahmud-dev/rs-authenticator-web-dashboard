package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rs/auth/configs"
	"rs/auth/internal/app/routes"
)

func main() {
	port := configs.ConfigInstance().Port
	router := mux.NewRouter()

	routes.Init(router)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
