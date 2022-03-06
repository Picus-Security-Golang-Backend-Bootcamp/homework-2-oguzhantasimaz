package models

import (
	"fmt"
	"strconv"
)

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
	if !b.isDeleted {
		return b.author.name + " " + b.author.surname
	} else {
		fmt.Println("You can't get author name of a deleted book")
		return "Deleted book"
	}
}

//Set stock count of the book
func (b *Book) SetStockCount(count int) {
	if !b.isDeleted {
		b.stockCount = count
	} else {
		fmt.Println("You can't set stock count of a deleted book")
	}
}

//Set isDeleted of the book
func (b *Book) SetIsDeleted(isDeleted bool) {
	if !b.isDeleted {
		b.isDeleted = isDeleted
	} else {
		fmt.Println("You can't set isDeleted of a deleted book")
	}
}

//Return all information of the book as a string
func (b *Book) GetBookInformation() string {
	if !b.isDeleted {
		return "ID: " + strconv.Itoa(b.id) + "\n" +
			"Title: " + b.title + "\n" +
			"Stock Code: " + b.stockCode + "\n" +
			"Stock Count: " + strconv.Itoa(b.stockCount) + "\n" +
			"ISBN: " + strconv.Itoa(b.isbn) + "\n" +
			"Page Count: " + strconv.Itoa(b.pageCount) + "\n" +
			"Price: " + strconv.FormatFloat(b.price, 'f', 2, 64) + "\n" +
			"Author: " + b.AuthorName() + "\n" +
			"Is Deleted: " + strconv.FormatBool(b.isDeleted)
	} else {
		fmt.Println("You can't get information of a deleted book")
		return "Deleted book"
	}
}
