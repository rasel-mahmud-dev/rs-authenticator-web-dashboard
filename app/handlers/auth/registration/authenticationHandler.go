package registration

import (
	"fmt"
	"net/http"
	"rs/auth/app/handlers"
	"rs/auth/app/models"
	"rs/auth/app/net"
)

type AuthenticationHandler struct {
	handlers.BaseHandler
}

func (h *AuthenticationHandler) Handle(w http.ResponseWriter, r *http.Request) bool {
	user := r.Context().Value("user").(*models.User)
	user.Password = ""
	credentials := net.Success.LOGOUT_SUCCESS
	fmt.Println(credentials)
	//response.Respond(w, status.Success.LOGOUT_SUCCESS, "Account successful created", user)
	return false
}
