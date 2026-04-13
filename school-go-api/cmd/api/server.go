package main

import (
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func main() {
	port := ":3000"

	// Root Route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello root route"))
		fmt.Println("Hello root route")
	})

	// Teachers Route
	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET method on teachers route"))
			fmt.Println("Hello GET method on teachers route")
			return
		case http.MethodPost:
			w.Write([]byte("Hello POST method on teachers route"))
			fmt.Println("Hello POST s method on teachers route")
			return
		case http.MethodPut:
			w.Write([]byte("Hello PUT method on teachers route"))
			fmt.Println("Hello PUT method on teachers route")
			return
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH method on teachers route"))
			fmt.Println("Hello PATCH method on teachers route")
			return
		case http.MethodDelete:
			w.Write([]byte("Hello DELETE method on teachers route"))
			fmt.Println("Hello DELETE method on teachers route")
			return
		default:
			w.Write([]byte("Hello teachers route"))
			fmt.Println("Hello teachers route")
		}
	})

	// Students Route
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Route")
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET method on Students route"))
			fmt.Println("Hello GET method on Students route")
			return
		case http.MethodPost:
			w.Write([]byte("Hello POST method on Students route"))
			fmt.Println("Hello POST s method on Students route")
			return
		case http.MethodPut:
			w.Write([]byte("Hello PUT method on Students route"))
			fmt.Println("Hello PUT method on Students route")
			return
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH method on Students route"))
			fmt.Println("Hello PATCH method on Students route")
			return
		case http.MethodDelete:
			w.Write([]byte("Hello DELETE method on Students route"))
			fmt.Println("Hello DELETE method on Students route")
			return
		default:
		}
		w.Write([]byte("Hello Students route"))
		fmt.Println("Hello Students route")

	})

	// EXECS Route
	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Route")
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET method on Execs route"))
			fmt.Println("Hello GET method on Execs route")
			return
		case http.MethodPost:
			w.Write([]byte("Hello POST method on Execs route"))
			fmt.Println("Hello POST s method on Execs route")
			return
		case http.MethodPut:
			w.Write([]byte("Hello PUT method on Execs route"))
			fmt.Println("Hello PUT method on Execs route")
			return
		case http.MethodPatch:
			w.Write([]byte("Hello PATCH method on Execs route"))
			fmt.Println("Hello PATCH method on Execs route")
			return
		case http.MethodDelete:
			w.Write([]byte("Hello DELETE method on Execs route"))
			fmt.Println("Hello DELETE method on Execs route")
			return
		default:
		}
		w.Write([]byte("Hello Execs route"))
		fmt.Println("Hello Execs route")

	})

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
