package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/faysal0x1/go-bookstore/pkg/responses"
	"github.com/faysal0x1/go-bookstore/pkg/services"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	authService services.AuthService
}

func NewAuthMiddleware(authService services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}
}

func (m *AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			responses.Error(w, http.StatusUnauthorized, "Authorization header is required")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			responses.Error(w, http.StatusUnauthorized, "Invalid authorization header format")
			return
		}

		token, err := m.authService.ValidateToken(parts[1])
		if err != nil || !token.Valid {
			responses.Error(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			responses.Error(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		// Inject claims into context
		ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
		ctx = context.WithValue(ctx, "email", claims["email"])
		ctx = context.WithValue(ctx, "role", claims["role"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
