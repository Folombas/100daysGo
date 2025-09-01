package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ServerStats struct {
	StartTime      time.Time `json:"start_time"`
	TotalRequests  int64     `json:"total_requests"`
	ActiveConnections int64 `json:"active_connections"`
	MemoryUsageMB  float64   `json:"memory_usage_mb"`
}

var stats = ServerStats{
	StartTime: time.Now(),
}

func startServer() {
	port := getEnv("PORT", "8080")
	
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/network", networkHandler)

	log.Printf("🚀 Сервер запущен на порту %s", port)
	log.Printf("📊 Режим CI/CD: %s", getEnv("CI", "false"))
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	stats.TotalRequests++
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
	<title>CI/CD Network Demo</title>
	<style>body{font-family: Arial, sans-serif; margin: 40px;}</style>
</head>
<body>
	<h1>🚀 CI/CD Network Demo</h1>
	<p>Сервер работает: %s</p>
	<p>Всего запросов: %d</p>
	<p><a href="/health">Health Check</a></p>
	<p><a href="/stats">Статистика</a></p>
	<p><a href="/api/users">API Users</a></p>
</body>
</html>
`, time.Since(stats.StartTime).String(), stats.TotalRequests)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	stats.TotalRequests++
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"version":   "1.0.0",
		"ci_mode":   getEnv("CI", "false"),
	})
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats.TotalRequests++
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(stats)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	stats.TotalRequests++
	
	users := []map[string]interface{}{
		{"id": 1, "name": "Иван Иванов", "email": "ivan@example.com"},
		{"id": 2, "name": "Мария Петрова", "email": "maria@example.com"},
		{"id": 3, "name": "Алексей Сидоров", "email": "alex@example.com"},
	}
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(users)
}

func networkHandler(w http.ResponseWriter, r *http.Request) {
	stats.TotalRequests++
	
	networkInfo := map[string]interface{}{
		"client_ip":    r.RemoteAddr,
		"user_agent":   r.UserAgent(),
		"method":       r.Method,
		"content_type": r.Header.Get("Content-Type"),
		"ci_mode":      getEnv("CI", "false"),
	}
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(networkInfo)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}