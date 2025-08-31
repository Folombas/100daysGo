package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error": "Требуется аутентификация", "code": 401}`)
			return
		}

		// Простая проверка токена (в реальном приложении нужно использовать JWT или другую систему)
		if authToken != "Bearer secret-token-123" {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, `{"error": "Неверный токен доступа", "code": 403}`)
			return
		}

		next(w, r)
	}
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
	}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{
		"message": "Добро пожаловать в панель администратора",
		"features": ["статистика", "пользователи", "настройки"],
		"access_level": "admin"
	}`)
}

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
	<title>Панель управления</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 40px; }
		.dashboard { background: #e8f5e9; padding: 20px; border-radius: 10px; }
	</style>
</head>
<body>
	<div class="dashboard">
		<h1>📊 Панель управления администратора</h1>
		<p>Статус системы: <strong>активна</strong></p>
		<p>Пользователей онлайн: <strong>15</strong></p>
		<p>Запросов сегодня: <strong>1,234</strong></p>
	</div>
</body>
</html>
`)
}