package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "school-go-api/internal/api/middleware"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello root route"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on teachers route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on teachers route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on teachers route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on teachers route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on teachers route"))
	default:
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Students route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on Students route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Students route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Students route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Students route"))
	default:
	}

}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Execs route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on Execs route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Execs route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Execs route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Execs route"))
	default:
	}

}

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	// Root Route
	mux.HandleFunc("/", rootHandler)

	// Teachers Route
	mux.HandleFunc("/teachers/", teachersHandler)

	// Students Route
	mux.HandleFunc("/students/", studentsHandler)

	// EXECS Route
	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQery:                   true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy, sortOrder", "name", "age", "city"},
	// }

	secureMux := mw.SecurityHeaders(mux)

	// secureMux := applyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTime, rl.RateLimiter, mw.CORS)

	server := &http.Server{
		TLSConfig: tlsConfig,
		Addr:      port,
		Handler:   secureMux, //secureMux,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}

type Middleware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
