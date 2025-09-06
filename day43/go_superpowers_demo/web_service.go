package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// User представляет структуру пользователя
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Обработчик для /api/users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Алексей Петров", Email: "alex@example.com", CreatedAt: time.Now()},
		{ID: 2, Name: "Мария Сидорова", Email: "maria@example.com", CreatedAt: time.Now()},
		{ID: 3, Name: "Иван Иванов", Email: "ivan@example.com", CreatedAt: time.Now()},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// Обработчик для /api/health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
		"runtime":   "Go",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// StartWebServer запускает веб-сервер
func StartWebServer() {
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/health", healthHandler)

	port := ":8080"
	fmt.Printf("🚀 Веб-сервер запущен на http://localhost%s\n", port)
	fmt.Println("📊 Доступные эндпоинты:")
	fmt.Printf("   GET %s/api/users\n", port)
	fmt.Printf("   GET %s/api/health\n", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}