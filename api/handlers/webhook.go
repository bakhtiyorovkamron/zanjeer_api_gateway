package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) WebhookHandler(c *gin.Context) {
	req := make(map[string]interface{})

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	r, _ := json.MarshalIndent(req, " ", " ")
	fmt.Println("REQUEST FROM WEBHOOK: ", string(r))
}
