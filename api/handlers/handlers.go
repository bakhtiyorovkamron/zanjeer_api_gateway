package handlers

import (
	"strconv"

	"github.com/Projects/zanjeer_api_gateway/config"
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/Projects/zanjeer_api_gateway/models/flespi"
	"github.com/Projects/zanjeer_api_gateway/pkg/logger"
	"github.com/Projects/zanjeer_api_gateway/storage"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log     *logger.Logger
	cfg     config.Config
	storage storage.StorageI
}
type DataFromFlespiWebhook struct {
	IsNew bool
	Data  map[string]interface{}
}

var tunnel = make(chan flespi.WebHookResponse)

var data *DataFromFlespiWebhook

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
func (h *handlerV1) handleResponse(c *gin.Context, info models.StandardResponse) bool {
	switch code := info.Code; {
	case code >= 400 && code < 500:
		c.JSON(code, gin.H{
			"status":  "error",
			"data":    nil,
			"message": info.Message,
		})

	case code >= 500 && code < 600:
		c.JSON(code, gin.H{
			"status":  "error",
			"data":    nil,
			"message": info.Message,
		})
	case code >= 200 && code < 300:
		c.JSON(code, gin.H{
			"status":  "success",
			"data":    info.Data,
			"message": info.Message,
		})
	}

	return false
}

func ParseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", "10"))
}

func ParsePageQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("offset", "1"))
}
