package routes

import (
		"github.com/go-chi/chi/v5"
		"github.com/zzibo/cvwo-web-forum/internal/handlers"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/api/hello", handlers.HelloHandler)
	}
}
