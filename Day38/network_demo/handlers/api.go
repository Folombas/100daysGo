package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Иван Иванов", Email: "ivan@example.com", CreatedAt: "2024-01-15"},
		{ID: 2, Name: "Мария Петрова", Email: "maria@example.com", CreatedAt: "2024-01-16"},
		{ID: 3, Name: "Алексей Сидоров", Email: "alex@example.com", CreatedAt: "2024-01-17"},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"users":    users,
		"count":    len(users),
		"page":     1,
		"per_page": 10,
	})
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	stats := map[string]interface{}{
		"total_requests":   1423,
		"active_connections": 15,
		"requests_per_second": 2.34,
		"uptime_seconds":   time.Since(startTime).Seconds(),
		"memory_usage_mb":  getMemoryUsage(),
		"goroutines":       8,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(stats)
}