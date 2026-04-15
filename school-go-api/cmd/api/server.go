package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "school-go-api/internal/api/middleware"
	"school-go-api/internal/api/router"
)

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQery:                   true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy, sortOrder", "name", "age", "city"},
	// }

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	router := router.Router()
	secureMux := mw.SecurityHeaders(router)

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
