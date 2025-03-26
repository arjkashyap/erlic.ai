package routes

import (
	"github.com/arjkashyap/erlic.ai/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, h *handlers.Handlers) {
	r.Get("/health-check", h.HealthCheck.HealthCheckHandler)

	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}", h.AuthHandler.AuthInitiate)
		r.Get("/{provider}/callback", h.AuthHandler.AuthCallback)
		r.Get("/logout/{provider}", h.AuthHandler.AuthLogout)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", h.UserHandler.CreateUser)
	})

}
