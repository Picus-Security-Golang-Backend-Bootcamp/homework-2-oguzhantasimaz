package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

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
				fmt.Printf("\nAranan kitap: %s ", strings.Join(args[1:], " "))
				book := library.Search(strings.Join(args[1:], " "))
				if book != nil {
					fmt.Printf("\nBulunan kitap: %s\n", book.Title())
				} else {
					fmt.Println("\nBu kitap bulunamadı")
				}
			}
		} else if args[0] == "list" {
			library.List()
		}
	}
}
