package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-1-oguzhantasimaz/utilites"
)

// Library struct
type Library struct {
	Books []*Book
}

func NewLibrary(path string) (*Library, error) {
	//read books.txt from assets folder and create a new library
	books := []*Book{}
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	//read file line by line
	//and create a new book for each line
	//and add it to the books slice
	counter := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		books = append(books, &Book{
			id:         counter,
			title:      scanner.Text(),
			stockCode:  utilites.RandomString(5),
			stockCount: int(utilites.RandomInt(100)),
			isbn:       int(utilites.RandomInt(9999)),
			pageCount:  int(utilites.RandomInt(1000)),
			price:      float64(utilites.RandomInt(1000)),
			author:     &Author{name: utilites.RandomString(10), surname: utilites.RandomString(10)},
			isDeleted:  false,
		})
		counter++
	}

	//close the file after function ends
	defer file.Close()

	return &Library{
		Books: books,
	}, nil
}

//list all books
func (l *Library) List() {
	for _, book := range l.Books {
		fmt.Println(*book)
	}
}

//buy book from the library by id and price parameter
func (l *Library) Buy(id int, count int) {
	//search for a book in the library
	for _, book := range l.Books {
		if book.ID() == id {
			//Check if book is deleted
			if !book.IsDeleted() {
				//check if stock count is greater than 0
				if book.StockCount() > 0 {
					//if count is lesser than stock count
					//then set stock count to minus the count
					if count < book.StockCount() {
						book.SetStockCount(book.StockCount() - count)
						fmt.Println("You have bought", count, "copies of", book.Title())
						return
					} else {
						//if count is greater than stock count
						//then message that you have bought all the books
						book.SetStockCount(0)
						fmt.Println("You have bought all the books of", book.Title())
						return
					}
				} else {
					fmt.Println("Sorry, we don't have enough books in stock")
					return
				}
			} else {
				fmt.Println("Book is deleted")
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
		if book.ID() == id {
			//check if book is deleted
			if !book.IsDeleted() {
				//if book is found
				//then set isDeleted to true
				book.SetIsDeleted(true)
				fmt.Println("Book", book.Title(), "is deleted")
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
func (l *Library) SearchBy(value interface{}) *Book {
	//search for a book in the library
	switch v := value.(type) {
	case string:
		for _, book := range l.Books {
			if strings.Contains(strings.ToLower(book.StockCode()), strings.ToLower(value.(string))) || strings.Contains(strings.ToLower(book.Title()), strings.ToLower(value.(string))) || strings.Contains(strings.ToLower(book.AuthorName()), strings.ToLower(value.(string))) {
				if book.isDeleted == false {
					return book
				} else {
					fmt.Println("You found a deleted book")
				}
			}
		}
	case int:
		for _, book := range l.Books {
			if book.ISBN() == value.(int) {
				if book.isDeleted == false {
					return book
				} else {
					fmt.Println("You found a deleted book")
				}
			}
		}
	default:
		fmt.Printf("\nInvalid search type %T", v)
		return nil
	}
	return nil
}
