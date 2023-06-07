package handler

import (
	"book-store/admin"
	"book-store/auth"
	"book-store/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerAdmin struct {
	adminService admin.IService
	authService  auth.Service
}

func NewHandlerAdmin(adminService admin.IService, authService auth.Service) *handlerAdmin {
	return &handlerAdmin{adminService, authService}
}

func (h *handlerAdmin) Login(c *gin.Context) {
	// binding
	var input admin.InputAdminLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		errBinding := helper.ErrorFormater(err)
		response := helper.ResponseAPI(
			"failed",
			http.StatusBadRequest,
			errBinding,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// service
	adminLoggedin, httpCode, err := h.adminService.Login(input)
	if err != nil {
		response := helper.ResponseAPI(
			"failed",
			httpCode,
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	// generate token jwt
	token, err := h.authService.GenerateToken(adminLoggedin.ID)
	if err != nil {
		response := helper.ResponseAPI(
			"failed generate token",
			httpCode,
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	// format admin
	adminFormatted := admin.FormatterResponseAdmin(adminLoggedin, token)

	response := helper.ResponseAPI(
		"success",
		httpCode,
		adminFormatted,
	)

	c.JSON(httpCode, response)
}
