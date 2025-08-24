package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Инициализируем хранилище пользователей
	InitUserStorage()

	// Настраиваем обработчики
	http.HandleFunc("/", authMiddleware(indexHandler))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/dashboard", authMiddleware(dashboardHandler))

	// Обслуживание статических файлов
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Запускаем сервер
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}