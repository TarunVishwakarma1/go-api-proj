package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]int
	limit     int
	resetTime time.Duration
}

func NewRateLimiter(limit int, resetTime time.Duration) *rateLimiter {
	r1 := &rateLimiter{
		limit:     limit,
		visitors:  make(map[string]int),
		resetTime: resetTime,
	}

	go r1.resetVisitorCount()
	return r1
}

func (r *rateLimiter) resetVisitorCount() {
	for {
		time.Sleep(r.resetTime)
		r.mu.Lock()
		r.visitors = make(map[string]int)
		r.mu.Unlock()
	}
}

func (rl *rateLimiter) RateLimiter(next http.Handler) http.Handler {
	fmt.Println("Rate Limiter Middleware ....")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Rate Limiter middleware being returned")
		
		visitorIP := r.RemoteAddr //Might want to extract the IP in a more sophisticate way
		
		rl.mu.Lock()
		rl.visitors[visitorIP]++
		count := rl.visitors[visitorIP]
		rl.mu.Unlock()
		
		fmt.Printf("Visitor Count from %v is %v\n", visitorIP, count)

		if count > rl.limit {
			http.Error(w, "Too many request", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
		fmt.Println("Rate Limiter Middleware ends....")

	})
}
