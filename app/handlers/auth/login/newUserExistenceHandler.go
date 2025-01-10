package login

import (
	"context"
	"net/http"
	"rs/auth/app/db/repositories"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type UserExistenceHandler struct {
	handlers.BaseHandler
}

func (h *UserExistenceHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	loginRequest := r.Context().Value("loginRequest").(dto.LoginRequest)
	userRepo := repositories.NewUserRepository()
	user, err := userRepo.GetUserByEmail(loginRequest.Email)
	if err != nil || user == nil {
		utils.LoggerInstance.Info("User does not exist in database.")
		response.Respond(w, statusCode.INVALID_CREDENTIALS, "Invalid email or password", nil)
		return false
	}

	ctx := context.WithValue(r.Context(), "user", user)
	r = r.WithContext(ctx)

	return h.HandleNext(w, r)
}
