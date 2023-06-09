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

type responseBook struct {
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

func FormatterResponseBook(book entity.Book) *responseBook {
	return &responseBook{
		ID:            book.ID,
		Title:         book.Title,
		PublishedYear: book.PublishedYear,
		ISBN:          book.ISBN,
		Authors:       book.Authors,
	}
}

func FormatterResponseBooks(books []entity.Book) []*responseBook {
	var result []*responseBook

	for _, book := range books {
		bookFormatted := FormatterResponseBook(book)
		result = append(result, bookFormatted)
	}

	return result
}
