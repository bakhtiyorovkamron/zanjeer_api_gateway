package handlers

import (
	"testing"

	"github.com/Projects/zanjeer_api_gateway/config"
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/Projects/zanjeer_api_gateway/pkg/db"
	"github.com/Projects/zanjeer_api_gateway/pkg/logger"
	"github.com/Projects/zanjeer_api_gateway/storage"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateAdmin(t *testing.T) {
	var resp models.Admin

	resp.Login = "shaxzod"
	resp.Password = "shaxzod123"

	cfg := config.Load()

	logger := logger.New(cfg.LogLevel)

	db, err := db.New(cfg)
	if err != nil {
		logger.Error("Error while connecting to database", err)
	} else {
		logger.Info("Successfully connected to database")
	}

	h := NewHandlerV1(&HandlerV1Config{
		Logger:   nil,
		Cfg:      cfg,
		Postgres: storage.New(db, logger, cfg),
	})

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(resp.Password), bcrypt.DefaultCost)
	if err != nil {
		panic("Error generating password")
	}
	resp.Password = string(hashedPassword)

	_, err = h.storage.Postgres().CreateAdmin(resp)
	if err != nil {
		panic("error creating admin")
	}
}
