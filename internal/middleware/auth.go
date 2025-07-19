package middleware

import (
	"fmt"
	"marketplace/internal/api/http/response"
	"marketplace/pkg/utils"
	"net/http"
)

func WithAuthMiddleware(publicKey []byte, debugMode bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			login, err := utils.ProcessAuthHeader(r, publicKey)

			if err != nil {
				response.ProcessError(w, fmt.Errorf("Auth middleware: %w", err), debugMode)
				return
			}

			r.Header.Set("X-User-Login", login)
			next.ServeHTTP(w, r)
		})
	}
}
