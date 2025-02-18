package accountRecovery

import (
	"encoding/json"
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/response"
)

type PreparedContextState struct {
	handlers.BaseHandler
}

func (h *PreparedContextState) Handle(c *context.BaseContext) bool {

	body := make(map[string]string)
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		response.Respond(c.ResponseWriter, statusCode.INVALID_JSON_FORMAT, "Invalid JSON format", nil)
		return false
	}

	c.TwoFaSecurityContext.AccountRecoveryCode = body["code"]
	return h.HandleNext(c)
}
