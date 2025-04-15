package handlers

import (
	"github.com/arjkashyap/erlic.ai/internal/config"
	"github.com/arjkashyap/erlic.ai/internal/db/repositories"
)

type Handlers struct {
	HealthCheck      *HealthCheckHandler
	UserHandler      *UserHandler
	AuthHandler      *AuthHandler
	DashboardHandler *DashboardHandler
	ProfileHandler   *ProfileHandler
	ChatHandler      *ChatHandler
}

func NewHandlers(ur *repositories.UserRepository, cfg *config.Config) *Handlers {
	return &Handlers{
		HealthCheck:      NewHealthCheckHandler(),
		UserHandler:      NewUserHandler(ur),
		AuthHandler:      NewAuthHandler(ur),
		DashboardHandler: NewDashboardHandler(),
		ProfileHandler:   NewProfileHandler(ur),
		ChatHandler:      NewChatHandler(cfg),
	}
}
