package routers

import (
	"net/http"

	"github.com/riazahmedshah/todo/handlers"
	"github.com/riazahmedshah/todo/middlewares"
)

func SetupRoutes(mux *http.ServeMux) {
	// TODO Routes
	mux.HandleFunc("GET /todos", middlewares.AuthMiddleware(handlers.GetAllTodosHandler))
	mux.HandleFunc("GET /todos/{id}", middlewares.AuthMiddleware(handlers.GetTodoHandler))
	mux.HandleFunc("POST /todos", middlewares.AuthMiddleware(handlers.CreateTodoHandler))
	mux.HandleFunc("PATCH /todos/{id}", middlewares.AuthMiddleware(handlers.UpdateTodoHandler))
	mux.HandleFunc("DELETE /todos/{id}", middlewares.AuthMiddleware(handlers.DeleteTodoHandler))

	// USER Routes
	mux.HandleFunc("POST /register", handlers.RegisterHandler)
	mux.HandleFunc("POST /login", handlers.LoginHandler)
}
