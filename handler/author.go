package handler

import (
	"book-store/auth"
	"book-store/author"
	"book-store/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerAuthor struct {
	authorService author.IService
	authService   auth.Service
}

func NewHandlerAuthor(authorService author.IService, authService auth.Service) *handlerAuthor {
	return &handlerAuthor{authorService, authService}
}

func (h *handlerAuthor) Register(c *gin.Context) {
	// binding
	var input author.InputAuthorSession

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
	authorRegistered, httpCode, err := h.authorService.Register(input)
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
		authorRegistered,
	)

	c.JSON(httpCode, response)
}

func (h *handlerAuthor) Login(c *gin.Context) {
	// binding
	var input author.InputAuthorSession

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
	authorLoggedin, httpCode, err := h.authorService.Login(input)
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
	token, err := h.authService.GenerateToken(authorLoggedin.ID)
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
	authorFormatted := author.FormatterAuthorLogin(authorLoggedin, token)

	response := helper.ResponseAPI(
		"success",
		httpCode,
		authorFormatted,
	)

	c.JSON(httpCode, response)
}

func (h *handlerAuthor) CurrentAuthor(c *gin.Context) {
	currentAuthor := c.MustGet("currentAuthor").(author.Author)
	response := helper.ResponseAPI(
		"success",
		http.StatusOK,
		currentAuthor,
	)

	c.JSON(http.StatusOK, response)
}
