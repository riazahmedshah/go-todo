package main

import "errors"

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
