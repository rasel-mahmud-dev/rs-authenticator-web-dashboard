package response

import (
	"encoding/json"
	"net/http"
	"rs/auth/app/net"
)

type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Respond(w http.ResponseWriter, statusCode string, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	httpStatusCode := net.HttpStatus[statusCode]
	w.WriteHeader(httpStatusCode)

	response := APIResponse{
		Message: message,
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
