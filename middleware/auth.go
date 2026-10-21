
package middleware

import (
	"context"
	"fmt"
	"goMod/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const (
	UserIDKey ContextKey = "userId"
	EmailKey  ContextKey = "email"
)

func JWTMiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization Header not present", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorization Token not present", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			http.Error(w, "Token is Required", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}

				return []byte(config.GetString("JWT_STRING", "token")), nil
			},
		)

		if err != nil || token == nil || !token.Valid {
			http.Error(w, "Token Not Correct", http.StatusUnauthorized)
			return
		}

		userId, okId := claims["id"].(string)
		email, okEmail := claims["email"].(string)

		if !okId || !okEmail || userId == "" || email == "" {
			http.Error(w, "Invalid Token Error", http.StatusUnauthorized)
			return
		}

		fmt.Println("User id is:", userId, "Email is:", email)

		ctx := context.WithValue(r.Context(), UserIDKey, userId)
		ctx = context.WithValue(ctx, EmailKey, email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
