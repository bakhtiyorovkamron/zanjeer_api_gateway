package handlers

import (
	"github.com/Projects/zanjeer_api_gateway/models"
	"github.com/Projects/zanjeer_api_gateway/pkg/util"
	"github.com/Projects/zanjeer_api_gateway/pkg/validator"
	"github.com/gin-gonic/gin"
)

// @Router		/user/send-number [POST]
// @Summary		user sending number
// @Tags        Driver
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
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "Invalid request",
			Data:    nil,
			Code:    400,
		})
		return
	}
	valid := validator.IsValidPhone(req.Phone)
	if !valid {
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "Invalid phone number",
			Data:    nil,
			Code:    400,
		})
		return
	}

	req.Phone = util.FormatPhone(req.Phone)

	exists, err := h.storage.Postgres().SearchDriver(models.DriverSearchRequest{
		Phone:  req.Phone,
		Limit:  1,
		Offset: 1,
	})
	if err != nil {
		h.log.Error("Error while checking phone exists in db", err)
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "Internal error",
			Data:    nil,
			Code:    500,
		})
		return
	}
	if exists.Count == 0 {
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "User Not Found",
			Data:    nil,
			Code:    404,
		})
		return
	}

	code, err := util.GenerateCode(4)
	if err != nil {
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "Error while generating code",
			Data:    nil,
			Code:    500,
		})
		return
	}

	// send code to user
	otp, err := h.storage.Postgres().CreateOTP(models.SmsOtp{
		Phone: req.Phone,
		Code:  code,
	})
	if err != nil {
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "Error while sending code",
			Data:    nil,
			Code:    500,
		})
		return
	}

	h.handleResponse(c, models.StandardResponse{
		Status: "success",
		Data:   otp,
		Code:   200,
	})
}

// @Router		/user/verify-number [POST]
// @Summary		user verifying number
// @Tags        Driver
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
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "Bad request",
			Data:    nil,
			Code:    400,
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
	// otp := models.ConfirmOTP{
	// 	SmsId: req.SmsId,
	// 	Code:  req.Code,
	// }
	phone := ""

	err := h.storage.Postgres().ConfirmOTP(models.ConfirmOTP{
		SmsId: req.SmsId,
		Code:  req.Code,
		Phone: &phone,
	})
	if err != nil {
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
			Code:    502,
		})
		return
	}
	resp, err := h.storage.Postgres().VerifyDriver(phone)
	if err != nil {
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "Error",
			Data:    nil,
			Code:    502,
		})
		return
	}
	token, err := validator.GenerateToken(resp.Id, "driver")
	if err != nil {
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "Error",
			Data:    nil,
			Code:    502,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User verified successfully",
		"token":   token,
	})
}

// @Router		/user/register [POST]
// @Summary		user registeration
// @Tags        Driver
// @Description	Here users can be registered.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param       post   body       models.Driver true "driver"
// @Success		200 	{object}  models.Driver
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) UserRegister(c *gin.Context) {
	var req models.Driver
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Error("Error while binding request", err)
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "Bad request",
			Data:    nil,
			Code:    400,
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

	req.Phone = util.FormatPhone(req.Phone)

	resp, err := h.storage.Postgres().CreateDriver(req)
	if err != nil {
		h.handleResponse(c, models.StandardResponse{
			Status:  "error",
			Message: "Error",
			Data:    nil,
			Code:    502,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User registreted successfully",
		"data":    resp,
	})
}
