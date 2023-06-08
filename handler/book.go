package handler

import (
	"book-store/book"
	"book-store/entity"
	"book-store/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerBook struct {
	bookService book.IService
}

func NewHandlerBook(bookService book.IService) *handlerBook {
	return &handlerBook{bookService}
}

func (h *handlerBook) Create(c *gin.Context) {
	// ambil infor author yang create new book
	currentAuthor := c.MustGet("currentAuthor").(entity.Author)

	// input binding
	var input book.InputNewBook
	input.Author = currentAuthor

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
	bookSaved, httpCode, err := h.bookService.Save(input)
	if err != nil {
		response := helper.ResponseAPI(
			"failed",
			httpCode,
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	// formatting
	bookSavedFormatted := book.FormatterResponseNewBook(bookSaved)

	response := helper.ResponseAPI(
		"success",
		httpCode,
		bookSavedFormatted,
	)

	c.JSON(httpCode, response)
}
