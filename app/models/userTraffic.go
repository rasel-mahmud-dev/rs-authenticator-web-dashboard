package models

import "time"

type UserTraffic struct {
	ID           int           `json:"id"`
	RoutePath    string        `json:"route_path"`
	HTTPMethod   string        `json:"http_method"`
	UserAgent    string        `json:"user_agent"`
	IPAddress    string        `json:"ip_address"`
	RequestTime  time.Time     `json:"request_time"`
	ResponseTime time.Duration `json:"response_time"`
}
