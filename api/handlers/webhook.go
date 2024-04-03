package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) WebhookHandler(c *gin.Context) {
	req := make(map[string]interface{})

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	tunnel <- req
	c.JSON(200, gin.H{"status": "success"})
}
