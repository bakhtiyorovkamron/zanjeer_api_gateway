package handlers

import (
	"net/http"

	"github.com/Projects/zanjeer_api_gateway/pkg/validator"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := validator.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
