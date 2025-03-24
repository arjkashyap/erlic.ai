package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/arjkashyap/erlic.ai/internal/config"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
<<<<<<< HEAD
	// establish connection with db
=======

>>>>>>> 71dd5c4 (Added/modified files to new branch: develop)
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

	return db, nil
}
