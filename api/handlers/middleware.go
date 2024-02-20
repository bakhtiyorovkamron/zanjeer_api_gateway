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
func (h *handlerV1) SuperAdminCheckType() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType := validator.GetUserTypeFromToken(c)
		if userType == "" {
			c.String(http.StatusUnauthorized, "user type is empty")
			c.Abort()
			return
		}
		if userType != "superadmin" {
			c.String(http.StatusUnauthorized, "user is not superadmin")
			c.Abort()
			return
		}
		c.Next()
	}
}
