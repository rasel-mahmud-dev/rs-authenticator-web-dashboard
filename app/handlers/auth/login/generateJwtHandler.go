package login

import (
	"context"
	"fmt"
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/services/jwt"
	"rs/auth/app/utils"
	"time"
)

type GenerateJwtHandler struct {
	handlers.BaseHandler
}

func (h *GenerateJwtHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	user := (*r).Context().Value("user").(*models.User)
	user.Password = ""
	token, err := jwt.Jwt.GenerateToken(jwt.JwtPayload{UserId: user.ID}, time.Hour*24*60)
	if err != nil {
		utils.LoggerInstance.Error(fmt.Sprintf("Jwt token generation error - %s", err.Error()))
		response.Respond(w, statusCode.INTERNAL_SERVER_ERROR, "Internal error", nil)
		return false
	}
	ctx := context.WithValue((*r).Context(), "token", token)
	*r = (*r).WithContext(ctx)
	return h.HandleNext(w, r)
}
