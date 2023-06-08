package book

import (
	"book-store/entity"
	"fmt"
	"net/http"
)

type IService interface {
	Save(input InputNewBook) (entity.Book, int, error)
	GetAll() ([]entity.Book, int, error)
	Update(input InputUpdateBook, title string) (entity.Book, int, error)
	Delete(bookTitle string, author entity.Author) (int, error)
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

func (s *service) Update(input InputUpdateBook, title string) (entity.Book, int, error) {
	// cek apakah buku ada atau tidak
	bookByTitle, err := s.repoBook.FindByBookTitle(title)
	if err != nil {
		return bookByTitle, http.StatusInternalServerError, err
	}

	if bookByTitle.ID == 0 {
		return bookByTitle, http.StatusNotFound, fmt.Errorf("book %v not found", title)
	}

	// cek apakah user merupakan author buku atau bukan
	isAuthor := false
	for _, author := range bookByTitle.Authors {
		if input.Author.ID == author.ID {
			isAuthor = true
		}
	}

	if !isAuthor {
		return bookByTitle, http.StatusNotFound, fmt.Errorf("%v not author this book", input.Author.Name)
	}

	// binding
	bookByTitle.ISBN = input.ISBN
	bookByTitle.PublishedYear = input.PublishedYear
	bookByTitle.Title = input.Title

	// update
	bookUpdated, err := s.repoBook.Update(bookByTitle)
	if err != nil {
		return bookUpdated, http.StatusInternalServerError, err
	}

	return bookUpdated, http.StatusOK, nil
}

func (s *service) Delete(bookTitle string, author entity.Author) (int, error) {
	// cari buku berdasarkan title
	bookByTitle, err := s.repoBook.FindByBookTitle(bookTitle)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if bookByTitle.ID == 0 {
		return http.StatusNotFound, fmt.Errorf("book %v not found", bookTitle)
	}

	// apakah user merupakan author
	isAuthor := false
	for _, bookAuthor := range bookByTitle.Authors {
		if author.ID == bookAuthor.ID {
			isAuthor = true
		}
	}

	if !isAuthor {
		return http.StatusNotFound, fmt.Errorf("%v not author this book", author.Name)
	}

	// delete
	if err := s.repoBook.Delete(bookByTitle); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
