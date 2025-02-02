package login

import (
	"fmt"
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/services/jwt"
	"rs/auth/app/utils"
	"time"
)

type GenerateJwtHandler struct {
	handlers.BaseHandler
}

func (h *GenerateJwtHandler) Handle(c *context2.BaseContext) bool {
	user := c.User
	user.Password = ""
	token, err := jwt.Jwt.GenerateToken(jwt.JwtPayload{UserId: user.ID}, time.Hour*24*60)
	if err != nil {
		utils.LoggerInstance.Error(fmt.Sprintf("Jwt token generation error - %s", err.Error()))
		response.Respond(c.ResponseWriter, statusCode.INTERNAL_SERVER_ERROR, "Internal error", nil)
		return false
	}
	c.AccessToken = token
	return h.HandleNext(c)
}
