package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Обработчик главной страницы
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Обработчик входа
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if AuthenticateUser(username, password) {
			// Создаем сессию
			sessionID, err := CreateSession(username)
			if err != nil {
				http.Error(w, "Ошибка создания сессии", http.StatusInternalServerError)
				return
			}

			// Устанавливаем куки
			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    sessionID,
				Expires:  time.Now().Add(24 * time.Hour),
				HttpOnly: true,
				Path:     "/",
			})

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			tmpl := template.Must(template.ParseFiles("templates/login.html"))
			tmpl.Execute(w, map[string]interface{}{
				"Error": "Неверные учетные данные",
			})
		}
	}
}

// Обработчик регистрации
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		if err := CreateUser(username, password, email); err != nil {
			tmpl := template.Must(template.ParseFiles("templates/register.html"))
			tmpl.Execute(w, map[string]interface{}{
				"Error": fmt.Sprintf("Ошибка регистрации: %v", err),
			})
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// Обработчик личного кабинета
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем пользователя из контекста
	user, ok := r.Context().Value("user").(*User)
	if !ok {
		http.Error(w, "Ошибка авторизации", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Username": user.Username,
		"Email":    user.Email,
	})
}

// Обработчик выхода
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Удаляем куки
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	// Удаляем сессию
	cookie, err := r.Cookie("session_id")
	if err == nil {
		DeleteSession(cookie.Value)
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}