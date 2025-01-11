package dto

type Completed2FASecretBody struct {
	Id          string `json:"id" validate:"required"`
	AppName     string `json:"provider" validate:"required"`
	IsCompleted bool   `json:"isCompleted" validate:"required"`
}
