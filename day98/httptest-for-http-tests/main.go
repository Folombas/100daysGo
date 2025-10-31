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

	// Регистрируем обработчики
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
	fmt.Println("📚 Доступные endpoint:")
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
