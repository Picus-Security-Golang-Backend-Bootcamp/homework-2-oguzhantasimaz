package models

import (
	"fmt"

	utilites "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-oguzhantasimaz/utilities"
)

// Library struct
type Library struct {
	Books []*Book
}

func NewLibrary(bookList []*Book) (*Library, error) {
	//create a new library
	return &Library{
		Books: bookList,
	}, nil
}

//List all books
func (l *Library) List() {
	for _, book := range l.Books {
		fmt.Print("\n" + book.GetBookInformation() + "\n")
	}
}

//buy book from the library by id and price parameter
func (l *Library) Buy(id int, count int) {
	//search for a book in the library
	for _, book := range l.Books {
		if book.id == id {
			//Check if book is deleted
			if !book.isDeleted {
				//check if stock count is greater than 0
				if book.stockCount > 0 {
					//if count is less than 1 then return error
					if count > 0 {

						//if count is lesser than stock count
						//then set stock count to minus the count
						if count < book.stockCount {
							book.SetStockCount(book.stockCount - count)
							fmt.Println("You have bought", count, "copies of", book.title)
							return
						} else {
							//if count is greater than stock count
							//then message that you have bought all the books
							book.SetStockCount(0)
							fmt.Println("You have bought all the books of", book.title)
							return
						}
					} else {
						fmt.Println("Count should be greater than 0")
						return
					}
				} else {
					fmt.Println("Sorry, we don't have enough books in stock")
					return
				}
			} else {
				fmt.Println("Sorry book is deleted, you can't buy it")
				return
			}
		}
	}
	fmt.Println("Book not found")
	return //if book is not found
}

//delete book from the library by id
func (l *Library) Delete(id int) {
	//search for a book in the library
	for _, book := range l.Books {
		if book.id == id {
			//check if book is deleted
			if !book.isDeleted {
				//if book is found
				//then set isDeleted to true
				book.SetIsDeleted(true)
				fmt.Println("Book", book.title, "is deleted")
				return
			} else {
				fmt.Println("Book is already deleted")
				return
			}
		}
	}
	fmt.Println("Book not found")
	return //if book is not found
}

//Search book by stockCode, isbn, title
func (l *Library) SearchByString(value string) *Book {
	//search for a book in the library
	for _, book := range l.Books {
		if utilites.StrContains(book.stockCode, value) || utilites.StrContains(book.title, value) || utilites.StrContains(book.AuthorName(), value) {
			if book.isDeleted == false {
				return book
			} else {
				fmt.Println("You found a deleted book")
			}
		}
	}
	return nil
}

//Search book by integer id, isbn
func (l *Library) SearchByInt(value int) *Book {
	for _, book := range l.Books {
		if book.isbn == value || book.id == value {
			if book.isDeleted == false {
				return book
			} else {
				fmt.Println("You found a deleted book")
			}
		}
	}
	return nil
}
