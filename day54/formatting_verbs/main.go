package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// getCategory определяет категорию глагола форматирования
func getCategory(verb string) string {
	switch {
	case strings.ContainsAny(verb, "sq"):
		return "string"
	case strings.ContainsAny(verb, "dfebx+-."):
		return "number"
	case strings.ContainsAny(verb, "vT"):
		return "struct"
	default:
		return "special"
	}
}

func main() {
	// Настройка обработчиков
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Сервер запущен на http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Установка UTF-8 для корректного отображения кириллицы
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Создаем карту функций для шаблона
	funcMap := template.FuncMap{
		"getCategory": getCategory,
	}

	// Парсинг шаблона с функциями
	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Получение примеров форматирования
	examples := GetFormattingExamples()

	// Выполнение шаблона
	data := map[string]interface{}{
		"Examples": examples,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка выполнения шаблона: "+err.Error(), http.StatusInternalServerError)
	}
}
