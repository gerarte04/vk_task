package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf(
			"[HTTP/1.1] %s | %s | %s | %d | %s",
			r.Method, r.URL.Path, r.RemoteAddr, ww.Status(), time.Since(start).String(), 
		)
	})
}
