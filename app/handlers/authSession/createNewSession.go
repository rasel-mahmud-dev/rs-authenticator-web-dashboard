package authSession

import (
	"context"
	"fmt"
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/repositories"
	"rs/auth/app/utils"
)

type NewSessionHandler struct {
	handlers.BaseHandler
}

func (h *NewSessionHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	user := (*r).Context().Value("user").(*models.User)
	accessToken := (*r).Context().Value("token").(string)

	authSession, err := repositories.AuthSessionRepository.InsertAuthSession(models.AuthSession{
		UserId:       user.ID,
		IPAddress:    utils.GetUserIP(*r),
		UserAgent:    utils.GetUserAgent(*r),
		AccessToken:  accessToken,
		RefreshToken: accessToken,
	})
	if err != nil {
		utils.LoggerInstance.Error(fmt.Sprintf("Auth session creation failed %s", err))
	}

	ctx := context.WithValue((*r).Context(), "authSession", authSession)
	*r = (*r).WithContext(ctx)

	return h.HandleNext(w, r)
}
