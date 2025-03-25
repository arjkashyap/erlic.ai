package config

import (
	"database/sql"

	"github.com/arjkashyap/erlic.ai/internal/api/handlers"
	"github.com/arjkashyap/erlic.ai/internal/db/repositories"
	"github.com/arjkashyap/erlic.ai/internal/env"
)

type DBConfig struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type Config struct {
	Port string
	Env  string
	DB   DBConfig
}

func NewConfig() *Config {

	db_conf := DBConfig{
		DSN:          env.GetString("DSN", "postgres://postgres:password@localhost:5432/erlic?sslmode=disable"),
		MaxOpenConns: 25,
		MaxIdleConns: 25,
		MaxIdleTime:  "15m",
	}

	return &Config{
		Port: env.GetString("PORT", ":8080"),
		Env:  env.GetString("ENV", "DEV"),
		DB:   db_conf,
	}
}

func (c *Config) InitializeHandlers(r *repositories.Repositories) *handlers.Handlers {
	return handlers.NewHandlers(r.UserRepository)
}

func (c *Config) InitializeRepositories(db *sql.DB) *repositories.Repositories {
	return repositories.NewRepositories(db)
}
