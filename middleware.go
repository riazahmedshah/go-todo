package main

import (
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			writeJson(w, http.StatusBadRequest, map[string]string{"error": "no cookie found"})
			return
		}

		tokenStr := cookie.Value

		token, err := jwt.ParseWithClaims(
			tokenStr,
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			writeJson(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
			return
		}

		claims := token.Claims.(*Claims)
		ctx := context.WithValue(r.Context(), "userId", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
