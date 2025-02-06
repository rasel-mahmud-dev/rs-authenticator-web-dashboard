package dto

type Completed2FASecretBody struct {
	Id          string `json:"id" validate:"required"`
	AppName     string `json:"provider" validate:"required"`
	IsCompleted bool   `json:"isCompleted" validate:"required"`
}

type GenerateMfaQRRequestPayload struct {
	Provider string `json:"provider"`
	IsNew    bool   `json:"isNew"`
}
