// web/server.go
package web

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"fmt"

	"go_ai/ai_simulator"
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

	// Генерируем "ИИ-ответ" на фиксированный промпт
	prompt := "расскажи о человеке, который учит Go, чтобы купить новую видеокарту"
	aiResponse := ai_simulator.GenerateText(prompt)

	// Собираем системную информацию
	sysInfo := map[string]string{
		"OS":       runtime.GOOS,
		"Arch":     runtime.GOARCH,
		"CPUs":     fmt.Sprintf("%d", runtime.NumCPU()),
		"GoVer":    runtime.Version(),
		"MemHint":  "Go эффективен даже на 16 ГБ ОЗУ!",
		"GPUHint":  "Go не требует GPU — учись где угодно!",
	}

	data := struct {
		Title     string
		AiText    string
		SysInfo   map[string]string
	}{
		Title:   "Go & ИИ — Твой путь к новому железу",
		AiText:  aiResponse,
		SysInfo: sysInfo,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка рендеринга", http.StatusInternalServerError)
		log.Println("❌ Ошибка рендеринга:", err)
	}
}
