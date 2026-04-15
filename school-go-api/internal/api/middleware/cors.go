package middleware

import "net/http"

// Allowed Origins
var allowedOrigins = []string{
	"https://mu-origin.com",
	"http://localhost:3000",
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Accept-Encoding")
		if !isOriginAllowedOrigin(origin) {
			http.Error(w, "Not Allowed", http.StatusForbidden)
			return
		} else {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "Content-Type, true")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		next.ServeHTTP(w, r)
	})
}

func isOriginAllowedOrigin(origin string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin != allowedOrigin {
			return false
		}
	}
	return true
}
