package main

import (
	"github.com/arjkashyap/erlic.ai/internal/api"
	"github.com/arjkashyap/erlic.ai/internal/config"
	"github.com/arjkashyap/erlic.ai/internal/db"
	"github.com/arjkashyap/erlic.ai/internal/logger"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	defer logger.CloseLogger()
	config := config.NewConfig()

	db, err := db.Connect(config)
	if err != nil {
		logger.Logger.Fatal("Failed to connect to the database", zap.Error(err))
		panic(err)
	}

	defer db.Close()

	hr := config.InitializeHandlers(config.InitializeRepositories(db))
	srv := api.NewAPI(logger.Logger, config, hr)

	err = srv.Run()
	if err != nil {
		logger.Logger.Fatal("Failed to start the server", zap.Error(err))
	}
}
