package dto

type RegisterRequestBody struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=3,max=100"`
}
