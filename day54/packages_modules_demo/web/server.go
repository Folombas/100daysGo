// web/server.go
package web

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"packages_modules_demo/calculator"
	"packages_modules_demo/formatter"
)

func StartServer() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", homeHandler)

	log.Printf("🚀 Веб-сервер запущен на http://localhost:%s", port)
	return http.ListenAndServe(":"+port, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		log.Println("❌ Ошибка шаблона:", err)
		return
	}

	// Используем импортированные пакеты для генерации данных
	a, b := 100.0, 25.5
	examples := []string{
		formatter.FormatWithLabel("Сложение", "+", a, b, calculator.Add(a, b)),
		formatter.FormatWithLabel("Вычитание", "-", a, b, calculator.Subtract(a, b)),
		formatter.FormatWithLabel("Умножение", "×", a, b, calculator.Multiply(a, b)),
	}

	// Обработка деления
	if result, ok := calculator.Divide(a, b); ok {
		examples = append(examples, formatter.FormatWithLabel("Деление", "÷", a, b, result))
	} else {
		examples = append(examples, "Деление: невозможно (деление на ноль)")
	}

	data := struct {
		Title    string
		Examples []string
	}{
		Title:    "Пакеты, модули и импорты в Go — Демо Day 54",
		Examples: examples,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка рендеринга", http.StatusInternalServerError)
		log.Println("❌ Ошибка рендеринга:", err)
	}
}
