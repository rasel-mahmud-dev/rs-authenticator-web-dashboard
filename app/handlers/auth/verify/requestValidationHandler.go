package verify

import (
	"context"
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type RequestValidationHandler struct {
	handlers.BaseHandler
}

func (h *RequestValidationHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	token := utils.GetToken(r)
	if token == "" {
		response.Respond(w, statusCode.ACCESS_TOKEN_MISSED, "Access required.", nil)
		return false
	}
	ctx := context.WithValue(r.Context(), "accessToken", token)
	r = r.WithContext(ctx)
	return h.HandleNext(w, r)
}
