package middleware

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the client accept gun zip encoding header
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		// Wrap the response writer

		w = &gzipResponseWriter{
			ResponseWriter: w, writer: gz,
		}

		next.ServeHTTP(w, r)
		fmt.Println("Sent response from Compression Middleware")
	})
}

// gzip response writer wraps http.ResponseWriter to write gzipped responses
type gzipResponseWriter struct {
	http.ResponseWriter
	writer *gzip.Writer
}

func (gw *gzipResponseWriter) Write(b []byte) (int, error) {
	return gw.writer.Write(b)
}
