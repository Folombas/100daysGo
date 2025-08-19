package main

import (
	"net/http"
	"fmt"
)

// Настройка роутинга
func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	// Статические файлы
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// API endpoints
	mux.HandleFunc("/api/hello", logRequest(helloHandler))
	mux.HandleFunc("/api/time", logRequest(timeHandler))
	mux.HandleFunc("/api/echo", logRequest(echoHandler))
	mux.HandleFunc("/api/external", logRequest(externalAPIHandler))
	
	// Главная страница
	mux.HandleFunc("/", logRequest(homeHandler))
	
	return mux
}

// Обработчик главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<h1>Добро пожаловать в HTTP мастерскую!</h1>
		<p>Используйте эндпоинты:</p>
		<ul>
			<li><a href="/api/hello">/api/hello</a> - Приветствие</li>
			<li><a href="/api/time">/api/time</a> - Текущее время</li>
			<li><a href="/static/index.html">Статические файлы</a></li>
		</ul>
	`)
}