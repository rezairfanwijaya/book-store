package book

import "book-store/author"

type responseNewBook struct {
	ID            int             `json:"id"`
	Title         string          `json:"title"`
	PublishedYear string          `json:"pusblished_year"`
	ISBN          string          `json:"isbn"`
	Author        []author.Author `json:"author"`
}

func FormatterResponseNewBook(book Book) *responseNewBook {
	return &responseNewBook{
		ID:            book.ID,
		Title:         book.Title,
		PublishedYear: book.PublishedYear,
		ISBN:          book.ISBN,
		Author:        book.Author,
	}
}
