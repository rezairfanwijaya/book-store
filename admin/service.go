package admin

import (
	"book-store/helper"
	"fmt"
	"net/http"
)

type IService interface {
	Login(input InputAdminLogin) (Admin, int, error)
}

type service struct {
	repoAdmin IRepository
}

func NewService(repoAdmin IRepository) *service {
	return &service{repoAdmin}
}

func (s *service) Login(input InputAdminLogin) (Admin, int, error) {
	// cari email yang diinput
	adminByEmail, err := s.repoAdmin.FindByEmail(input.Email)
	if err != nil {
		return adminByEmail, http.StatusInternalServerError, err
	}

	// jika email salah
	if adminByEmail.ID == 0 {
		return adminByEmail, http.StatusBadRequest, fmt.Errorf("email %v not found", input.Email)
	}

	// jika password salah
	if err := helper.VerifyPassword(input.Password, adminByEmail.Password); err != nil {
		return adminByEmail, http.StatusBadRequest, err
	}

	// sukses login
	return adminByEmail, http.StatusOK, nil
}
