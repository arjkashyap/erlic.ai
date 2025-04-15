package config

import (
	"github.com/arjkashyap/erlic.ai/internal/env"
)

type DBConfig struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type Config struct {
	Port       string
	Env        string
	MLProvider string
	DB         DBConfig
}

func NewConfig() *Config {
	db_conf := DBConfig{
		DSN:          env.GetString("DSN", "postgres://postgres:password@localhost:5432/erlic?sslmode=disable"),
		MaxOpenConns: 25,
		MaxIdleConns: 25,
		MaxIdleTime:  "15m",
	}

	return &Config{
		Port:       env.GetString("PORT", ":8080"),
		Env:        env.GetString("ENV", "DEV"),
		MLProvider: env.GetString("ML_PROVIDER", "vertex"),
		DB:         db_conf,
	}
}
