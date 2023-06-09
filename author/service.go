package author

import (
	"book-store/entity"
	"fmt"
	"net/http"
)

type IService interface {
	Register(input InputAuthorSession) (entity.Author, int, error)
	Login(input InputAuthorSession) (entity.Author, int, error)
	GetByName(name string) (entity.Author, int, error)
	GetByID(id int) (entity.Author, int, error)
	GetAll() ([]entity.Author, int, error)
}

type service struct {
	repoAtuhor IRepository
}

func NewService(repoAtuhor IRepository) *service {
	return &service{repoAtuhor}
}

func (s *service) Register(input InputAuthorSession) (entity.Author, int, error) {
	// find by name
	authorByName, err := s.repoAtuhor.FindByName(input.Name)
	if err != nil {
		return authorByName, http.StatusInternalServerError, err
	}

	if authorByName.ID != 0 {
		return authorByName, http.StatusBadRequest, fmt.Errorf("author %v already registered", input.Name)
	}

	// binding
	newAuthor := entity.Author{
		Name:    input.Name,
		Country: input.Country,
	}

	// save
	authorRegistered, err := s.repoAtuhor.Save(newAuthor)
	if err != nil {
		return authorRegistered, http.StatusInternalServerError, err
	}

	return authorRegistered, http.StatusOK, nil
}

func (s *service) GetByName(name string) (entity.Author, int, error) {
	// find by name
	authorByName, err := s.repoAtuhor.FindByName(name)
	if err != nil {
		return authorByName, http.StatusInternalServerError, err
	}

	if authorByName.ID == 0 {
		return authorByName, http.StatusBadRequest, fmt.Errorf("author %v not found", name)
	}

	return authorByName, http.StatusOK, nil
}

func (s *service) Login(input InputAuthorSession) (entity.Author, int, error) {
	// get by name
	authorByName, err := s.repoAtuhor.FindByName(input.Name)
	if err != nil {
		return authorByName, http.StatusInternalServerError, err
	}

	if authorByName.ID == 0 {
		return authorByName, http.StatusBadRequest, fmt.Errorf("author %v not registered", input.Name)
	}

	// cek country
	if input.Country != authorByName.Country {
		return authorByName, http.StatusBadRequest, fmt.Errorf("country not match")
	}

	return authorByName, http.StatusOK, nil
}

func (s *service) GetByID(id int) (entity.Author, int, error) {
	authorByID, err := s.repoAtuhor.FindByID(id)
	if err != nil {
		return authorByID, http.StatusInternalServerError, err
	}

	if authorByID.ID == 0 {
		return authorByID, http.StatusBadRequest, fmt.Errorf("id %v not found", id)
	}

	return authorByID, http.StatusOK, nil
}

func (s *service) GetAll() ([]entity.Author, int, error) {
	authors, err := s.repoAtuhor.FindAll()
	if err != nil {
		return authors, http.StatusInternalServerError, err
	}

	return authors, http.StatusOK, nil
}
