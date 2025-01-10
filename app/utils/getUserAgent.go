package utils

import "net/http"

func GetUserAgent(r *http.Request) string {
	return r.UserAgent()
}
