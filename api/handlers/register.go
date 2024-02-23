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
		c.JSON(500, models.StandardResponse{
			Status:  "error",
			Message: "Error while generating code",
		})
		return
	}

	req.Phone = util.FormatPhone(req.Phone)
	// send code to user
	otp, err := h.storage.Postgres().CreateOTP(models.SmsOtp{
		Phone: req.Phone,
		Code:  code,
	})
	if err != nil {
		c.JSON(500, models.StandardResponse{
			Status:  "error",
			Message: "Error while sending code",
		})
		return
	}

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
// @Param       post   body       models.Sms true "admin"
// @Success		200 	{object}  models.Sms
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) VerifyNumber(c *gin.Context) {
	var req models.Sms
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error while binding request", err)
		c.JSON(400, models.StandardResponse{
			Status:  "error",
			Message: "Invalid request",
		})
		return
	}

	// if req.Code != "1234" {
	// 	c.JSON(400, models.StandardResponse{
	// 		Status:  "error",
	// 		Message: "Invalid code",
	// 	})
	// 	return
	// }
	otp := models.ConfirmOTP{
		SmsId: req.SmsId,
		Code:  req.Code,
	}
	phone := "string"
	otp.Phone = &phone
	err := h.storage.Postgres().ConfirmOTP(otp)
	if err != nil {
		c.JSON(500, models.StandardResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	resp, err := h.storage.Postgres().CreateDriver(models.Driver{
		Phone: *otp.Phone,
	})
	if err != nil {
		c.JSON(500, models.StandardResponse{
			Status:  "error",
			Message: "Error while creating driver",
		})
		return
	}
	token, err := validator.GenerateToken(resp.Id, "driver")
	if err != nil {

	}
	c.JSON(200, gin.H{
		"message": "User verified successfully",
		"token":   token,
	})
}
