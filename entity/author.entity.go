package entity

type Author struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Books   []Book `json:"-" gorm:"many2many:author_books;"`
}
