package main

import "net/http"

func setupTodoRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /todos", getAllTodosHandler)
	mux.HandleFunc("GET /todos/{id}", getTodoHandler)
	mux.HandleFunc("POST /todos", createTodoHandler)
	mux.HandleFunc("PATCH /todos/{id}", updateTodohandler)
	mux.HandleFunc("DELETE /todos/{id}", deleteTodoHandler)
}
