package generate2FASecret

import (
	"encoding/json"
	"rs/auth/app/context"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type PreparedContextState struct {
	handlers.BaseHandler
}

func (h *PreparedContextState) Handle(c *context.BaseContext) bool {

	var payload dto.GenerateMfaQRRequestPayload
	err := json.NewDecoder((*c.Request).Body).Decode(&payload)
	if err != nil {
		response.Respond(c.ResponseWriter, statusCode.INVALID_JSON_FORMAT, "Invalid JSON format", nil)
		return false
	}

	c.TwoFaSecurityContext.GenerateMfaBody = payload
	authSession := c.Request.Context().Value("authSession").(*models.AuthSession)
	c.AuthSession = authSession
	return h.HandleNext(c)
}
