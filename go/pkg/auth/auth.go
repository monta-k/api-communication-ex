package auth

import (
	"context"
	"net/http"
	"strings"
)

type userContextKey struct{}

var userCtxKey = &userContextKey{}

type User struct {
	Name string
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Authorization token is required", http.StatusUnauthorized)
			return
		}

		if token != "token" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		user := &User{
			Name: "user",
		}

		r = r.WithContext(context.WithValue(r.Context(), userCtxKey, user))

		next.ServeHTTP(w, r)
	})
}

func UserFromContext(ctx context.Context) *User {
	user, ok := ctx.Value(userCtxKey).(*User)
	if !ok {
		return nil
	}

	return user
}
