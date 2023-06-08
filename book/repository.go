package book

import "gorm.io/gorm"

type IRepository interface {
	Save(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(book Book) (Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}
