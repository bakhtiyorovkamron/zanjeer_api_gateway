package handlers

import (
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/gin-gonic/gin"
)

// @Router		/user/send-number [POST]
// @Summary		user sending number
// @Tags        User
// @Description	Here users can be registered.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param       post   body       models.UserRegister true "admin"
// @Success		200 	{object}  models.UserRegister
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) SendNumber(c *gin.Context) {
	var req models.UserRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error while binding request", err)
		c.JSON(400, models.StandardResponse{
			Status:  "error",
			Message: "Invalid request",
		})
		return
	}

	c.JSON(200, models.StandardResponse{
		Status:  "success",
		Message: "User registered successfully",
		Data:    req,
	})
}
