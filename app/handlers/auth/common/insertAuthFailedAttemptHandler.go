package common

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/repositories"
)

type InsertAuthFailedAttemptHandler struct {
	handlers.BaseHandler
}

func (h *InsertAuthFailedAttemptHandler) Handle(c *context.BaseContext) bool {
	repositories.AuthSessionRepository.InsertAuthFailedAttempt(c.LoginContext.UserAuthAttempt)
	return h.HandleNext(c)
}
