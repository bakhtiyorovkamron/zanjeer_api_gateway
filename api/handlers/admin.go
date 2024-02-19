package handlers

import (
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/gin-gonic/gin"
)

// @Router		/superadmin/add/admin [POST]
// @Summary		Add admin by superadmin
// @Tags        Admin
// @Description	Here admins can be created.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param       post   body       models.Admin true "admin"
// @Success		200 	{object}  models.Admin
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) CreateAdmin(c *gin.Context) {
	var resp models.Admin

	if err := c.ShouldBindJSON(&resp); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := h.storage.Postgres().CreateAdmin(resp)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"error":  err.Error()},
		)
		return
	}
	c.JSON(200, gin.H{
		"status":  "OK",
		"message": "Admin created successfully",
	})
}
