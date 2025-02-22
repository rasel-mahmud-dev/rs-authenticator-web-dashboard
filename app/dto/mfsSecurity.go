package dto

type Completed2FASecretBody struct {
	CodeName  string `json:"codeName"`
	Secret    string `json:"secret" validate:"required"`
	QrCodeURL string `json:"qrCodeURL"`
}

type GenerateMfaQRRequestPayload struct {
	Provider string `json:"provider"`
	IsNew    bool   `json:"isNew"`
}
