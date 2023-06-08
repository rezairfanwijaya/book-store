package book

import "book-store/author"

type InputNewBook struct {
	Title         string        `json:"title"`
	PublishedYear string        `json:"pusblished_year"`
	ISBN          string        `json:"isbn"`
	Author        author.Author `json:"author"`
}
