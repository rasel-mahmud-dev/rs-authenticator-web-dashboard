package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories/userProfile"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

func InsertOrUpdateUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authSession := (*r).Context().Value("authSession").(*models.AuthSession)

	if authSession == nil {
		response.Respond(w, statusCode.UNAUTHORIZED, "UNAUTHORIZED", nil)
		return
	}

	var profile dto.UpdateProfilePayload
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		utils.LoggerInstance.Error(err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	profile.UserID = authSession.UserId

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

	authSession := (*r).Context().Value("authSession").(*models.AuthSession)

	if authSession == nil {
		response.Respond(w, statusCode.UNAUTHORIZED, "UNAUTHORIZED", nil)
		return
	}

	userProfileRepo := userProfile.NewRepository()
	result, err := userProfileRepo.GetUserProfile(authSession.UserId)
	if err != nil {
		fmt.Println(err)
		utils.LoggerInstance.Error(err.Error())
		http.Error(w, "Failed to save profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}
