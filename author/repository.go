package author

import "gorm.io/gorm"

type IRepository interface {
	Save(author Author) (Author, error)
	FindByName(name string) (Author, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(author Author) (Author, error) {
	if err := r.db.Create(&author).Error; err != nil {
		return author, err
	}

	return author, nil
}

func (r *repository) FindByName(name string) (Author, error) {
	var author Author
	if err := r.db.Where("name = ?", name).Find(&author).Error; err != nil {
		return author, err
	}

	return author, nil
}
