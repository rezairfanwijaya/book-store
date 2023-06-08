package author

type Author struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Country string `json:"country"`
}
