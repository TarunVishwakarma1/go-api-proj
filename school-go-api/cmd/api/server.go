package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "school-go-api/internal/api/middleware"
	"school-go-api/internal/api/router"
	"school-go-api/internal/repositories/sqlconnect"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error", err)
	}

	_, err := sqlconnect.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
	}
	port := os.Getenv("API_PORT")

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
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
