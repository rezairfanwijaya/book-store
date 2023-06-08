package author

import (
	"book-store/entity"

	"gorm.io/gorm"
)

type IRepository interface {
	Save(author entity.Author) (entity.Author, error)
	FindByName(name string) (entity.Author, error)
	FindByID(id int) (entity.Author, error)
	FindAll() ([]entity.Author, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(author entity.Author) (entity.Author, error) {
	if err := r.db.Create(&author).Error; err != nil {
		return author, err
	}

	return author, nil
}

func (r *repository) FindByName(name string) (entity.Author, error) {
	var author entity.Author
	if err := r.db.Where("name = ?", name).Find(&author).Error; err != nil {
		return author, err
	}

	return author, nil
}

func (r *repository) FindByID(id int) (entity.Author, error) {
	var author entity.Author
	if err := r.db.Where("id = ?", id).Find(&author).Error; err != nil {
		return author, err
	}

	return author, nil
}

func (r *repository) FindAll() ([]entity.Author, error) {
	var authors []entity.Author
	if err := r.db.Find(&authors).Error; err != nil {
		return authors, err
	}

	return authors, nil
}
