package models

//Book struct
type Book struct {
	title string
}

//Get title of the book
func (b *Book) Title() string {
	return b.title
}
