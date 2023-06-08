package book

import (
	"book-store/author"
	"net/http"
)

type IService interface {
	Save(input InputNewBook) (Book, int, error)
}

type service struct {
	repoBook IRepository
}

func NewService(repoBook IRepository) *service {
	return &service{repoBook}
}

func (s *service) Save(input InputNewBook) (Book, int, error) {
	// binding
	book := Book{
		Title:         input.Title,
		PublishedYear: input.PublishedYear,
		ISBN:          input.ISBN,
		Author: []author.Author{
			input.Author,
		},
	}

	bookSaved, err := s.repoBook.Save(book)
	if err != nil {
		return bookSaved, http.StatusInternalServerError, err
	}

	return bookSaved, http.StatusOK, nil
}
