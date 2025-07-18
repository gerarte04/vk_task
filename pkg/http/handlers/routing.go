package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RouterOption func(r chi.Router)

func WithRecovery() RouterOption {
	return func(r chi.Router) {
		r.Use(middleware.Recoverer)
	}
}
