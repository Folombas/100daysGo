package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("🌐 HTTPTEST ДЕМОНСТРАЦИЯ")
	fmt.Println("========================")
	fmt.Println()
	fmt.Println("🎯 Гоша, добро пожаловать в мир тестирования HTTP!")
	fmt.Println("💡 Этот модуль покажет, как тестировать API без реальных серверов")
	fmt.Println()

	handler := NewAPIHandler()

	// Обработчик для корневого пути
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>HTTPTest Демонстрация</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .endpoint { background: #f5f5f5; padding: 10px; margin: 10px 0; }
        code { background: #eee; padding: 2px 5px; }
    </style>
</head>
<body>
    <h1>🌐 HTTPTest Демонстрация</h1>
    <p>🎯 Гоша, добро пожаловать в мир тестирования HTTP!</p>

    <h2>📚 Доступные endpoint:</h2>
    <div class="endpoint">
        <strong>GET /user?id=1</strong> - получить пользователя по ID<br>
        <a href="/user?id=1" target="_blank">/user?id=1</a>
    </div>
    <div class="endpoint">
        <strong>GET /users</strong> - получить всех пользователей<br>
        <a href="/users" target="_blank">/users</a>
    </div>
    <div class="endpoint">
        <strong>POST /user</strong> - создать пользователя (используйте Postman или curl)
    </div>

    <h2>🧪 Тестирование:</h2>
    <p>Запустите в терминале:</p>
    <code>go test -v</code> - запустить все тесты<br>
    <code>go test -bench=.</code> - запустить бенчмарки

    <h2>💡 Информация:</h2>
    <p>httptest позволяет тестировать HTTP код без запуска реальных серверов - это быстрее и надежнее!</p>
</body>
</html>
		`)
	})

	// Регистрируем обработчики API
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetUserHandler(w, r)
		case http.MethodPost:
			handler.CreateUserHandler(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/users", handler.GetAllUsersHandler)

	fmt.Println("🚀 Запуск сервера на http://localhost:8080")
	fmt.Println()
	fmt.Println("📚 Теперь доступны:")
	fmt.Println("   GET  /              - HTML страница с информацией")
	fmt.Println("   GET  /user?id=1     - получить пользователя по ID")
	fmt.Println("   POST /user          - создать пользователя")
	fmt.Println("   GET  /users         - получить всех пользователей")
	fmt.Println()
	fmt.Println("🧪 Запуск тестов:")
	fmt.Println("   go test -v          - запустить все тесты")
	fmt.Println("   go test -bench=.    - запустить бенчмарки")
	fmt.Println()
	fmt.Println("💡 Помни: httptest позволяет тестировать HTTP код")
	fmt.Println("   без запуска реальных серверов - это быстрее и надежнее!")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
