package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func main() {
	// Создаем роутер
	router := mux.NewRouter()
	
	// Регистрируем обработчики
	router.HandleFunc("/articles", GetArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", GetArticle).Methods("GET")
	router.HandleFunc("/articles", CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", UpdateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", DeleteArticle).Methods("DELETE")
	
	// Старт сервера
	fmt.Println("Сервер запущен: http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}