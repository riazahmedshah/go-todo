package main

import (
	"encoding/json"
	"net/http"
)

func getAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(UserIDKey).(int64)
	todos, err := getAllTodos(r.Context(), userId)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]any{"error": err})
		return
	}
	writeJson(w, http.StatusOK, todos)
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(UserIDKey).(int64)
	var t Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}
	t.UserID = userId

	if err := createTodo(r.Context(), t); err != nil {
		writeJson(w, http.StatusInternalServerError, err)
		return
	}
	writeJson(w, http.StatusOK, t)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
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
