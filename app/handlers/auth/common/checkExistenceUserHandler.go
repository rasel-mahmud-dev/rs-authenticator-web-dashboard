package common

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type CheckExistenceUserHandler struct {
	handlers.BaseHandler
}

func (h *CheckExistenceUserHandler) Handle(c context.BaseContext) bool {

	userRepo := repositories.NewUserRepository()
	user, err := userRepo.GetUserByEmail(c.Email)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		response.Respond(c.ResponseWriter, statusCode.INTERNAL_SERVER_ERROR, "Internal error", nil)
		return false
	}

	if user != nil {
		utils.LoggerInstance.Info("User already exist in database.")
		response.Respond(c.ResponseWriter, statusCode.DUPLICATE_ENTITY, "User already onboarded", nil)
		return false
	}

	return h.HandleNext(c)
}
