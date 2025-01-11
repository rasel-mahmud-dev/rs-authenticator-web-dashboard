package generate2FASecret

import (
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net/statusCode"
	"rs/auth/app/repositories"
	"rs/auth/app/response"
	"rs/auth/app/utils"
)

type CheckInitTokenHandler struct {
	handlers.BaseHandler
}

func (h *CheckInitTokenHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	authSession := (*r).Context().Value("authSession").(*models.AuthSession)

	initToken, err := repositories.MfaSecurityTokenRepo.GetLastInit(authSession.UserId)
	if err != nil {
		utils.LoggerInstance.Error(err.Error())
	}

	if initToken != nil {
		utils.LoggerInstance.Error("Already  exists previous token.")
		response.Respond(w, statusCode.OK, "Success", initToken)
		return false
	}

	return h.HandleNext(w, r)
}
