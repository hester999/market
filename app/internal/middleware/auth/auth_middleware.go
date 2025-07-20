package auth

import (
	"context"
	"errors"
	"market/app/internal/apperr"
	"net/http"
	"strings"
)

type AuthUsecaseMiddleware interface {
	ValidateSession(token string) (string, error)
}

func AuthMiddleware(auth AuthUsecaseMiddleware) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := r.Header.Get("Authorization")

			if token == "" {
				http.Error(w, "Token required", http.StatusUnauthorized)
				return
			}
			if !strings.HasPrefix(token, "Bearer ") {

				http.Error(w, "Invalid token format", http.StatusUnauthorized)
				return
			}
			userId, err := auth.ValidateSession(token)

			if err != nil {
				if errors.Is(err, apperr.ErrSessionExpired) {
					http.Error(w, "Session expired", http.StatusUnauthorized)
					return
				}
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			ctx := context.WithValue(r.Context(), "user_id", userId)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func OptionalAuth(auth AuthUsecaseMiddleware) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token := strings.TrimPrefix(authHeader, "Bearer ")
				userID, err := auth.ValidateSession(token)
				if err == nil {
					ctx := context.WithValue(r.Context(), "user_id", userID)
					r = r.WithContext(ctx)
				} else {
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
