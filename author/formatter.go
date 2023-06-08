package author

type responseAuthorLogin struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Token   string `json:"token"`
}

func FormatterAuthorLogin(author Author, token string) *responseAuthorLogin {
	return &responseAuthorLogin{
		Name:    author.Name,
		Country: author.Country,
		Token:   token,
	}
}
