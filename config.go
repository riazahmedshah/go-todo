package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func initDB() error {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

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
