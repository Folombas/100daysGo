package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

// Структура статьи (перенесена из models.go для простоты)
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Хранилище в памяти
var articles = []Article{
	{"1", "Привет, Go!", "Go - отличный язык для API"},
	{"2", "REST за 5 минут", "Создаём API с помощью gorilla/mux"},
}

// Обработчик главной страницы
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

// Обработчик формы добавления
func FormHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	article := Article{
		ID:      r.FormValue("id"),
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}
	
	articles = append(articles, article)
	
	// Перенаправляем обратно на главную
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Получение всех статей
func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(articles)
}

func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/add", FormHandler).Methods("POST")
	router.HandleFunc("/articles", GetArticles).Methods("GET")
	
	fmt.Println("Сервер запущен: http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}