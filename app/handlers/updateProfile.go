package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rs/auth/app/models"
	"rs/auth/app/repositories/userProfile"
	"rs/auth/app/utils"
)

func InsertOrUpdateUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var profile models.UserProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if profile.UserID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userProfileRepo := userProfile.NewRepository()
	err := userProfileRepo.InsertOrUpdateUserProfile(profile)
	if err != nil {
		fmt.Println(err)
		utils.LoggerInstance.Error(err.Error())
		http.Error(w, "Failed to save profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "Profile saved successfully"})
}

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userProfileRepo := userProfile.NewRepository()
	result, err := userProfileRepo.GetUserProfile("4a0755bb-c107-4427-89db-fa584dae1479")
	if err != nil {
		fmt.Println(err)
		utils.LoggerInstance.Error(err.Error())
		http.Error(w, "Failed to save profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}
