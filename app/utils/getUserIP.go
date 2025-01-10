package utils

import (
	"net/http"
	"strings"
)

func GetUserIP(r *http.Request) string {
	realIP := r.Header.Get("X-Forwarded-For")
	if realIP != "" {
		ips := strings.Split(realIP, ",")
		return strings.TrimSpace(ips[0])
	}
	return r.RemoteAddr
}
