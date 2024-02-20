package handlers

import "github.com/gin-gonic/gin"

// @Router		/user/register [POST]
// @Summary		user register
// @Tags        User
// @Description	Here users can be registered.
// @Accept      json
// @Produce		json
// @Security    BearerAuth
// @Param       post   body       models.UserRegister true "admin"
// @Success		200 	{object}  models.UserRegister
// @Failure     default {object}  models.StandardResponse
func (h *handlerV1) UserRegister(c *gin.Context) {

}
