package main

import (
	"encoding/xml"
	"os"
)

type Library struct {
	XMLName xml.Name `xml:"library"`
	Books   []Book   `xml:"book"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      string   `xml:"id,attr"`
	Title   string   `xml:"title"`
	Author  string   `xml:"author"`
	Year    int      `xml:"year"`
	Price   float64  `xml:"price"`
}

func GenerateXML() error {
	library := Library{
		Books: []Book{
			{
				ID:     "201",
				Title:  "Clean Code",
				Author: "Robert C. Martin",
				Year:   2008,
				Price:  49.99,
			},
			{
				ID:     "202",
				Title:  "Design Patterns",
				Author: "Erich Gamma",
				Year:   1994,
				Price:  59.99,
			},
		},
	}

	file, err := os.Create("generated.xml")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	return encoder.Encode(library)
}