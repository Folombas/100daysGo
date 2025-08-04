package main

import (
	"fmt"
	"log"
	"net/http"

	//"path/filepath"
	"time"
)

func main() {
	// Создаем маршрутизатор
	mux := http.NewServeMux()

	// Обслуживаем статические файлы
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", http.StripPrefix("/", fs))

	// Специальный маршрут для гофера
	mux.HandleFunc("GET /gopher", gopherHandler)

	// Middleware для логирования
	handler := loggingMiddleware(mux)

	// Настройка сервера
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Запуск сервера
	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	fmt.Println("👉 Статика доступна по: http://localhost:8080/")
	fmt.Println("👉 Страница с гофером: http://localhost:8080/gopher")

	log.Fatal(server.ListenAndServe())
}

// Обработчик страницы с гофером
func gopherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
	<title>Гофер!</title>
	<link rel="stylesheet" href="/css/style.css">
</head>
<body>
	<div class="container">
		<h1>Привет, это Гофер!</h1>
		<div class="gopher-container">
			<pre class="gopher">
				<img src="/images/gopher.png" alt="gopher">
			</pre>
		</div>
		<p>Символ языка Go - милый гофер!</p>
		<a href="/" class="btn">На главную</a>
	</div>
</body>
</html>
	`)
}

// Middleware для логирования
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}
