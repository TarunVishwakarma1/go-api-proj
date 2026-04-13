package main

import (
	"encoding/json"
	"fmt"
	"io"
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

			// Parse the form data (necessary for x-www-form-urlencoded)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			fmt.Println("Form:", r.Form)

			// Prepare response data
			response := make(map[string]any)
			for k, v := range r.Form {
				response[k] = v[0]
			}
			fmt.Println("Processed Response Map:", response)

			// RAW BODY

			body, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()

			fmt.Println("Raw body:", string(body))

			var userInstance user

			err = json.Unmarshal(body, &userInstance)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error while marshlling %s", err), http.StatusBadRequest)
				return
			}

			fmt.Println("Unmarshalled json into user instance", userInstance)
			fmt.Println("Recieved User name as:", userInstance.Name)

			// Prepare response data
			response1 := make(map[string]any)
			for k, v := range r.Form {
				response[k] = v[0]
			}

			err = json.Unmarshal(body, &response1)
			if err != nil {
				http.Error(w, "Error while marshlling", http.StatusBadRequest)
				return
			}
			fmt.Println("Unmarshalled json into a map", response1)

			fmt.Println("Body:", r.Body)
			fmt.Println("Form:", r.Form)
			fmt.Println("Header:", r.Header)
			fmt.Println("Context:", r.Context())
			fmt.Println("Context Length:", r.ContentLength)
			fmt.Println("Host:", r.Host)
			fmt.Println("Method:", r.Method)
			fmt.Println("Protocol:", r.Proto)
			fmt.Println("Remote Addr:", r.RemoteAddr)
			fmt.Println("Request URI:", r.RequestURI)
			fmt.Println("TLS:", r.TLS)
			fmt.Println("Trailer:", r.Trailer)
			fmt.Println("Transfer Encoding:", r.TransferEncoding)
			fmt.Println("URL:", r.URL)
			fmt.Println("UserAgent:", r.UserAgent())
			fmt.Println("Port:", r.URL.Port())
			fmt.Println("Scheme:", r.URL.Scheme)

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
		w.Write([]byte("Hello students route"))
		fmt.Println("Hello students route")

	})

	// EXECS Route
	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello execs route"))
		fmt.Println("Hello execs route")

	})

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
