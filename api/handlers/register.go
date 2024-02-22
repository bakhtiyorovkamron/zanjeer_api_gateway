package handlers

import (
	"fmt"

	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/Projects/zanjeer_api_gateway/pkg/util"
	"github.com/Projects/zanjeer_api_gateway/pkg/validator"
	"github.com/gin-gonic/gin"
)

// @Router		/user/send-number [POST]
// @Summary		user sending number
// @Tags        User
// @Description	Here users can be registered.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param       post   body       models.Sms true "admin"
// @Success		200 	{object}  models.Sms
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) SendNumber(c *gin.Context) {
	var req models.Sms
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error while binding request", err)
		c.JSON(400, models.StandardResponse{
			Status:  "error",
			Message: "Invalid request",
		})
		return
	}
	valid := validator.IsValidPhone(req.Phone)
	if !valid {
		c.JSON(400, models.StandardResponse{
			Status:  "error",
			Message: "Invalid phone number",
		})
		return
	}

	code, err := util.GenerateCode(4)
	if err != nil {
		return
	}

	fmt.Println("Code", code)
	// send code to user
	otp, err := h.storage.Postgres().CreateOTP(models.SmsOtp{
		Phone: req.Phone,
		Code:  code,
	})

	c.JSON(200, models.StandardResponse{
		Status:  "success",
		Message: "User registered successfully",
		Data:    otp,
	})
}

// @Router		/user/verify-number [POST]
// @Summary		user verifying number
// @Tags        User
// @Description	Here users can be registered.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param       post   body       models.VerifyNumber true "admin"
// @Success		200 	{object}  models.VerifyNumber
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) VerifyNumber(c *gin.Context) {
	var req models.VerifyNumber
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error while binding request", err)
		c.JSON(400, models.StandardResponse{
			Status:  "error",
			Message: "Invalid request",
		})
		return
	}

	if req.Code != "1234" {
		c.JSON(400, models.StandardResponse{
			Status:  "error",
			Message: "Invalid code",
		})
		return
	}

	c.JSON(200, models.StandardResponse{
		Status:  "success",
		Message: "User registered successfully",
		Data:    req,
	})
}
