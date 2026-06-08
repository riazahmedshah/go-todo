package main

import (
	"encoding/json"
	"net/http"
)

func getAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := getAllTodos(r.Context())
	if err != nil {
		writeJson(w, http.StatusInternalServerError, err)
		return
	}
	writeJson(w, http.StatusOK, todos)
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var t Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	if err := createTodo(r.Context(), t); err != nil {
		writeJson(w, http.StatusInternalServerError, err)
		return
	}
	writeJson(w, http.StatusOK, t)
}

func updateTodohandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var body UpdateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	if err := updateTodo(r.Context(), id, body); err != nil {
		writeJson(w, http.StatusInternalServerError, err)
		return
	}

	writeJson(w, http.StatusOK, map[string]string{"message": "Todo updated"})
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := deleteTodo(r.Context(), id); err != nil {
		writeJson(w, http.StatusInternalServerError, err)
		return
	}

	writeJson(w, http.StatusOK, map[string]string{"message": "Todo deleted"})
}

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	todo, err := getTodo(r.Context(), id)

	if err != nil {
		writeJson(w, http.StatusInternalServerError, err)
		return
	}

	writeJson(w, http.StatusOK, todo)
}
