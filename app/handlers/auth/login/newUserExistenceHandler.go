package login

import (
	"context"
	"fmt"
	"net/http"
	"rs/auth/app/db/repositories"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type UserExistenceHandler struct {
	handlers.BaseHandler
	userRepo *repositories.UserRepository
}

func NewUserExistenceHandler() *UserExistenceHandler {
	userRepo := repositories.NewUserRepository()
	return &UserExistenceHandler{userRepo: userRepo}
}

func (h *UserExistenceHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	loginRequest := r.Context().Value("loginRequest").(dto.LoginRequest)
	user, err := h.userRepo.GetUserByEmail(loginRequest.Email)
	fmt.Println("sfsdf", user, nil)
	if err != nil || user == nil {
		utils.LoggerInstance.Info("User does not exist in database.")
		response.Respond(w, http.StatusUnauthorized, "Invalid email or password", nil)
		return false
	}

	ctx := context.WithValue(r.Context(), "user", user)
	r = r.WithContext(ctx)

	return h.HandleNext(w, r)
}
