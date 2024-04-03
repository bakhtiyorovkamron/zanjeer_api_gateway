package handlers

import (
	"github.com/Projects/zanjeer_api_gateway/models/flespi"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) WebhookHandler(c *gin.Context) {
	var req flespi.WebHookResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	tunnel <- req
	c.JSON(200, gin.H{"status": "success"})
}
