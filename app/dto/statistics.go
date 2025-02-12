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
	RoutePath       string    `json:"route_path"`
	Count           int       `json:"request_count"`
	AvgResponseTime float64   `json:"avg_response_time"`
	LastAccessed    time.Time `json:"last_accessed"`
	RequestTime     time.Time `json:"date"`
}
type TrafficCountStats struct {
	Count       int       `json:"request_count"`
	RequestTime time.Time `json:"date"`
}
