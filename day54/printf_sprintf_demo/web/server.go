// web/server.go
package web

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func StartServer() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

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

	data := struct {
		Title string
		Examples []Example
	}{
		Title: "Printf & Sprintf в Go — Демо Day 54",
		Examples: []Example{
			{Code: `fmt.Printf("Привет, %s!", "мир")`, Output: "Привет, мир!"},
			{Code: `fmt.Printf("Число: %d, Плавающее: %.2f", 42, 3.1415)`, Output: "Число: 42, Плавающее: 3.14"},
			{Code: `s := fmt.Sprintf("Сообщение: %s", "Готово")`, Output: `s = "Сообщение: Готово"`},
			{Code: `fmt.Printf("%10s | %-10s", "Право", "Лево")`, Output: "     Право | Лево      "},
			{Code: `fmt.Printf("Цена: %8.2f₽", 1234.5)`, Output: "Цена:  1234.50₽"},
			{Code: `fmt.Printf("Город: %s", "Санкт-Петербург")`, Output: "Город: Санкт-Петербург"},
		},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка рендеринга", http.StatusInternalServerError)
		log.Println("❌ Ошибка рендеринга:", err)
	}
}

type Example struct {
	Code   string
	Output string
}
