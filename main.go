package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/google/uuid"
)

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

var (
	todos = map[string]Todo{}
	mu    sync.Mutex
)

func saveToFile() error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("todos.json", data, 0644)
}

func loadFromFile() error {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return json.Unmarshal(data, &todos)

}

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
		saveToFile()
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
		saveToFile()
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
		saveToFile()
		mu.Unlock()

		writeJson(w, http.StatusOK, map[string]string{"message": "deleted"})
	default:
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func main() {
	if err := initDB(); err != nil {
		log.Fatalf("Failed to Connect DataBase: %v", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", todosHandler)
	mux.HandleFunc("/todos/", todoHandler)

	const port = ":3000"
	addr := "http://localhost" + port

	log.Printf("Server started on %s", addr)
	log.Println("Press Ctrl+C to stop the server...")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
