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
		Title   string
		Examples []WidthExample
	}{
		Title: "Форматирование ширины в Go — Демо Day 54x",
		Examples: []WidthExample{
			{
				Description: "Выравнивание по правому краю (ширина 10)",
				Code:        `fmt.Printf("|%10s|", "Go")`,
				Output:      "|        Go|",
			},
			{
				Description: "Выравнивание по левому краю (ширина 10)",
				Code:        `fmt.Printf("|%-10s|", "Go")`,
				Output:      "|Go        |",
			},
			{
				Description: "Заполнение нулями (ширина 6)",
				Code:        `fmt.Printf("ID: %06d", 42)`,
				Output:      "ID: 000042",
			},
			{
				Description: "Динамическая ширина через *",
				Code:        `fmt.Printf("|%*s|", 15, "Динамически")`,
				Output:      "|   Динамически|",
			},
			{
				Description: "Ширина + точность для float",
				Code:        `fmt.Printf("|%10.2f|", 3.1415)`,
				Output:      "|      3.14|",
			},
			{
				Description: "Таблица с кириллицей",
				Code: `fmt.Printf("| %-10s | %6s | %-15s |", "Алексей", "28", "Москва")`,
				Output: "| Алексей    |     28 | Москва          |",
			},
		},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка рендеринга", http.StatusInternalServerError)
		log.Println("❌ Ошибка рендеринга:", err)
	}
}

type WidthExample struct {
	Description string
	Code        string
	Output      string
}
