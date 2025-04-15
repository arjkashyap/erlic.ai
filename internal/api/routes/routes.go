package routes

import (
	"github.com/arjkashyap/erlic.ai/internal/api/handlers"
	"github.com/arjkashyap/erlic.ai/internal/api/middleware"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, h *handlers.Handlers) {
	r.Get("/health-check", h.HealthCheck.HealthCheckHandler)

	r.Route("/api", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Get("/{provider}", h.AuthHandler.AuthInitiate)
			r.Get("/{provider}/callback", h.AuthHandler.AuthCallback)
			r.Get("/logout/{provider}", h.AuthHandler.AuthLogout)
			r.Get("/me", h.AuthHandler.GetCurrentUser)
		})

		r.Route("/user", func(r chi.Router) {
			r.Post("/register", h.UserHandler.CreateUser)

		})
		// some
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)

			r.Route("/dashboard", func(r chi.Router) {
				r.Get("/", h.DashboardHandler.GetDashboard)
			})

			r.Route("/profile", func(r chi.Router) {
				r.Get("/", h.ProfileHandler.GetProfile)
			})

			r.Route("/chat", func(r chi.Router) {
				r.Post("/prompt", h.ChatHandler.HandlePrompt)
			})
		})
	})
}
