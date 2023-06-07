package admin

type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
