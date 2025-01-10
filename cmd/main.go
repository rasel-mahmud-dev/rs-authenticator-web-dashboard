package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"rs/auth/app/configs"
	"rs/auth/app/routes"
)

func main() {
	port := configs.Config.Port
	router := mux.NewRouter()
	routes.Init(router)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3200"},                             // Allow all origins, or specify your frontend domain here
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},           // Allow specific methods
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"}, // Allowed headers
		AllowCredentials: true,
		// Debug: true,
	})

	handler := corsHandler.Handler(router)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
