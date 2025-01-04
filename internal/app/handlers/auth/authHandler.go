package authHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rs/auth/internal/app/dto"
	"rs/auth/internal/app/validators"
	"rs/auth/internal/db/repositories"
	"rs/auth/internal/response"
	"rs/auth/internal/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, "Invalid JSON format", nil)
		return
	}

	err = validators.ValidateLoginRequest(&loginRequest)
	if err != nil {
		response.Respond(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	repo := repositories.NewUserRepository()
	user, err := repo.GetUserByEmail(loginRequest.Email)

	if user == nil {
		utils.LoggerInstance.Info(fmt.Sprintf("User %s no exists on database.", loginRequest.Email))
		response.Respond(w, http.StatusUnauthorized, "Invalid username or password", nil)
		return
	}

	if loginRequest.Email == "admin" && loginRequest.Password == "password" {
		response.Respond(w, http.StatusOK, "Login successful", nil)
	} else {
		response.Respond(w, http.StatusUnauthorized, "Invalid username or password", nil)
	}
}
