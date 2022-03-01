package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-1-oguzhantasimaz/internal/models"
)

func main() {
	args := os.Args
	// dirPath, err := os.Getwd()
	//create library from books.txt
	library, err := models.NewLibrary("../../assets/books.txt")
	if err != nil {
		log.Fatal(err)
	}
	// read args if search or list

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
