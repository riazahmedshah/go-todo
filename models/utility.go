package models

import "errors"

type contextKey string

const UserIDKey contextKey = "userID"

var ErrNotFound = errors.New("not found")
