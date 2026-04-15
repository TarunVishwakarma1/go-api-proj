package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	mw "school-go-api/internal/api/middleware"
	"strconv"
	"strings"
	"sync"
)

type Teacher struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Class     string `json:"class,omitempty"`
	Subject   string `json:"subject,omitempty"`
}

var teachers = make(map[int]Teacher)
var mutex = &sync.Mutex{}
var nextId = 1

func init() {
	teachers[nextId] = Teacher{
		ID:        nextId,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Maths",
	}
	nextId++
	teachers[nextId] = Teacher{
		ID:        nextId,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextId++
	teachers[nextId] = Teacher{
		ID:        nextId,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "11A",
		Subject:   "Biology",
	}
	nextId++
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	if idStr == "" {
		firstName := r.URL.Query().Get("firstName")
		lastName := r.URL.Query().Get("lastName")

		teacherList := make([]Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || lastName == teacher.LastName) {
				teacherList = append(teacherList, teacher)
			}
		}

		response := struct {
			Status string    `json:"status"`
			Count  int       `json:"count"`
			Data   []Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teachers),
			Data:   teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// Handler Path parameter
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
	defer mutex.Unlock()
	mutex.Lock()

	var newTeachers []Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	addedTeacher := make([]Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextId
		teachers[nextId] = newTeacher
		addedTeacher[i] = newTeacher
		nextId++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string    `json:"status"`
		Count  int       `json:"count"`
		Data   []Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeacher),
		Data:   addedTeacher,
	}

	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello root route"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// call get method handler func
		getTeachersHandler(w, r)
	case http.MethodPost:
		// Post request handler
		addTeacherHandler(w, r)
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
