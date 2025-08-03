package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Создаем маршрутизатор
	mux := http.NewServeMux()
	
	// Регистрируем обработчики
	mux.HandleFunc("GET /search", searchHandler)
	mux.HandleFunc("GET /user/{id}", userHandler)
	mux.HandleFunc("POST /register", registerHandler)
	mux.HandleFunc("POST /json", jsonHandler)
	
	// Настраиваем middleware
	handler := addHeadersMiddleware(mux)
	handler = loggingMiddleware(handler)
	
	// Конфигурация сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	// Запуск сервера
	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	fmt.Println("👉 Доступные эндпоинты:")
	fmt.Println("   GET  /search?q=...&category=...")
	fmt.Println("   GET  /user/{id}")
	fmt.Println("   POST /register (form-data)")
	fmt.Println("   POST /json (application/json)")
	
	log.Fatal(server.ListenAndServe())
}

// Middleware для добавления заголовков
func addHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

// Middleware для логирования
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}