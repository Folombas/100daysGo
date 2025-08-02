package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// Middleware для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

// Middleware для обработки паник
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `
				<!DOCTYPE html>
				<html>
				<head><title>Ошибка сервера</title></head>
				<body>
					<h1>500 - Внутренняя ошибка сервера</h1>
					<p>Попробуйте перезагрузить страницу позже</p>
					<a href="/">Вернуться на главную</a>
				</body>
				</html>
				`)
				log.Printf("Паника: %v\n%s", err, debug.Stack())
			}
		}()
		next.ServeHTTP(w, r)
	})
}