package middleware

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func RateLimiter(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Every(time.Second*1), 10)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
