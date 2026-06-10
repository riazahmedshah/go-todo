package stores

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/riazahmedshah/todo/models"
)

func CreateUser(ctx context.Context, u models.User) error {

	_, err := db.Exec(
		ctx,
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		u.Username,
		u.Email,
		u.Password,
	)
	return err
}

func FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	var u models.User
	row := db.QueryRow(ctx, "SELECT id, password FROM users WHERE email=$1", email)

	if err := row.Scan(&u.ID, &u.Password); err != nil {
		if err == pgx.ErrNoRows {
			return models.User{}, models.ErrNotFound
		}
		return models.User{}, err
	}

	return u, nil
}

func GetUserById(ctx context.Context, id int64) (models.User, error) {
	var u models.User
	row := db.QueryRow(ctx, "SELECT id, username, email FROM users WHERE id=$1", id)

	if err := row.Scan(&u.ID, &u.Username, &u.Email); err != nil {
		return models.User{}, err
	}

	return u, nil
}
