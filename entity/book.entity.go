package entity

type Book struct {
	ID            int      `json:"id" gorm:"primaryKey"`
	Title         string   `json:"title"`
	PublishedYear string   `json:"pusblished_year"`
	ISBN          string   `json:"isbn"`
	Authors       []Author `gorm:"many2many:author_books;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
