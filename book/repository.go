package book

import (
	"book-store/entity"

	"gorm.io/gorm"
)

type IRepository interface {
	Save(book entity.Book) (entity.Book, error)
	FindByBookTitle(bookTitle string) (entity.Book, error)
	FindAll() ([]entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
	Delete(book entity.Book) error
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

func (r *repository) FindAll() ([]entity.Book, error) {
	var books []entity.Book
	if err := r.db.Model(&entity.Book{}).Preload("Authors").Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}

func (r *repository) Update(book entity.Book) (entity.Book, error) {
	if err := r.db.Save(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) Delete(book entity.Book) error {
	var singleBook entity.Book

	if err := r.db.Model(&book).Association("Authors").Clear(); err != nil {
		return err
	}

	if err := r.db.Where("title = ?", book.Title).Delete(&singleBook).Error; err != nil {
		return err
	}

	return nil
}
