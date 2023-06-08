package book

import (
	"book-store/entity"
	"net/http"
)

type IService interface {
	Save(input InputNewBook) (entity.Book, int, error)
}

type service struct {
	repoBook IRepository
}

func NewService(repoBook IRepository) *service {
	return &service{repoBook}
}

func (s *service) Save(input InputNewBook) (entity.Book, int, error) {
	// mapping
	book := entity.Book{
		Title:         input.Title,
		PublishedYear: input.PublishedYear,
		ISBN:          input.ISBN,
		Authors: []entity.Author{
			input.Author,
		},
	}

	bookSaved, err := s.repoBook.Save(book)
	if err != nil {
		return bookSaved, http.StatusInternalServerError, err
	}

	return bookSaved, http.StatusOK, nil
}
