package models

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		books = append(books, &Book{
			title: scanner.Text(),
		})
	}

	//close the file after function ends
	defer file.Close()

	return &Library{
		Books: books,
	}, nil
}

func (l *Library) Search(title string) *Book {
	//search for a book in the library
	for _, book := range l.Books {
		if strings.Contains(strings.ToLower(book.Title()), strings.ToLower(title)) {
			return book
		}
	}
	return nil
}

//list all books
func (l *Library) List() {
	for _, book := range l.Books {
		fmt.Println(book.Title())
	}
}
