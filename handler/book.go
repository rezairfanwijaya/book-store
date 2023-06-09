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
	// ambil info author yang create new book
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

func (h *handlerBook) GetAll(c *gin.Context) {
	books, httpCode, err := h.bookService.GetAll()
	if err != nil {
		response := helper.ResponseAPI(
			"failed",
			httpCode,
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	// formatter
	booksFormatted := book.FormatterResponseBooks(books)

	response := helper.ResponseAPI(
		"success",
		httpCode,
		booksFormatted,
	)

	c.JSON(httpCode, response)
}

func (h *handlerBook) Update(c *gin.Context) {
	title := c.Param("title")

	if title == "" {
		response := helper.ResponseAPI(
			"failed",
			http.StatusBadRequest,
			"param not be empty",
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil info author yang create new book
	currentAuthor := c.MustGet("currentAuthor").(entity.Author)

	// input binding
	var input book.InputUpdateBook
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
	bookUpdated, httpCode, err := h.bookService.Update(input, title)
	if err != nil {
		response := helper.ResponseAPI(
			"failed",
			httpCode,
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	// formatter
	bookUpdatedFormatted := book.FormatterResponseBook(bookUpdated)
	response := helper.ResponseAPI(
		"success",
		httpCode,
		bookUpdatedFormatted,
	)

	c.JSON(httpCode, response)
}

func (h *handlerBook) Delete(c *gin.Context) {
	title := c.Param("title")

	if title == "" {
		response := helper.ResponseAPI(
			"failed",
			http.StatusBadRequest,
			"param not be empty",
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil info author yang create new book
	currentAuthor := c.MustGet("currentAuthor").(entity.Author)

	// service
	httpCode, err := h.bookService.Delete(title, currentAuthor)
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
		"success deleted",
	)

	c.JSON(httpCode, response)
}

func (h *handlerBook) GetByTitle(c *gin.Context) {
	bookTitle := c.Param("title")

	bookByTitle, httpCode, err := h.bookService.GetByBookTitle(bookTitle)
	if err != nil {
		response := helper.ResponseAPI(
			"failed",
			httpCode,
			err.Error(),
		)

		c.JSON(httpCode, response)
		return
	}

	// formatter
	bookByTitleFormatted := book.FormatterResponseBook(bookByTitle)
	response := helper.ResponseAPI(
		"success",
		httpCode,
		bookByTitleFormatted,
	)

	c.JSON(httpCode, response)
}
