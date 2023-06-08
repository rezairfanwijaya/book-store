package book

import (
	"book-store/entity"
)

type InputNewBook struct {
	Title         string        `json:"title"`
	PublishedYear string        `json:"pusblished_year"`
	ISBN          string        `json:"isbn"`
	Author        entity.Author `json:"author"`
}

type InputUpdateBook struct {
	Title         string        `json:"title"`
	PublishedYear string        `json:"pusblished_year"`
	ISBN          string        `json:"isbn"`
	Author        entity.Author `json:"author"`
}
