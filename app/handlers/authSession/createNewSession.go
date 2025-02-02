package authSession

import (
	"fmt"
	context2 "rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/repositories"
	"rs/auth/app/utils"
)

type NewSessionHandler struct {
	handlers.BaseHandler
}

func (h *NewSessionHandler) Handle(c *context2.BaseContext) bool {
	user := c.User
	accessToken := c.AccessToken

	authSession, err := repositories.AuthSessionRepository.InsertAuthSession(models.AuthSession{
		UserId:       user.ID,
		IPAddress:    utils.GetUserIP(c.Request),
		UserAgent:    utils.GetUserAgent(c.Request),
		AccessToken:  accessToken,
		RefreshToken: accessToken,
	})
	if err != nil {
		utils.LoggerInstance.Error(fmt.Sprintf("Auth session creation failed %s", err))
	}

	c.AuthSession = authSession

	return h.HandleNext(c)
}
