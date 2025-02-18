package dto

import "time"

type UserRegistrationStats struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type AuthenticatorStats struct {
	Date          string `json:"date"`
	Authenticator int    `json:"authenticator"`
	Password      int    `json:"password"`
	RecoveryCode  int    `json:"recovery_code"`
}

type GetAttemptRateStatsResult struct {
	Failed  int `json:"failed"`
	Success int `json:"success"`
	Total   int `json:"total"`
}

type AttemptRateStatsDetail struct {
	Date    time.Time `json:"date"`
	Failed  int       `json:"failed"`
	Success int       `json:"success"`
}

type DetailedTrafficStats struct {
	RoutePath       string    `json:"route_path,omitempty"`
	Count           int       `json:"request_count,omitempty"`
	AvgResponseTime float64   `json:"avg_response_time,omitempty"`
	LastAccessed    time.Time `json:"last_accessed,omitempty"`
	RequestTime     time.Time `json:"date,omitempty"`
}
type TrafficCountStats struct {
	Count       int       `json:"request_count"`
	RequestTime time.Time `json:"date"`
}
