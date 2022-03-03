package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-oguzhantasimaz/internal/models"
	utilites "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-oguzhantasimaz/utilities"
)

var books []*models.Book

func init() {
	// create information for books
	bookTitleList := [5]string{"Go Programming", "Python Programming", "Java Programming", "C Programming", "C++ Programming"}
	authorNameList := [5]string{"Harry", "Hermonie", "Jhon", "Peter", "Paul"}
	authorSurnameList := [5]string{"Potter", "Granger", "Doe", "Smith", "Jones"}
	isDeletedList := [5]bool{true, false, false, false, false}

	//create book list
	for i := 0; i < 5; i++ {
		book := models.NewBook(
			i+1,
			bookTitleList[i],
			utilites.RandomString(5),
			utilites.RandomInt(100),
			utilites.RandomInt(9999),
			utilites.RandomInt(1000),
			float64(utilites.RandomInt(1000)),
			isDeletedList[i],
			models.NewAuthor(authorNameList[i], authorSurnameList[i]),
		)
		books = append(books, book)
	}

}

func main() {
	// read args from command line
	args := os.Args
	//create new library from book list created above
	library, err := models.NewLibrary(books)

	if err != nil {
		log.Fatal(err)
	}

	if len(args) == 1 {
		projectName := path.Base(args[0])
		fmt.Printf("\n %s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n buy => satin alma islemi icin\n delete => silme islemi icin", projectName)
		os.Exit(1)
	} else {
		args = args[1:]
		if args[0] == "search" {
			var book *models.Book

			//check argument to see if it is an int or string
			if utilites.IsInt(args[1]) {
				book = library.SearchByInt(utilites.StringToInt(args[1]))
			} else {
				book = library.SearchByString(args[1])
			}
			//if book is not found
			if book != nil {
				fmt.Printf("\nBook found: %v\n", book.GetBookInformation())
			} else {
				fmt.Println("\nBook not found")
			}
		} else if args[0] == "list" {
			library.List()
		} else if args[0] == "buy" {
			if len(args) >= 2 {
				//get id and count args as integer and buy book by id and count
				id, err := strconv.Atoi(args[1])
				if err != nil {
					//create error if id is not integer
					error := fmt.Errorf("%s is not integer", args[1])
					log.Fatal(error)
				}
				count, err := strconv.Atoi(args[2])
				if err != nil {
					//create error if count is not integer
					error := fmt.Errorf("%s is not integer", args[2])
					log.Fatal(error)
				}
				library.Buy(id, count)
			}
		} else if args[0] == "delete" {
			//delete book by id
			if len(args) >= 2 {
				//convert id to integer
				id, err := strconv.Atoi(args[1])
				if err != nil {
					//create error if id is not integer
					error := fmt.Errorf("%s is not integer", args[1])
					log.Fatal(error)
				}
				library.Delete(id)
			}
		}
	}
}
