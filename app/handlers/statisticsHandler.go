package handlers

import (
	"encoding/json"
	"net/http"
	"rs/auth/app/repositories"
	"rs/auth/app/repositories/trafficRepo"
	"strconv"
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

func LoginAttemptSlatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	value := r.URL.Query().Get("t")
	userRepo := repositories.NewUserRepository()
	var stats interface{}
	if value == "detail" {
		stats, _ = userRepo.GetAttemptRateDetailStats()

	} else {
		stats = userRepo.GetAttemptRateStats()
	}

	_ = json.NewEncoder(w).Encode(stats)
}

func GetApiLatencyStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	trafficRepository := trafficRepo.TrafficRepository
	stats, _ := trafficRepository.GetApiLatencyStats()
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"cachedTime": nil,
		"data":       stats,
	})
}

func FetchTrafficStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	value := r.URL.Query().Get("t")
	trafficRepository := trafficRepo.TrafficRepository
	var stats interface{}
	if value == "detail" {
		stats, _ = trafficRepository.GetTrafficDetailStats()
	} else {
		stats, _ = trafficRepository.GetTrafficCountStats()
	}
	_ = json.NewEncoder(w).Encode(stats)
}

func FetchUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err != nil || parsedPage <= 0 {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
		page = parsedPage
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 10
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit <= 0 {
			http.Error(w, "Invalid limit number", http.StatusBadRequest)
			return
		}
		limit = parsedLimit
	}

	userRepo := repositories.NewUserRepository()
	users, totalItems, err := userRepo.GetAllUsers(page, limit)
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{"data": users, "totalItems": totalItems})
}
