package verify

import (
	context2 "rs/auth/app/context"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/services/jwt"
)

type ValidateAccessTokenHandler struct {
	handlers.BaseHandler
}

func (h *ValidateAccessTokenHandler) Handle(c context2.BaseContext) bool {
	token := c.AccessToken
	parseToken, err := jwt.Jwt.ParseToken(token)
	if err != nil {
		response.Respond(c.ResponseWriter, statusCode.INTERNAL_SERVER_ERROR, err.Error(), nil)
		return false
	}

	authSession := repositories.AuthSessionRepository.GetAuthSessionByAccessToken(token)
	if authSession == nil {
		response.Respond(c.ResponseWriter, statusCode.UNAUTHORIZED, "Unauthorized", nil)
		return false
	}

	if authSession.UserId != parseToken.UserId {
		response.Respond(c.ResponseWriter, statusCode.UNAUTHORIZED_SESSION_MISMATCH, "Session mismatch", nil)
		return false
	}

	if authSession.IsRevoked {
		response.Respond(c.ResponseWriter, statusCode.UNAUTHORIZED_SESSION_REVOKED, "Session revoked", dto.AuthVerify{
			ID:        authSession.UserId,
			SessionId: authSession.ID,
			IsRevoked: authSession.IsRevoked,
			Username:  authSession.Username,
			Email:     authSession.Email,
			Avatar:    authSession.Avatar,
		})
		return false
	}

	c.AuthSession = authSession

	return h.HandleNext(c)
}
