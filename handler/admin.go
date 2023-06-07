package handler

import (
	"book-store/admin"
	"book-store/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerAdmin struct {
	adminService admin.IService
}

func NewHandlerAdmin(adminService admin.IService) *handlerAdmin {
	return &handlerAdmin{adminService}
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

	response := helper.ResponseAPI(
		"success",
		httpCode,
		adminLoggedin,
	)

	c.JSON(httpCode, response)
}
