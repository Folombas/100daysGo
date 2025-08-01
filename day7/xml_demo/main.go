package main

import (
	"fmt"
	"log"
)

func main() {
	// Чтение и парсинг XML
	library, err := ParseXML("books.xml")
	if err != nil {
		log.Fatalf("XML parsing error: %v", err)
	}
	PrintBooks(library)

	// Генерация нового XML
	if err := GenerateXML(); err != nil {
		log.Fatalf("XML generation error: %v", err)
	}
	fmt.Println("XML file 'generated.xml' created successfully")
}
