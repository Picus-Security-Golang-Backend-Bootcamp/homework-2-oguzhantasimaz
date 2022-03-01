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
	author     *Author
	isDeleted  bool
}

//Get title of the book
func (b *Book) Title() string {
	return b.title
}

//Get page count of the book
func (b *Book) PageCount() int {
	return b.pageCount
}

//Get price of the book
func (b *Book) Price() float64 {
	return b.price
}

//Get ISBN of the book
func (b *Book) ISBN() int {
	return b.isbn
}

//Get stock count of the book
func (b *Book) StockCount() int {
	return b.stockCount
}

//Get author of the book
func (b *Book) Author() *Author {
	return b.author
}

//Get id of the book
func (b *Book) ID() int {
	return b.id
}

//Get isDeleted of the book
func (b *Book) IsDeleted() bool {
	return b.isDeleted
}

//Get author name of the book
func (b *Book) AuthorName() string {
	return b.author.name + " " + b.author.surname
}

//Get Stock Code of the book
func (b *Book) StockCode() string {
	return b.stockCode
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
