package book

import (
	"book-store/entity"
	"fmt"
	"net/http"
)

type IService interface {
	Save(input InputNewBook) (entity.Book, int, error)
	GetAll() ([]entity.Book, int, error)
}

type service struct {
	repoBook IRepository
}

func NewService(repoBook IRepository) *service {
	return &service{repoBook}
}

func (s *service) Save(input InputNewBook) (entity.Book, int, error) {
	// get book by title
	bookByTitle, err := s.repoBook.FindByBookTitle(input.Title)
	if err != nil {
		return bookByTitle, http.StatusInternalServerError, err
	}

	// cek apakah author sama dengan current user
	// untuk menghindari duplikasi data
	for _, author := range bookByTitle.Authors {
		if author.ID == input.Author.ID {
			return bookByTitle, http.StatusBadRequest, fmt.Errorf("duplicate author %v on book %v", input.Author.Name, input.Title)
		}
	}

	if bookByTitle.ID != 0 {
		// jika book sudah ada maka tambahkan saja authornya
		bookByTitle.Authors = append(bookByTitle.Authors, input.Author)
		bookUpdated, err := s.repoBook.Update(bookByTitle)
		if err != nil {
			return bookUpdated, http.StatusInternalServerError, err
		}

		return bookUpdated, http.StatusOK, err
	}

	// mapping
	book := entity.Book{
		Title:         input.Title,
		PublishedYear: input.PublishedYear,
		ISBN:          input.ISBN,
		Authors: []entity.Author{
			input.Author,
		},
	}

	// simpan
	bookSaved, err := s.repoBook.Save(book)
	if err != nil {
		return bookSaved, http.StatusInternalServerError, err
	}

	return bookSaved, http.StatusOK, nil
}

func (s *service) GetAll() ([]entity.Book, int, error) {
	books, err := s.repoBook.FindAll()
	if err != nil {
		return books, http.StatusInternalServerError, err
	}

	return books, http.StatusOK, err
}
