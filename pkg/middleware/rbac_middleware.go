package middleware

import (
	"net/http"

	"github.com/faysal0x1/go-bookstore/pkg/responses"
)

func RoleMiddleware(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := r.Context().Value("role")
			if userRole == nil {
				responses.Error(w, http.StatusUnauthorized, "User role not found in context")
				return
			}

			roleStr, ok := userRole.(string)
			if !ok {
				responses.Error(w, http.StatusUnauthorized, "Invalid role format")
				return
			}

			authorized := false
			for _, role := range roles {
				if role == roleStr {
					authorized = true
					break
				}
			}

			if !authorized {
				responses.Error(w, http.StatusForbidden, "You do not have permission to access this resource")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
