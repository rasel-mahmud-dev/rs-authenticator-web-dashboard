package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func Init(router *mux.Router) {
	router.HandleFunc("/health", healthHandler).Methods("GET")

	AuthRoutes(router)
}
