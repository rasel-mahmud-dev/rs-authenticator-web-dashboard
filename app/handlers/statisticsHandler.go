package handlers

import (
	"encoding/json"
	"net/http"
	"rs/auth/app/repositories"
)

func RegistrationSlatsHandler(w http.ResponseWriter, _r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userRepo := repositories.NewUserRepository()
	stats, err := userRepo.GetUserRegistrationStats()
	if err != nil {
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{"data": stats})
}

func AuthenticationSlatsHandler(w http.ResponseWriter, _r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userRepo := repositories.NewUserRepository()
	stats, err := userRepo.GetAuthenticationStats()
	if err != nil {
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{"data": stats})
}

func LoginAttemptSlatsHandler(w http.ResponseWriter, _r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userRepo := repositories.NewUserRepository()
	stats := userRepo.GetAttemptRateStats()
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"data": stats})
}
