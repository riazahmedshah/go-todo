package middlewares

import (
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/riazahmedshah/todo/models"
	"github.com/riazahmedshah/todo/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			utils.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "no cookie found"})
			return
		}

		tokenStr := cookie.Value

		token, err := jwt.ParseWithClaims(
			tokenStr,
			&models.Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			utils.WriteJson(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
			return
		}

		claims := token.Claims.(*models.Claims)
		ctx := context.WithValue(r.Context(), models.UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
