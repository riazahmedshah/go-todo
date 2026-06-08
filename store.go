package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func initDB() error {
	dbURL := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), dbURL)

	if err != nil {
		return err
	}

	err = pool.Ping(context.Background())

	if err != nil {
		return err
	}

	db = pool
	return nil
}

func getALlTodos(ctx context.Context) ([]Todo, error) {
	rows, err := db.Query(ctx, "SELECT id, title, conten, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todoList []Todo

	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.Done); err != nil {
			return nil, err
		}
		todoList = append(todoList, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todoList, nil
}
