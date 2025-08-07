package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// User данные для передачи в шаблон
type User struct {
	Name  string
	Email string
	Posts int
}

func main() {
	// Регистрируем обработчики
	http.HandleFunc("/inline", inlineHandler)
	http.HandleFunc("/from-file", fileHandler)
	http.HandleFunc("/", homeHandler)

	// Запускаем сервер
	fmt.Println("Сервер запущен: http://localhost:8080")
	fmt.Println("Доступные пути: /inline, /from-file")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Ошибка сервера: %v", err)
		os.Exit(1)
	}
}

// Обработчик главной страницы
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Демо работы с шаблонами</h1>
		<ul>
			<li><a href="/inline">Inline-шаблон</a></li>
			<li><a href="/from-file">Шаблон из файла</a></li>
		</ul>`)
}

// Обработчик встроенного шаблона
func inlineHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Создаем шаблон прямо в коде
	tmpl := `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Inline-шаблон</title>
</head>
<body>
	<h1>Привет, {{.Name}}!</h1>
	<p>Текущее время: {{.Time}}</p>
	<ul>
		{{range .Items}}
		<li>{{.}}</li>
		{{end}}
	</ul>
</body>
</html>`

	// 2. Парсим шаблон
	t, err := template.New("inline").Parse(tmpl)
	if err != nil {
		http.Error(w, "Ошибка создания шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Подготавливаем данные
	data := struct {
		Name  string
		Time  string
		Items []string
	}{
		Name:  "Гость",
		Time:  time.Now().Format("15:04:05"),
		Items: []string{"Go", "Шаблоны", "Встроенный код"},
	}

	// 4. Устанавливаем кодировку и выполняем шаблон
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Ошибка выполнения шаблона: "+err.Error(), http.StatusInternalServerError)
	}
}

// Обработчик шаблона из файла
func fileHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Получаем абсолютный путь к шаблону
	tmplPath, err := filepath.Abs("templates/layout.html")
	if err != nil {
		http.Error(w, "Ошибка пути: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Парсим несколько файлов шаблонов
	t, err := template.ParseFiles(tmplPath, "templates/user.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Подготавливаем данные
	user := User{
		Name:  "Гоша",
		Email: "gosha_gofer@example.com",
		Posts: 42,
	}

	// 4. Выполняем шаблон с данными
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := t.ExecuteTemplate(w, "layout.html", user); err != nil {
		http.Error(w, "Ошибка выполнения: "+err.Error(), http.StatusInternalServerError)
	}
}