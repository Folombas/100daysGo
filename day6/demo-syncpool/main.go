package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Регистрация обработчиков
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testHandler)
	
	// Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}