package handlers

import (
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/gin-gonic/gin"
)

// @Router		/user/edit-info [PATCH]
// @Summary		user modifying information
// @Tags        Driver
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

// @Router /user/{id} [GET]
// @Summary Get info about driver
// @Tags Driver
// @Description Here driver's info can be fetched.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "id"
// @Success 200 {object} models.Driver
// @Failure default {object} models.StandardResponse
func (h *handlerV1) GetDriverInfo(c *gin.Context) {

	driverId := c.Param("id")

	data, err := h.storage.Postgres().GetDriverInfo(driverId)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "error",
			"data":   nil,
		})
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   data,
	})
}

// @Router		/user/{id} [DELETE]
// @Summary		user delete
// @Tags        Driver
// @Description	Here drivers can be deleted.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param 		id 		path 	 string true "id"
// @Success 	200 	{object} models.StandardResponse
// @Failure 	default {object} models.StandardResponse
func (h *handlerV1) DeleteDriverInfo(c *gin.Context) {
	driverId := c.Param("id")

	if driverId == "" {
		c.JSON(400, gin.H{
			"error": "driver id is required",
		})
		return
	}

	if err := h.storage.Postgres().DeleteDriver(driverId); err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, gin.H{"status": "success"})
}

// @Router /user/get-list [GET]
// @Summary Get info about all drivers driver
// @Tags Driver
// @Description Here drivers' info can be fetched.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Success 200 {object} models.Driver
// @Failure default {object} models.StandardResponse
func (h *handlerV1) GetDriversList(c *gin.Context) {

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error":  err,
			"status": "error",
		})
		return
	}
	offset, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error":  err,
			"status": "error",
		})
		return
	}

	data, err := h.storage.Postgres().GetDriverList(int64(limit), int64(offset))
	if err != nil {
		c.JSON(400, gin.H{
			"error":  err,
			"status": "error",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   data,
	})
}

// @Router /user/search [POST]
// @Summary Get info about driver
// @Tags Driver
// @Description Here driver's info can be fetched by name and phone
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param  search  body models.DriverSearchRequest true "search driver"
// @Success 200 {object} []models.Driver
// @Failure default {object} models.StandardResponse
func (h *handlerV1) GetDriversSearch(c *gin.Context) {
	var res models.DriverSearchRequest
	if err := c.ShouldBindJSON(&res); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	if res.Limit == 0 {
		res.Limit = 1
	}
	if res.Offset == 0 {
		res.Offset = 1
	}
	data, err := h.storage.Postgres().SearchDriver(res)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "ok",
		"data":   data,
	})
}
