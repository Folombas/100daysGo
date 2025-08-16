package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	
	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

type User struct {
	ID    int
	Name  string
	Email string
}

var db *sql.DB

func main() {
	// Подключение к PostgreSQL
	connStr := "user=postgres password=mysecretpassword dbname=godb sslmode=disable host=localhost"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer db.Close()

	// Создание таблицы
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT
		);
	`)
	if err != nil {
		log.Println("Ошибка создания таблицы:", err)
	}

	// Роуты
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/static/style.css", cssHandler)
	
	log.Println("Сервер запущен: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Ошибка БД: "+err.Error(), 500)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			http.Error(w, "Ошибка сканирования: "+err.Error(), 500)
			return
		}
		users = append(users, u)
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, users)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 302)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	
	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	if err != nil {
		log.Println("Ошибка добавления:", err)
	}

	http.Redirect(w, r, "/", 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", 302)
		return
	}

	id := r.FormValue("id")
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Println("Ошибка удаления:", err)
	}

	http.Redirect(w, r, "/", 302)
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "style.css")
}