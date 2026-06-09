package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func createUser(ctx context.Context, u User) error {

	_, err := db.Exec(
		ctx,
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		u.Username,
		u.Email,
		u.Password,
	)
	return err
}

func findUserByEmail(ctx context.Context, email string) (User, error) {
	var u User
	row := db.QueryRow(ctx, "SELECT id, username, email FROM users WHERE email=$1", email)

	if err := row.Scan(&u.ID, &u.Username, &u.Email); err != nil {
		if err == pgx.ErrNoRows {
			return User{}, ErrNotFound
		}
		return User{}, err
	}

	return u, nil
}

func getUserById(ctx context.Context, id int) (User, error) {
	var u User
	row := db.QueryRow(ctx, "SELECT id, username, email FROM users WHERE id=$1", id)

	if err := row.Scan(&u.ID, &u.Username, &u.Email); err != nil {
		return User{}, err
	}

	return u, nil
}
