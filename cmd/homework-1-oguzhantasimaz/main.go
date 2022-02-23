package main

import (
	"fmt"
	"log"
	"os"
	"path"

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
		fmt.Printf("\n %s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n", projectName)
		os.Exit(1)
	} else {
		// read args if search or list
		args = args[1:]
		if args[0] == "search" {
			if len(args) >= 2 {
				book := library.Search(args[1])
				if book != nil {
					fmt.Println(book.Title())
				} else {
					fmt.Println("Bu kitap bulunamadı")
				}
			}
		} else if args[0] == "list" {
			library.List()
		}
	}
}
