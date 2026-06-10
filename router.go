package main

import "net/http"

func setupRoutes(mux *http.ServeMux) {
	// TODO Routes
	mux.HandleFunc("GET /todos", authMiddleware(getAllTodosHandler))
	mux.HandleFunc("GET /todos/{id}", authMiddleware(getTodoHandler))
	mux.HandleFunc("POST /todos", authMiddleware(createTodoHandler))
	mux.HandleFunc("PATCH /todos/{id}", authMiddleware(updateTodoHandler))
	mux.HandleFunc("DELETE /todos/{id}", authMiddleware(deleteTodoHandler))

	// USER Routes
	mux.HandleFunc("POST /register", registerHandler)
	mux.HandleFunc("POST /login", loginHandler)
}
