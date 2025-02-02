package registration

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type AuthenticationHandler struct {
	handlers.BaseHandler
}

func (h *AuthenticationHandler) Handle(c context.BaseContext) bool {
	user := c.User
	user.Password = ""
	response.Respond(c.ResponseWriter, statusCode.ACCOUNT_CREATED, "Account successful created", user)
	return false
}
