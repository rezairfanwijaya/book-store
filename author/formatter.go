package author

import "book-store/entity"

type responseAuthorLogin struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Token   string `json:"token"`
}

type responseCurrentAuthor struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func FormatterAuthorLogin(author entity.Author, token string) *responseAuthorLogin {
	return &responseAuthorLogin{
		Name:    author.Name,
		Country: author.Country,
		Token:   token,
	}
}

func FormatterCurrentAuthor(author entity.Author) *responseCurrentAuthor {
	return &responseCurrentAuthor{
		ID:      author.ID,
		Name:    author.Name,
		Country: author.Country,
	}
}

func FormatterCurrentAuthors(authors []entity.Author) []*responseCurrentAuthor {
	var result []*responseCurrentAuthor

	for _, author := range authors {
		singleAuthor := FormatterCurrentAuthor(author)
		result = append(result, singleAuthor)
	}

	return result
}
