package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func getAllTodos(ctx context.Context, userId int64) ([]Todo, error) {
	rows, err := db.Query(ctx, "SELECT id, user_id, title, content, done FROM todos WHERE user_id=$1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todoList []Todo

	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Content, &t.Done); err != nil {
			return nil, err
		}
		todoList = append(todoList, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todoList, nil
}

func createTodo(ctx context.Context, t Todo) error {
	_, err := db.Exec(
		ctx,
		"INSERT INTO todos (user_id, title, content) VALUES ($1, $2, $3)",
		t.UserID,
		t.Title,
		t.Content,
	)
	return err
}

func updateTodo(ctx context.Context, id string, input UpdateTodoInput) error {
	existing, err := getTodo(ctx, id)

	if err != nil {
		return err
	}

	if input.Title != nil {
		existing.Title = *input.Title
	}
	if input.Content != nil {
		existing.Content = *input.Content
	}
	if input.Done != nil {
		existing.Done = *input.Done
	}

	_, err = db.Exec(
		ctx,
		"UPDATE todos SET title=$1, content=$2, done=$3 WHERE id=$4",
		existing.Title,
		existing.Content,
		existing.Done,
		id,
	)
	return err
}

func deleteTodo(ctx context.Context, id string) error {
	_, err := db.Exec(
		ctx,
		"DELETE FROM todos WHERE id=$1",
		id,
	)

	return err
}

func getTodo(ctx context.Context, id string) (Todo, error) {
	row := db.QueryRow(
		ctx,
		"SELECT id, title, content, done FROM todos WHERE id=$1",
		id,
	)
	var t Todo
	if err := row.Scan(&t.ID, &t.Title, &t.Content, &t.Done); err != nil {
		if err == pgx.ErrNoRows {
			return Todo{}, ErrNotFound
		}
		return Todo{}, err
	}

	return t, nil
}
