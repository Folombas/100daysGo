package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

// Структуры данных
type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type PageData struct {
	GetResult  string
	PostResult string
}

func main() {
	// Статические файлы
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Обработчики
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{Timeout: 10 * time.Second}
	
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		http.Error(w, "Ошибка GET-запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var post Post
	json.Unmarshal(body, &post)

	// Форматируем результат
	result := fmt.Sprintf(`
		<div class="result-card success">
			<h3>📨 Получен пост #%d</h3>
			<p><strong>Заголовок:</strong> %s</p>
			<p><strong>Текст:</strong> %s</p>
			<p><strong>UserID:</strong> %d</p>
			<p><strong>Статус:</strong> %s</p>
		</div>
	`, post.ID, post.Title, post.Body, post.UserID, resp.Status)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(result))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{Timeout: 10 * time.Second}
	
	newPost := Post{
		Title:  "Мой пост из Day12",
		Body:   "Создано в рамках марафона '100 дней программирования на Go'!",
		UserID: 1,
	}

	jsonData, _ := json.Marshal(newPost)
	resp, err := client.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		http.Error(w, "Ошибка POST-запроса: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var createdPost Post
	json.Unmarshal(body, &createdPost)

	// Форматируем результат
	result := fmt.Sprintf(`
		<div class="result-card success">
			<h3>✅ Успешно создан новый пост!</h3>
			<p><strong>ID:</strong> %d</p>
			<p><strong>Заголовок:</strong> %s</p>
			<p><strong>Текст:</strong> %s</p>
			<p><strong>Статус:</strong> %s</p>
		</div>
	`, createdPost.ID, createdPost.Title, createdPost.Body, resp.Status)

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(result))
}