package registration

import (
	"rs/auth/app/context"
	"rs/auth/app/handlers"
	"rs/auth/app/utils"
)

type PasswordValidationHandler struct {
	handlers.BaseHandler
}

func (h *PasswordValidationHandler) Handle(c context.BaseContext) bool {
	//registerRequestBody := (*r).Context().Value("payload").(dto.RegisterRequestBody)
	utils.LoggerInstance.Info("Check password strange.")
	//fmt.Println(registerRequestBody)
	return h.HandleNext(c)
}
