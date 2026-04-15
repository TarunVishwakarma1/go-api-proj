package router

import (
	"net/http"
	"school-go-api/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	// Root Route
	mux.HandleFunc("/", handlers.RootHandler)

	// Teachers Route
	mux.HandleFunc("/teachers/", handlers.TeachersHandler)

	// Students Route
	mux.HandleFunc("/students/", handlers.StudentsHandler)

	// EXECS Route
	mux.HandleFunc("/execs/", handlers.ExecsHandler)

	return mux
}
