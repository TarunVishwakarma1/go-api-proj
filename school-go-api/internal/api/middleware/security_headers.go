package middleware

import (
	"fmt"
	"net/http"
)

func SecurityHeaders(next http.Handler) http.Handler {
	fmt.Println("Security Headers Middlewares...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Security Headers Middlewares being returned...")

		w.Header().Set("X-DNS-Prefetch-Control", "off")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1;mode=block")
		w.Header().Set("X-Content-Type-Options", "nonsniff")
		w.Header().Set("Strict Transport Security", "max-age=63072000;includeSubDomains;preload")
		w.Header().Set("Referred-Policy", "no-referrer")

		next.ServeHTTP(w, r)
		fmt.Println("Security Headers Middlewares ends...")

	})
}
