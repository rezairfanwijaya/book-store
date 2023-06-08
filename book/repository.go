package book

import (
	"book-store/entity"

	"gorm.io/gorm"
)

type IRepository interface {
	Save(book entity.Book) (entity.Book, error)
	FindByBookTitle(bookTitle string) (entity.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(book entity.Book) (entity.Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) FindByBookTitle(bookTitle string) (entity.Book, error) {
	var book entity.Book
	if err := r.db.Model(&entity.Book{}).Preload("Authors").Where("title = ?", bookTitle).Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}
