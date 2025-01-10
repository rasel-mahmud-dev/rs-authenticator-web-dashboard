package dto

type AuthVerify struct {
	ID        string `json:"id,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
	IsRevoked bool   `json:"isRevoked"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Avatar    string `json:"avatar"`
}
