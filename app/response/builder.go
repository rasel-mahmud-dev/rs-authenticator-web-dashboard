package response

import (
	"encoding/json"
	"net/http"
	"rs/auth/app/net/statusCode"
)

type APIResponse struct {
	Message    string      `json:"message"`
	StatusCode string      `json:"statusCode"`
	Data       interface{} `json:"data,omitempty"`
}

func Respond(w http.ResponseWriter, statusCodeT statusCode.StatusCodeType, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	httpStatusCode, exists := statusCode.HttpStatus[statusCodeT]
	if !exists {
		httpStatusCode = 200
	}
	w.WriteHeader(httpStatusCode)

	response := APIResponse{
		Message:    message,
		StatusCode: string(statusCodeT),
		Data:       data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
