package server

import (
	"fmt"
	"net/http"
	"network_demo/handlers"
	"network_demo/middleware"
	"time" // Добавляем импорт time
)

func StartHTTPServer() {
	mux := http.NewServeMux()

	// Статические файлы
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Основные маршруты
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/api/health", handlers.HealthHandler)
	mux.HandleFunc("/api/time", handlers.TimeHandler)
	mux.HandleFunc("/api/users", handlers.UsersHandler)
	mux.HandleFunc("/network/test", handlers.NetworkTestHandler)
	mux.HandleFunc("/network/stats", handlers.StatsHandler)

	// Защищенные маршруты
	mux.HandleFunc("/admin", middleware.AuthMiddleware(middleware.AdminHandler))
	mux.HandleFunc("/admin/dashboard", middleware.AuthMiddleware(middleware.AdminDashboardHandler))

	// Настройка middleware
	wrappedMux := middleware.LoggingMiddleware(mux.ServeHTTP)

	fmt.Println("🌐 HTTP сервер запущен на http://localhost:8080")
	fmt.Println("📊 Доступные эндпоинты:")
	fmt.Println("   /              - Главная страница")
	fmt.Println("   /api/health    - Статус сервера")
	fmt.Println("   /api/time      - Текущее время")
	fmt.Println("   /api/users     - Данные пользователей")
	fmt.Println("   /admin         - Панель администратора (требует токен)")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      wrappedMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}