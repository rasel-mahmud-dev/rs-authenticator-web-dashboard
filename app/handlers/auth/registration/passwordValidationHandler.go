package registration

import (
	"fmt"
	"net/http"
	"rs/auth/app/dto"
	"rs/auth/app/handlers"
	"rs/auth/app/utils"
)

type PasswordValidationHandler struct {
	handlers.BaseHandler
}

func (h *PasswordValidationHandler) Handle(w http.ResponseWriter, r **http.Request) bool {
	registerRequestBody := (*r).Context().Value("payload").(dto.RegisterRequestBody)
	utils.LoggerInstance.Info("Check password strange.")
	fmt.Println(registerRequestBody)
	return h.HandleNext(w, r)
}
