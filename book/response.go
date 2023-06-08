package book

import (
	"book-store/entity"
)

type responseNewBook struct {
	ID            int             `json:"id"`
	Title         string          `json:"title"`
	PublishedYear string          `json:"pusblished_year"`
	ISBN          string          `json:"isbn"`
	Authors       []entity.Author `json:"authors"`
}

func FormatterResponseNewBook(book entity.Book) *responseNewBook {
	return &responseNewBook{
		ID:            book.ID,
		Title:         book.Title,
		PublishedYear: book.PublishedYear,
		ISBN:          book.ISBN,
		Authors:       book.Authors,
	}
}
