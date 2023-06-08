package author

type InputAuthorSession struct {
	Name    string `json:"name" binding:"required,min=6"`
	Country string `json:"country" binding:"required"`
}
