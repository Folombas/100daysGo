package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Создаем кастомный HTTP-сервер
	server := &http.Server{
		Addr:         ":8080",
		Handler:      setupRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	fmt.Println("📚 Доступные эндпоинты:")
	fmt.Println("   GET  /api/hello - Базовый обработчик")
	fmt.Println("   GET  /api/time - Текущее время")
	fmt.Println("   POST /api/echo - Эхо-ответ")
	fmt.Println("   GET  /api/external - Внешний API")
	
	log.Fatal(server.ListenAndServe())
}