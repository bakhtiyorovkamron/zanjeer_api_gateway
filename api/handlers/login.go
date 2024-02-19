package handlers

import (
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/gin-gonic/gin"
)

// @Router		/login [POST]
// @Summary		Login user
// @Tags        User
// @Description	Here user can be logged.
// @Security    BearerAuth
// @Accept      json
// @Produce		json
// @Param       post   body       models.Login true "login users"
// @Success		200 	{object}  models.Login
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) Login(c *gin.Context) {
	var resp models.Login

	if err := c.ShouldBindJSON(&resp); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	response, err := h.storage.Postgres().Login(resp)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Response from login", response)

	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   response,
	})
}
