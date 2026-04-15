package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a custom response writer to capture the satus code
		wrappedWriter := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		duration := time.Since(start)
		w.Header().Set("X-Response-Time", duration.String())
		next.ServeHTTP(wrappedWriter, r)
		duration = time.Since(start)
		//. Log the request details
		fmt.Printf("Method %s, URL: %s, Status: %d, Duration: %v\n", r.Method, r.URL, wrappedWriter.status, duration.String())
		fmt.Println("Sent Response from Response time middlerware")
	})
}

// Response Writer
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriterHeader(code int) {
	rw.status = code
	rw.WriteHeader(code)
}
