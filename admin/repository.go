package admin

import "gorm.io/gorm"

type IRepository interface {
	FindByEmail(email string) (Admin, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByEmail(email string) (Admin, error) {
	var admin Admin
	if err := r.db.Where("email = ?", email).Find(&admin).Error; err != nil {
		return admin, err
	}

	return admin, nil
}
