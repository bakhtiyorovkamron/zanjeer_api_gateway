package main

import (
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/api"
	"github.com/Projects/zanjeer_api_gateway/config"
	"github.com/Projects/zanjeer_api_gateway/pkg/db"
	"github.com/Projects/zanjeer_api_gateway/pkg/logger"
	"github.com/Projects/zanjeer_api_gateway/storage"
)

func main() {
	cfg := config.Load()

	logger := logger.New(cfg.LogLevel)

	db, err := db.New(cfg)
	if err != nil {
		logger.Error("Error while connecting to database", err)
	} else {
		logger.Info("Successfully connected to database")
	}
	fmt.Println("Database :", db)

	r := api.New(cfg, storage.New(db, logger, cfg), logger)
	r.Run(":7777")
}
