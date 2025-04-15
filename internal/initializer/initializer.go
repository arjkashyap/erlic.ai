package initializer

import (
	"database/sql"

	"github.com/arjkashyap/erlic.ai/internal/api/handlers"
	"github.com/arjkashyap/erlic.ai/internal/config"
	"github.com/arjkashyap/erlic.ai/internal/db/repositories"
)

func InitializeHandlers(r *repositories.Repositories, cfg *config.Config) *handlers.Handlers {
	return handlers.NewHandlers(r.UserRepository, cfg)
}

func InitializeRepositories(db *sql.DB) *repositories.Repositories {
	return repositories.NewRepositories(db)
}
