package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"rs/auth/app/configs"
	"rs/auth/app/middlewares"
	"rs/auth/app/routes"
	"strings"
)

func main() {
	port := configs.Config.Port
	router := mux.NewRouter()

	routes.Init(router)

	router.Use(middlewares.Traffic)

	allowedOrigins := strings.Split(configs.Config.CORS_WISHLIST, ",")
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		// Debug: true,
	})

	handler := corsHandler.Handler(router)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
