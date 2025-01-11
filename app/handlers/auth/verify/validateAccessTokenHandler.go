package verify

import (
	"context"
	"fmt"
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/services/jwt"
)

type ValidateAccessTokenHandler struct {
	handlers.BaseHandler
}

func (h *ValidateAccessTokenHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	token := (*r).Context().Value("accessToken").(string)
	parseToken, err := jwt.Jwt.ParseToken(token)
	if err != nil {
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, err.Error(), nil)
		return false
	}

	authSession := repositories.AuthSessionRepository.GetAuthSessionByAccessToken(token)
	if authSession == nil {
		response.Respond(w, statusCode.UNAUTHORIZED, "Unauthorized", nil)
		return false
	}

	if authSession.UserId != parseToken.UserId {
		response.Respond(w, statusCode.UNAUTHORIZED_SESSION_MISMATCH, "Session mismatch", nil)
		return false
	}

	if authSession.IsRevoked {
		response.Respond(w, statusCode.UNAUTHORIZED_SESSION_REVOKED, "Session revoked", dto.AuthVerify{
			ID:        authSession.UserId,
			SessionId: authSession.ID,
			IsRevoked: authSession.IsRevoked,
			Username:  authSession.Username,
			Email:     authSession.Email,
			Avatar:    authSession.Avatar,
		})
		return false
	}

	ctx := context.WithValue((*r).Context(), "authSession", authSession)
	*r = (*r).WithContext(ctx)

	authSession2 := (*r).Context().Value("authSession").(*models.AuthSession)
	fmt.Println("authSession2", authSession2)

	return h.HandleNext(w, r)
}
