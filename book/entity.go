package book

import "book-store/author"

type Book struct {
	ID            int             `json:"id" gorm:"primaryKey"`
	Title         string          `json:"title"`
	PublishedYear string          `json:"pusblished_year"`
	ISBN          string          `json:"isbn"`
	Author        []author.Author `gorm:"many2many:author_books;"`
}
