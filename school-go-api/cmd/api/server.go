package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Root Route")
	w.Write([]byte("Hello root route"))
	fmt.Println("Hello root route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// teachers/{id}
		//teachers/?key=value&query=value&sortby=email&sortorder=ASC
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userId := strings.TrimSuffix(path, "/")

		fmt.Println("The id is:", userId)

		fmt.Println("Query Params:", r.URL.Query())
		queryParams := r.URL.Query()
		sortBy := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortOrder := queryParams.Get("sortorder")

		if sortOrder == "" {
			sortOrder = "DESC"
		}

		fmt.Printf("SortBy: %v, Key: %v, SortOrder: %v\n", sortBy, key, sortOrder)

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
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
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

}

func execsHandler(w http.ResponseWriter, r *http.Request) {
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

}
func main() {
	port := ":3000"

	// Root Route
	http.HandleFunc("/", rootHandler)

	// Teachers Route
	http.HandleFunc("/teachers/", teachersHandler)

	// Students Route
	http.HandleFunc("/students/", studentsHandler)

	// EXECS Route
	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}

}
