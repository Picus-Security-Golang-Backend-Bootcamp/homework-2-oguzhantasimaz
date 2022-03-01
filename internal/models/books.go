package models

import "strconv"

//Book struct
type Book struct {
	id         int
	title      string
	stockCode  string
	stockCount int
	isbn       int
	pageCount  int
	price      float64
	isDeleted  bool
	author     *Author
}

//Create a new book
func NewBook(id int, title string, stockCode string, stockCount int, isbn int, pageCount int, price float64, isDeleted bool, author *Author) *Book {
	return &Book{id, title, stockCode, stockCount, isbn, pageCount, price, isDeleted, author}
}

//Get author name of the book
func (b *Book) AuthorName() string {
	return b.author.name + " " + b.author.surname
}

//Set stock count of the book
func (b *Book) SetStockCount(count int) {
	b.stockCount = count
}

//Set isDeleted of the book
func (b *Book) SetIsDeleted(isDeleted bool) {
	b.isDeleted = isDeleted
}

//Return all information of the book as a string
func (b *Book) GetBookInformation() string {
	return "ID: " + strconv.Itoa(b.id) + "\n" +
		"Title: " + b.title + "\n" +
		"Stock Code: " + b.stockCode + "\n" +
		"Stock Count: " + strconv.Itoa(b.stockCount) + "\n" +
		"ISBN: " + strconv.Itoa(b.isbn) + "\n" +
		"Page Count: " + strconv.Itoa(b.pageCount) + "\n" +
		"Price: " + strconv.FormatFloat(b.price, 'f', 2, 64) + "\n" +
		"Author: " + b.AuthorName() + "\n" +
		"Is Deleted: " + strconv.FormatBool(b.isDeleted)
}
