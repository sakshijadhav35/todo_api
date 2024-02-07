package main

import (
    "log"
    "net/http"
)

func main() {
    // Initialize router
    router := NewRouter()

    // Set up routes
    router.HandleFunc("/todos", ListTodos).Methods("GET")
    router.HandleFunc("/todos", CreateTodo).Methods("POST")
    router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
    router.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
    router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

    // Start server
    log.Fatal(http.ListenAndServe(":8080", router))
}
