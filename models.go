package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
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

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

var ErrNotFound = errors.New("not found")

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
