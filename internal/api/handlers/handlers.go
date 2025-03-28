package handlers

import (
	"github.com/arjkashyap/erlic.ai/internal/db/repositories"
)

type Handlers struct {
	HealthCheck *HealthCheckHandler
	UserHandler *UserHandler
	AuthHandler *AuthHandler
}

func NewHandlers(ur *repositories.UserRepository) *Handlers {
	return &Handlers{
		HealthCheck: NewHealthCheckHandler(),
		UserHandler: NewUserHandler(ur),
		AuthHandler: NewAuthHandler(),
	}
}
