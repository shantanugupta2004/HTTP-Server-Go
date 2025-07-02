package middlewares

import (
	"context"
	"net/http"
	"strings"
	"http-server-go/authutils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userEmailKey contextKey = "userEmail"

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &authutils.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("shantanu@2004"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Store user info in context
		ctx := context.WithValue(r.Context(), userEmailKey, claims.Email)

		// Pass new context with user email to handler
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
