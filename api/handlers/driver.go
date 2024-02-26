package handlers

import (
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/gin-gonic/gin"
)

// @Router		/user/edit-info [PATCH]
// @Summary		user modifying information
// @Tags        User
// @Description	Here users' info can be modified.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param       post   body       models.Driver true "admin"
// @Success		200 	{object}  models.Driver
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) UpdateDriverInfo(c *gin.Context) {
	var driver models.Driver

	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println("driver id: ", driver.Id)
	if (driver.Id) == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}

	data, err := h.storage.Postgres().UpdateDriverInfo(driver)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": data})
}
