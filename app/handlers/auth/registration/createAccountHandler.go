package registration

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/services/hash"
	"rs/auth/app/utils"
)

type CreateAccountHandler struct {
	handlers.BaseHandler
}

func (h *CreateAccountHandler) Handle(c *context.BaseContext) bool {
	payload := c.RegistrationContext.Payload
	utils.LoggerInstance.Debug("Create account chain.")
	user, err := repositories.UserRepositoryInstance.CreateAccount(models.User{
		Username: payload.Username,
		Password: hash.Hash.GenerateHash(payload.Password),
		Email:    payload.Email,
	})
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
		response.Respond(c.ResponseWriter, statusCode.ACCOUNT_CREATION_FAILED, err.Error(), nil)
		return false
	}

	if user == nil {
		utils.LoggerInstance.Info("Failed to create user account.")
		response.Respond(c.ResponseWriter, statusCode.ACCOUNT_CREATION_FAILED, "Failed to create user account.", nil)
		return false
	}

	c.User = user
	return h.HandleNext(c)
}
