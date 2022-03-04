package models

type Author struct {
	Name string
}

//struct constructor
func NewAuthor(authorName string) *Author {

	author := &Author{
		Name: authorName,
	}
	return author
}
