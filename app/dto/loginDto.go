package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}

type AuthenticatorLoginRequestBody struct {
	OtpCode string `json:"otpCode" validate:"required,min=6,max=6"`
}
