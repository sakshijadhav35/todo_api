package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type Todo struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

func ListTodos(w http.ResponseWriter, r *http.Request) {
	page, pageSize := getPaginationParams(r)

	// For simplicity, returning mock data
	todos := []Todo{
		{ID: "1", UserID: "user1", Title: "Task 1", Description: "Description 1", Status: "Pending", Created: time.Now(), Updated: time.Now()},
		{ID: "2", UserID: "user1", Title: "Task 2", Description: "Description 2", Status: "Completed", Created: time.Now(), Updated: time.Now()},
	}

	// Paginate todos
	start := (page - 1) * pageSize
	end := page * pageSize
	if start >= len(todos) {
		todos = []Todo{}
	} else if end > len(todos) {
		todos = todos[start:]
	} else {
		todos = todos[start:end]
	}

	// Convert todos to JSON and send response
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate unique ID
	todo.ID = gocql.TimeUUID().String()

	// Set created and updated timestamps
	now := time.Now()
	todo.Created = now
	todo.Updated = now

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	// Extract todo ID from request
	vars := mux.Vars(r)
	todoID := vars["id"]

	// For simplicity, returning mock data
	todo := Todo{ID: todoID, UserID: "user1", Title: "Task 1", Description: "Description 1", Status: "Pending", Created: time.Now(), Updated: time.Now()}

	// Convert todo to JSON and send response
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Extract todo ID from request
	vars := mux.Vars(r)
	todoID := vars["id"]

	var updatedTodo Todo
	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Extract todo ID from request
	vars := mux.Vars(r)
	todoID := vars["id"]

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted successfully"})
}

func getPaginationParams(r *http.Request) (int, int) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10 // Default page size
	}

	return page, pageSize
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", ListTodos).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
