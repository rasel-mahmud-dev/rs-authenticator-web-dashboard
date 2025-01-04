package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rs/auth/internal/app/config"
	"rs/auth/internal/app/routes"
)

func main() {
	port := config.ConfigInstance().Port
	router := mux.NewRouter()

	routes.Init(router)

	//router.HandleFunc("/", homeHandler).Methods("GET")
	//router.HandleFunc("/about", aboutHandler).Methods("GET")

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
