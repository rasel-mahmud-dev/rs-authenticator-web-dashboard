package dto

type UserRegistrationStats struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type AuthenticatorStats struct {
	Date          string `json:"date"`
	Authenticator int    `json:"authenticator"`
	Password      int    `json:"password"`
}

type GetAttemptRateStatsResult struct {
	Failed  int `json:"failed"`
	Success int `json:"success"`
	Total   int `json:"total"`
}
