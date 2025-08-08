package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

// Пользовательская структура с методами
type User struct {
	FirstName string
	LastName  string
	Birthday  time.Time
}

// Метод структуры для вычисления возраста
func (u User) Age() int {
	years := time.Since(u.Birthday).Hours() / 24 / 365
	return int(years)
}

// Метод для форматированного имени
func (u User) FullName() string {
	return fmt.Sprintf("%s %s", u.LastName, u.FirstName)
}

// Функции для регистрации в шаблонах
var funcMap = template.FuncMap{
	"upper":    strings.ToUpper,
	"rusMonth": russianMonth,
}

// Вспомогательная функция для месяца на русском
func russianMonth(t time.Time) string {
	months := []string{
		"января", "февраля", "марта", "апреля",
		"мая", "июня", "июля", "августа",
		"сентября", "октября", "ноября", "декабря",
	}
	return months[t.Month()-1]
}

func main() {
	// Создаем демо-пользователя
	birthDate := time.Date(1987, time.November, 30, 0, 0, 0, 0, time.UTC)
	user := User{"Гоша", "Golang Programmer", birthDate}

	// Обработчик главной страницы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Парсинг шаблона с функциями
		tmpl := template.Must(template.New("index.html").
			Funcs(funcMap).
			ParseFiles("templates/index.html"))

		// Выполняем шаблон с данными
		err := tmpl.Execute(w, map[string]interface{}{
			"User":  user,
			"Title": "День 14 - Методы и Функции",
		})

		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка шаблона: %s", err), http.StatusInternalServerError)
		}
	})

	// Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
