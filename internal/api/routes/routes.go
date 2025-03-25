package routes

import (
	"github.com/arjkashyap/erlic.ai/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, h *handlers.Handlers) {
	r.Get("/health-check", h.HealthCheck.HealthCheckHandler)

	r.Route("/users", func(r chi.Router) {
		r.Post("/create-user", h.UserHandler.CreateUser)
	})

}
