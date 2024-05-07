package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// AuthMiddleware is a middleware for authentication
func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/auth/login" || r.URL.Path == "/auth/register" {
				next.ServeHTTP(w, r)
				return
			}
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader == "" {
				http.Error(w, "No Authorization header provided", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Validate signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Extract user info from token claims
			claim, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user", claim["user"])
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AdminOnlyMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract the token from the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			tokenStr := authHeader[len("Bearer "):]
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				// Ensure the token uses the correct signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			// Extract role information from the token claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Unable to extract claims from token", http.StatusUnauthorized)
				return
			}

			role, ok := claims["role"].(string)
			if !ok || role != "admin" {
				http.Error(w, "Access forbidden: Admins only", http.StatusForbidden)
				return
			}

			// Pass the request to the next handler if authorization is valid
			next.ServeHTTP(w, r)
		})
	}
}
