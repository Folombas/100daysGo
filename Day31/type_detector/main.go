package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Настраиваем обработчики
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/detect", detectHandler)
	
	// Обслуживание статических файлов
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	// Запускаем сервер
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}