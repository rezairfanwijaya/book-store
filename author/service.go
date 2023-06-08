package author

import (
	"fmt"
	"net/http"
)

type IService interface {
	Register(input InputAuthorSession) (Author, int, error)
	Login(input InputAuthorSession) (Author, int, error)
	GetByName(name string) (Author, int, error)
	GetByID(id int) (Author, int, error)
}

type service struct {
	repoAtuhor IRepository
}

func NewService(repoAtuhor IRepository) *service {
	return &service{repoAtuhor}
}

func (s *service) Register(input InputAuthorSession) (Author, int, error) {
	// get by name
	authorByName, httpCode, err := s.GetByName(input.Name)
	if err != nil {
		return authorByName, httpCode, err
	}

	// binding
	newAuthor := Author{
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

func (s *service) GetByName(name string) (Author, int, error) {
	// find by name
	authorByName, err := s.repoAtuhor.FindByName(name)
	if err != nil {
		return authorByName, http.StatusInternalServerError, err
	}

	if authorByName.ID != 0 {
		return authorByName, http.StatusBadRequest, fmt.Errorf("author %v already registered", name)
	}

	return authorByName, http.StatusOK, nil
}

func (s *service) Login(input InputAuthorSession) (Author, int, error) {
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

func (s *service) GetByID(id int) (Author, int, error) {
	authorByID, err := s.repoAtuhor.FindByID(id)
	if err != nil {
		return authorByID, http.StatusInternalServerError, err
	}

	if authorByID.ID == 0 {
		return authorByID, http.StatusBadRequest, fmt.Errorf("id %v not found", id)
	}

	return authorByID, http.StatusOK, nil
}
