package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	mux := http.NewServeMux()
	
	// Регистрация обработчиков
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /about", aboutHandler)
	mux.HandleFunc("GET /user/{id}", userHandler)
	mux.HandleFunc("GET /status", statusHandler)
	
	// Статические файлы (без конфликтов)
	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static"+r.URL.Path)
	})
	
	// Настройка middleware
	handler := addContentTypeMiddleware(mux)
	handler = loggingMiddleware(handler)
	handler = recoveryMiddleware(handler)
	
	// Конфигурация сервера
	server := &http.Server{
		Addr:         HOST + ":" + PORT,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	// Запуск сервера
	fmt.Printf("🚀 HTTP сервер запущен на http://%s:%s\n", HOST, PORT)
	fmt.Println("👉 Попробуйте:")
	fmt.Println("   - http://localhost:8080")
	fmt.Println("   - http://localhost:8080/about")
	fmt.Println("   - http://localhost:8080/user/123")
	fmt.Println("   - http://localhost:8080/static/index.html")
	
	log.Fatal(server.ListenAndServe())
}