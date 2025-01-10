package verify

import (
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type ResponseHandler struct {
	handlers.BaseHandler
}

func (h *ResponseHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	authSession := r.Context().Value("authSession").(*models.AuthSession)
	response.Respond(w, statusCode.OK, "Ok", dto.AuthVerify{
		ID:        authSession.UserId,
		SessionId: authSession.ID,
		IsRevoked: authSession.IsRevoked,
		Username:  authSession.Username,
		Email:     authSession.Email,
	})
	return false
}
