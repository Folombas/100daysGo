package main

import (
	"net/http"
	"html/template"
	//"path/filepath"
)

// Главная страница
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, struct{}{})
}

// Создание файла/папки
func createHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	content := r.FormValue("content")
	isDir := r.FormValue("isDir") == "on"

	if isDir {
		createDir(name)
	} else {
		writeToFile(name, content)
	}
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Чтение файла
func readHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	content, _ := readFile(path)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(content))
}

// Удаление файла/папки
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	deletePath(path)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Список файлов
func listHandler(w http.ResponseWriter, r *http.Request) {
	files, _ := listDir(".")
	w.Header().Set("Content-Type", "application/json")
	w.Write(files)
}