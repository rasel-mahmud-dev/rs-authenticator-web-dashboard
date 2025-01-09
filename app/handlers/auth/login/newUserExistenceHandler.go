package login

import (
	"context"
	"net/http"
	"rs/auth/app/db/repositories"
	"rs/auth/app/dto"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type UserExistenceHandler struct {
	BaseHandler
	repo repositories.UserRepository
}

func NewUserExistenceHandler(repo repositories.UserRepository) *UserExistenceHandler {
	return &UserExistenceHandler{repo: repo}
}

func (h *UserExistenceHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	loginRequest := r.Context().Value("loginRequest").(dto.LoginRequest)
	user, err := h.repo.GetUserByEmail(loginRequest.Email)

	if err != nil || user == nil {
		utils.LoggerInstance.Info("User does not exist in database.")
		response.Respond(w, http.StatusUnauthorized, "Invalid email or password", nil)
		return false
	}

	ctx := context.WithValue(r.Context(), "user", user)
	r = r.WithContext(ctx)
	
	return h.HandleNext(w, r)
}
