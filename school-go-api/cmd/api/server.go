package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"school-go-api/internal/api/handlers"
	mw "school-go-api/internal/api/middleware"
)

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	// Root Route
	mux.HandleFunc("/", handlers.RootHandler)

	// Teachers Route
	mux.HandleFunc("/teachers/", handlers.TeachersHandler)

	// Students Route
	mux.HandleFunc("/students/", handlers.StudentsHandler)

	// EXECS Route
	mux.HandleFunc("/execs/", handlers.ExecsHandler)

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

	// secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTime, rl.RateLimiter, mw.CORS)

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
