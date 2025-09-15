package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"urok53-peremennye/internal/examples"
)

//go:embed web/templates/* web/static/**
var embeddedFS embed.FS

var (
	tmpl *template.Template
)

func mustLoadTemplates() *template.Template {
	// Parse all templates from embedded filesystem
	t, err := template.New("").Funcs(template.FuncMap{
		"join": func(items []string, sep string) string {
			res := ""
			for i, s := range items {
				if i > 0 {
					res += sep
				}
				res += s
			}
			return res
		},
	}).ParseFS(embeddedFS, "web/templates/*.tmpl")
	if err != nil {
		log.Fatalf("ошибка загрузки шаблонов: %v", err)
	}
	return t
}

func main() {
	tmpl = mustLoadTemplates()

	mux := http.NewServeMux()

	// Static files
	fs := http.FS(embeddedFS)
	staticPrefix := "/static/"
	mux.Handle(staticPrefix, http.StripPrefix(staticPrefix, http.FileServer(fs)))

	// Routes
	mux.HandleFunc("/", handleIndex)

	addr := ":8080"
	log.Printf("✅ Урок \"Властелин Переменных\" запущен: http://localhost%v", addr)
	if err := http.ListenAndServe(addr, withCommonHeaders(mux)); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func withCommonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	data := struct {
		Title    string
		Subtitle string
		Examples []examples.Example
	}{
		Title:    "Властелин Переменных",
		Subtitle: "Интерактивный мини-курс по переменным в Go",
		Examples: examples.GetExamples(),
	}

	if err := tmpl.ExecuteTemplate(w, "index.tmpl", data); err != nil {
		// try to show a simple error page without leaking sensitive info
		log.Printf("template error: %v", err)
		http.Error(w, "Ошибка рендеринга страницы", http.StatusInternalServerError)
	}
}
