package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Middleware для установки правильного Content-Type
func addContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Для HTML-страниц
		if r.URL.Path == "/" || r.URL.Path == "/about" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		}
		// Для JSON API
		if r.URL.Path == "/user/" || r.URL.Path == "/status" {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
		}
		next.ServeHTTP(w, r)
	})
}

// Обработчик главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>100 дней Go - Главная</title>
		<style>
			body {
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				max-width: 800px;
				margin: 0 auto;
				padding: 20px;
				background-color: #f5f7fa;
				color: #333;
			}
			header {
				background-color: #4b86b4;
				color: white;
				padding: 20px;
				border-radius: 10px;
				text-align: center;
				margin-bottom: 30px;
			}
			nav {
				display: flex;
				justify-content: center;
				gap: 15px;
				margin-bottom: 20px;
			}
			nav a {
				color: white;
				text-decoration: none;
				font-weight: bold;
				padding: 8px 15px;
				border-radius: 5px;
				background-color: #2a4d69;
				transition: background-color 0.3s;
			}
			nav a:hover {
				background-color: #1a3c5a;
			}
			.progress-bar {
				height: 20px;
				background-color: #e0e0e0;
				border-radius: 10px;
				margin: 20px 0;
				overflow: hidden;
			}
			.progress {
				height: 100%;
				width: 8%%;
				background-color: #4CAF50;
			}
		</style>
	</head>
	<body>
		<header>
			<h1>Челлендж "100 дней программирования на Go"</h1>
			<nav>
				<a href="/">Главная</a>
				<a href="/about">О проекте</a>
				<a href="/user/1">Мой прогресс</a>
			</nav>
		</header>
		
		<main>
			<h2>День 8: Основы HTTP</h2>
			<p>Сегодня мы изучаем создание HTTP-серверов на Go!</p>
			
			<h3>Наш прогресс:</h3>
			<div class="progress-bar">
				<div class="progress"></div>
			</div>
			<p>8 из 100 дней успешно завершено</p>
			
			<h3>Что мы уже изучили:</h3>
			<ul>
				<li>Синтаксис Go</li>
				<li>Функции и структуры</li>
				<li>Работу с файлами</li>
				<li>Тестирование кода</li>
				<li>Обработку XML</li>
			</ul>
		</main>
	</body>
	</html>
	`)
}

// Обработчик страницы "О проекте"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>О проекте</title>
		<style>
			body {
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
				max-width: 800px;
				margin: 0 auto;
				padding: 20px;
				background-color: #f5f7fa;
				color: #333;
			}
			.header {
				background-color: #4b86b4;
				color: white;
				padding: 20px;
				border-radius: 10px;
				text-align: center;
				margin-bottom: 30px;
			}
			.goal-card {
				background-color: white;
				border-radius: 10px;
				padding: 20px;
				margin-bottom: 20px;
				box-shadow: 0 2px 10px rgba(0,0,0,0.1);
			}
			.home-link {
				display: inline-block;
				margin-top: 20px;
				padding: 10px 20px;
				background-color: #2a4d69;
				color: white;
				text-decoration: none;
				border-radius: 5px;
			}
		</style>
	</head>
	<body>
		<div class="header">
			<h1>О нашем проекте</h1>
		</div>
		
		<div class="goal-card">
			<h2>Челлендж "100 дней программирования на Go"</h2>
			<p>Мы участвуем в интенсивном марафоне по изучению языка программирования Go (Golang). Каждый день мы осваиваем новую тему, решаем практические задачи и совершенствуем свои навыки.</p>
		</div>
		
		<div class="goal-card">
			<h3>Наши цели:</h3>
			<ul>
				<li>Освоить фундаментальные концепции Go</li>
				<li>Научиться создавать эффективные и надежные приложения</li>
				<li>Изучить работу с сетью, базами данных и параллельным программированием</li>
				<li>Создать портфолио проектов</li>
				<li>Подготовиться к работе с Go в реальных проектах</li>
			</ul>
		</div>
		
		<div class="goal-card">
			<h3>Почему именно Go?</h3>
			<ul>
				<li>Высокая производительность</li>
				<li>Простота и читаемость кода</li>
				<li>Мощная стандартная библиотека</li>
				<li>Встроенная поддержка многозадачности</li>
				<li>Растущая популярность в облачных технологиях</li>
			</ul>
		</div>
		
		<a href="/" class="home-link">Вернуться на главную</a>
	</body>
	</html>
	`)
}

// Обработчик пользовательского профиля
func userHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"user_id": id,
		"name":    "Иван Иванов",
		"email":   "user" + id + "@example.com",
		"progress": map[string]interface{}{
			"days_completed": 8,
			"last_date":      time.Now().Format("2006-01-02"),
			"next_topic":     "Базы данных в Go",
		},
	})
}

// Обработчик статуса сервера
func statusHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "ok",
		"version": "1.0.0",
		"uptime":  "8 дней",
		"stats": map[string]interface{}{
			"requests": 1248,
			"users":    5,
			"errors":   0,
		},
	})
}