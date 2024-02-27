package handlers

import (
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/gin-gonic/gin"
)

// @Router 		/device/create-device-type [POST]
// @Summary		Device types
// @Tags        GPS Device Type
// @Description	Device types can be created
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param       post   body       models.DeviceType true "admin"
// @Success		200 	{object}  models.DeviceType
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) CreateDeviceType(c *gin.Context) {
	var req models.DeviceType

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, err := h.storage.Postgres().CreateDeviceType(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success", "data": data})
}

// @Router /device/get-list-device-type [GET]
// @Summary Get device list
// @Tags  GPS Device Type
// @Description	Device types can be fetched from
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name  query string false "name"
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} []models.DeviceType
// @Failure default {object} models.StandardResponse
func (h *handlerV1) GetDeviceTypeList(c *gin.Context) {
	data, err := h.storage.Postgres().GetDeviceTypeList(models.GetDeviceTypeListRequest{
		Name: c.Query("name"),
	})
	if err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "ok",
		"data":   data,
	})
}
