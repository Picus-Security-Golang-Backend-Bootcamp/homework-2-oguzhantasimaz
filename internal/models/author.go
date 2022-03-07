package models

type Author struct {
	name    string
	surname string
}

//Create new author
func NewAuthor(name, surname string) *Author {
	return &Author{name: name, surname: surname}
}
