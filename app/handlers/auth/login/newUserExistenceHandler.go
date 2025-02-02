package login

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type UserExistenceHandler struct {
	handlers.BaseHandler
}

func (h *UserExistenceHandler) Handle(c context2.BaseContext) bool {
	loginRequest := c.LoginContext.LoginRequest
	userRepo := repositories.NewUserRepository()
	user, err := userRepo.GetUserByEmail(loginRequest.Email)
	if err != nil || user == nil {
		utils.LoggerInstance.Info("User does not exist in database.")
		response.Respond(c.ResponseWriter, statusCode.INVALID_CREDENTIALS, "Invalid email or password", nil)
		return false
	}

	c.User = user
	return h.HandleNext(c)
}
