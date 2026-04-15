package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"school-go-api/internal/models"
	"strconv"
	"strings"
	"sync"
)

var teachers = make(map[int]models.Teacher)
var mutex = &sync.Mutex{}
var nextId = 1

func init() {
	teachers[nextId] = models.Teacher{
		ID:        nextId,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Maths",
	}
	nextId++
	teachers[nextId] = models.Teacher{
		ID:        nextId,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextId++
	teachers[nextId] = models.Teacher{
		ID:        nextId,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "11A",
		Subject:   "Biology",
	}
	nextId++
}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
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

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	if idStr == "" {
		firstName := r.URL.Query().Get("firstName")
		lastName := r.URL.Query().Get("lastName")

		teacherList := make([]models.Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || lastName == teacher.LastName) {
				teacherList = append(teacherList, teacher)
			}
		}

		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
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

	var newTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	addedTeacher := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextId
		teachers[nextId] = newTeacher
		addedTeacher[i] = newTeacher
		nextId++
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeacher),
		Data:   addedTeacher,
	}

	json.NewEncoder(w).Encode(response)
}
