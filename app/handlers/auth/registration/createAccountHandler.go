package registration

import (
	"context"
	"net/http"
	"rs/auth/app/db/repositories"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/response"
	"rs/auth/app/services/hash"
	"rs/auth/app/utils"
)

type CreateAccountHandler struct {
	handlers.BaseHandler
}

func (h *CreateAccountHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	payload := r.Context().Value("payload").(dto.RegisterRequestBody)
	utils.LoggerInstance.Debug("Create account chain.")
	userRepo := repositories.NewUserRepository()
	user, err := userRepo.CreateAccount(models.User{
		Username: payload.Username,
		Password: hash.Instance.GenerateHash(payload.Password),
		Email:    payload.Email,
	})
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		response.Respond(w, http.StatusUnauthorized, err.Error(), nil)
		return false
	}

	if user == nil {
		utils.LoggerInstance.Info("Failed to create user account.")
		response.Respond(w, http.StatusInternalServerError, "User registration fail.", nil)
		return false
	}

	ctx := context.WithValue(r.Context(), "user", user)
	r = r.WithContext(ctx)
	return h.HandleNext(w, r)
}
