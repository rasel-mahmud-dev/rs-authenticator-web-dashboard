package middlewares

import (
	"log"
	"net/http"
	"rs/auth/app/models"
	"rs/auth/app/repositories/trafficRepo"
	"time"
)

func Traffic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		trafficData := models.UserTraffic{
			RoutePath:    r.URL.Path,
			HTTPMethod:   r.Method,
			UserAgent:    r.Header.Get("User-Agent"),
			IPAddress:    r.RemoteAddr,
			RequestTime:  time.Now(),
			ResponseTime: time.Duration(time.Since(start).Milliseconds()),
		}

		go func(trafficData models.UserTraffic) {
			if err := trafficRepo.TrafficRepository.InsertApiTraffic(trafficData); err != nil {
				log.Printf("Failed to log API traffic: %v", err)
			}
		}(trafficData)
	})
}
