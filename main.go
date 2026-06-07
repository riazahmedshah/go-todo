package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
)

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type Todo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

var (
	todos = map[string]Todo{}
	mu    sync.Mutex
)

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		list := make([]Todo, 0, len(todos))
		for _, todo := range todos {
			list = append(list, todo)
		}
		mu.Unlock()
		writeJson(w, http.StatusOK, list)

	case http.MethodPost:
		var t Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			writeJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid Body"})
			return
		}
		t.ID = uuid.NewString()
		mu.Lock()
		todos[t.ID] = t
		mu.Unlock()
		writeJson(w, http.StatusCreated, t)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/")
	if id == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "missing id"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		existing, ok := todos[id]
		if !ok {
			mu.Unlock()
			writeJson(w, http.StatusNotFound, map[string]string{"error": "Todo not found"})
			return
		}
		mu.Unlock()
		writeJson(w, http.StatusOK, existing)

	case http.MethodPut:
		var t Todo
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			writeJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid Body"})
			return
		}
		mu.Lock()
		existing, ok := todos[id]

		if !ok {
			mu.Unlock()
			writeJson(w, http.StatusNotFound, map[string]string{"error": "Todo not found"})
			return
		}

		existing.Title = t.Title
		existing.Content = t.Content
		existing.Done = t.Done
		todos[id] = existing
		mu.Unlock()

		writeJson(w, http.StatusOK, existing)

	case http.MethodDelete:
		mu.Lock()
		_, ok := todos[id]
		if !ok {
			mu.Unlock()
			writeJson(w, http.StatusNotFound, map[string]string{"error": "Todo not found"})
			return
		}

		delete(todos, id)
		mu.Unlock()

		writeJson(w, http.StatusOK, map[string]string{"message": "deleted"})
	default:
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", todosHandler)
	mux.HandleFunc("/todos/", todoHandler)
	http.ListenAndServe(":3000", mux)
}
