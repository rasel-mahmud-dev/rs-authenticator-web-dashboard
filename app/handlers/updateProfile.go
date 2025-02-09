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
	"time"
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

func UpdateProfileAvatarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get user session
	authSession, ok := r.Context().Value("authSession").(*models.AuthSession)
	if !ok || authSession == nil {
		response.Respond(w, statusCode.UNAUTHORIZED, "UNAUTHORIZED", nil)
		return
	}

	// Parse the uploaded file
	file, handler, err := r.FormFile("image")
	if err != nil {
		response.Respond(w, statusCode.BAD_REQUEST, "Invalid image", nil)
		return
	}
	defer file.Close()

	// Validate file type
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}
	if !allowedTypes[handler.Header.Get("Content-Type")] {
		response.Respond(w, statusCode.BAD_REQUEST, "Only JPEG, PNG, and WEBP are allowed", nil)
		return
	}

	// Create a unique filename
	filename := fmt.Sprintf("profile_%d", time.Now().UnixNano())

	// Upload the image to Cloudinary
	imageURL, err := utils.UploadToCloudinary(file, filename)
	if err != nil {
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, "Failed to upload image", nil)
		return
	}

	// Update the user profile with the new image URL
	var profile dto.UpdateProfilePayload
	profile.UserID = authSession.UserId
	profile.Avatar = imageURL
	profile.Cover = ""

	userProfileRepo := userProfile.NewRepository()
	err = userProfileRepo.InsertOrUpdateUserProfileAvatar(profile)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		http.Error(w, "Failed to upload profile image.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Profile updated successfully",
		"image":   imageURL,
	})
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
