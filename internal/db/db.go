package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/arjkashyap/erlic.ai/internal/config"
	"github.com/arjkashyap/erlic.ai/internal/logger"
	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	logger.Logger.Info("Connecting to DB")
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.DB.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	logger.Logger.Info("Connection to DB - Success")
	return db, nil
}
