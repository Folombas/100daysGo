package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func ParseXML(filename string) (Library, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Library{}, err
	}
	defer file.Close()

	var library Library
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&library)
	return library, err
}

func PrintBooks(library Library) {
	fmt.Println("\nBooks in library:")
	for _, book := range library.Books {
		fmt.Printf(
			"ID: %s\nTitle: %s\nAuthor: %s\nYear: %d\nPrice: $%.2f\n\n",
			book.ID, book.Title, book.Author, book.Year, book.Price,
		)
	}
}