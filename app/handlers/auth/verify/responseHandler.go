package verify

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type ResponseHandler struct {
	handlers.BaseHandler
}

func (h *ResponseHandler) Handle(c context2.BaseContext) bool {
	authSession := c.AuthSession
	response.Respond(c.ResponseWriter, statusCode.OK, "Ok", dto.AuthVerify{
		ID:        authSession.UserId,
		SessionId: authSession.ID,
		IsRevoked: authSession.IsRevoked,
		Username:  authSession.Username,
		Email:     authSession.Email,
		Avatar:    authSession.Avatar,
	})
	return false
}
