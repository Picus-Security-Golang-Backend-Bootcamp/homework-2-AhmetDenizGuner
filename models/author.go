package models

type Author struct {
	Name string
}

func NewAuthor(authorName string) *Author {

	author := &Author{
		Name: authorName,
	}
	return author
}
