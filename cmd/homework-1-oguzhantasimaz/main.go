package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-1-oguzhantasimaz/internal/models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-1-oguzhantasimaz/utilites"
)

var books []*models.Book

func init() {
	// create books

	bookTitleList := [5]string{"Go Programming", "Python Programming", "Java Programming", "C Programming", "C++ Programming"}
	authorNameList := [5]string{"Harry", "Hermonie", "Jhon", "Peter", "Paul"}
	authorSurnameList := [5]string{"Potter", "Granger", "Doe", "Smith", "Jones"}

	for i := 0; i < 5; i++ {
		book := models.NewBook(
			i,
			bookTitleList[i],
			utilites.RandomString(5),
			utilites.RandomInt(100),
			utilites.RandomInt(9999),
			utilites.RandomInt(1000),
			float64(utilites.RandomInt(1000)),
			false,
			models.NewAuthor(authorNameList[i], authorSurnameList[i]),
		)
		books = append(books, book)
	}

}

func main() {
	args := os.Args
	library, err := models.NewLibrary(books)

	if err != nil {
		log.Fatal(err)
	}

	if len(args) == 1 {
		projectName := path.Base(args[0])
		fmt.Printf("\n %s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n buy => satin alma islemi icin\n delete => silme islemi icin", projectName)
		os.Exit(1)
	} else {
		// read args if search or list
		args = args[1:]
		if args[0] == "search" {
			//search by any of stock code, isbn, title, author name
			book := library.SearchBy(args[1])
			//search by stock code
			if book != nil {
				fmt.Printf("\nBulunan kitap: %v\n", book.GetBookInformation())
			} else {
				fmt.Println("\nBu kitap bulunamadı")
			}

		} else if args[0] == "list" {
			library.List()
		} else if args[0] == "buy" {
			if len(args) >= 2 {
				//get id and count args as integer and buy book by id and count
				id, err := strconv.Atoi(args[1])
				if err != nil {
					log.Fatal(err)
				}
				count, err := strconv.Atoi(args[2])
				if err != nil {
					log.Fatal(err)
				}
				library.Buy(id, count)
			}
		} else if args[0] == "delete" {
			//delete book by id
			if len(args) >= 2 {
				id, err := strconv.Atoi(args[1])
				if err != nil {
					log.Fatal(err)
				}
				library.Delete(id)
			}
		} else if args[0] == "exit" {
			os.Exit(1)
		}
	}
}
