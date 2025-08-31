package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>Сетевой демо-сервер Go</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 800px; margin: 0 auto; }
        .endpoint { background: #f5f5f5; padding: 20px; margin: 10px 0; border-radius: 5px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Демонстрация сетевых возможностей Go</h1>
        <p>Сервер запущен: %s</p>
        
        <div class="endpoint">
            <h2>📊 API эндпоинты:</h2>
            <ul>
                <li><a href="/api/health">/api/health</a> - Статус сервера</li>
                <li><a href="/api/time">/api/time</a> - Текущее время</li>
                <li><a href="/api/users">/api/users</a> - Данные пользователей (JSON)</li>
            </ul>
        </div>

        <div class="endpoint">
            <h2>🔐 Защищенные роуты:</h2>
            <ul>
                <li><a href="/admin">/admin</a> - Требует аутентификацию</li>
                <li><a href="/admin/dashboard">/admin/dashboard</a> - Панель управления</li>
            </ul>
        </div>

        <div class="endpoint">
            <h2>📡 Сетевые тесты:</h2>
            <ul>
                <li><a href="/network/test">/network/test</a> - Тест сетевого соединения</li>
                <li><a href="/network/stats">/network/stats</a> - Статистика сервера</li>
            </ul>
        </div>
    </div>
</body>
</html>
`, time.Now().Format("2006-01-02 15:04:05"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{
	"status": "healthy",
	"timestamp": "%s",
	"uptime": "%.0f секунд",
	"memory_usage": "%.2f MB"
}`, time.Now().Format(time.RFC3339), time.Since(startTime).Seconds(), getMemoryUsage())
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{
	"current_time": "%s",
	"timezone": "%s",
	"unix_timestamp": %d
}`, time.Now().Format("2006-01-02 15:04:05"), time.Local.String(), time.Now().Unix())
}

func NetworkTestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{
	"network_test": "successful",
	"client_ip": "%s",
	"user_agent": "%s",
	"protocol": "%s"
}`, getClientIP(r), r.UserAgent(), r.Proto)
}

var startTime = time.Now()

func getMemoryUsage() float64 {
	// Заглушка для демонстрации
	return 12.34
}

func getClientIP(r *http.Request) string {
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	return r.RemoteAddr
}