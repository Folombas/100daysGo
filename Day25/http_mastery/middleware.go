package main

import (
	"fmt"
	"net/http"
	"time"
)

// Middleware для логирования
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Создаем кастомный ResponseWriter для перехвата статуса
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		
		// Вызываем следующий обработчик
		next(wrapped, r)
		
		// Логируем запрос
		duration := time.Since(start)
		fmt.Printf("[%s] %s %s - %d - %v\n", 
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method, 
			r.URL.Path, 
			wrapped.statusCode,
			duration,
		)
	}
}

// Кастомный ResponseWriter для перехвата статуса
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Middleware для CORS
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next(w, r)
	}
}