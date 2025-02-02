package verify

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type RequestValidationHandler struct {
	handlers.BaseHandler
}

func (h *RequestValidationHandler) Handle(c context2.BaseContext) bool {
	token := utils.GetToken(c.Request)
	if token == "" {
		response.Respond(c.ResponseWriter, statusCode.ACCESS_TOKEN_MISSED, "Access required.", nil)
		return false
	}

	c.AccessToken = token
	return h.HandleNext(c)
}
