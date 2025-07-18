package middleware

import "net/http"

func WithAuthMiddleware() func (h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return h
	}
}
