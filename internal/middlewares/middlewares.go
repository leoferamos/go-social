package middlewares

import (
	"go_social/internal/auth"
	"go_social/internal/responses"
	"net/http"
)

// AuthMiddleware is a middleware that checks if the request has a valid authentication token.
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.JSONError(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
