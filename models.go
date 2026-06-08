package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Todo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type UpdateTodoInput struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Done    *bool   `json:"done"`
}

var ErrNotFound = errors.New("not found")

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
