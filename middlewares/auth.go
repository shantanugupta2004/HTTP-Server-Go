package middlewares

import (
    "net/http"
    "strings"
    "http-server-go/handlers"
    "github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        claims := &handlers.Claims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte("shantanu@2004"), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    }
}