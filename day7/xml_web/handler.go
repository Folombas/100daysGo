package main

import (
	"encoding/xml"
	"html/template"
	"net/http"
	"os"
)

type Library struct {
	XMLName xml.Name `xml:"library"`
	Books   []Book   `xml:"book"`
}

type Book struct {
	ID          string  `xml:"id,attr"`
	Title       string  `xml:"title"`
	Author      string  `xml:"author"`
	Year        int     `xml:"year"`
	Price       float64 `xml:"price"`
	Description string  `xml:"description"`
}

func parseXML(filename string) (Library, error) {
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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	library, err := parseXML("books.xml")
	if err != nil {
		http.Error(w, "Error parsing XML: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, library.Books)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}