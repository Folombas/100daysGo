package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Статические файлы
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Роуты
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/convert", convertHandler)

	fmt.Println("Сервер запущен: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}