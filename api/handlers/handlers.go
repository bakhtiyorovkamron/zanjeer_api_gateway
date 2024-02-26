package handlers

import (
	"github.com/Projects/zanjeer_api_gateway/config"
	"github.com/Projects/zanjeer_api_gateway/pkg/logger"
	"github.com/Projects/zanjeer_api_gateway/storage"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log     *logger.Logger
	cfg     config.Config
	storage storage.StorageI
}

// NewHandlerV1 is a constructor for handlerV1
type HandlerV1Config struct {
	Logger   *logger.Logger
	Cfg      config.Config
	Postgres storage.StorageI
}

func NewHandlerV1(h *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:     h.Logger,
		cfg:     h.Cfg,
		storage: h.Postgres,
	}
}
func (h *handlerV1) HandleResponse(c *gin.Context, err error) bool {
	return false
}
