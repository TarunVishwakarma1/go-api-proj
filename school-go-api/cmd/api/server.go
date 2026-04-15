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
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST method on teachers route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on teachers route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on teachers route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on teachers route"))
		return
	default:
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Students route"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST method on Students route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Students route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Students route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Students route"))
		return
	default:
	}

}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Execs route"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST method on Execs route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Execs route"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Execs route"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Execs route"))
		return
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

	server := &http.Server{
		TLSConfig: tlsConfig,
		Addr:      port,
		Handler:   mw.Compression(mux),
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
