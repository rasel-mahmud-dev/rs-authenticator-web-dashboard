package registration

import (
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type CheckExistenceUserHandler struct {
	handlers.BaseHandler
}

func (h *CheckExistenceUserHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	payload := r.Context().Value("payload").(dto.RegisterRequestBody)
	userRepo := repositories.NewUserRepository()
	user, err := userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		response.Respond(w, statusCode.INTERNAL_ERROR, "Internal error", nil)
		return false
	}

	if user != nil {
		utils.LoggerInstance.Info("User already exist in database.")
		response.Respond(w, statusCode.DUPLICATE_ENTITY, "User already onboarded", nil)
		return false
	}

	return h.HandleNext(w, r)
}
