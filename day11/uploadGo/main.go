package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Статическая раздача HTML формы
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "upload.html")
	})

	// Обработчик загрузки файлов
	http.HandleFunc("/upload", uploadHandler)

	// Создаем папку для загрузок
	if err := os.MkdirAll("upload", 0755); err != nil {
		log.Fatal("Не удалось создать папку upload:", err)
	}

	port := ":8080"
	log.Printf("Сервер запущен на http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Ограничение 20MB
	if err := r.ParseMultipartForm(20 << 20); err != nil {
		http.Error(w, "Файл слишком большой (макс. 20MB)", http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка при получении файла: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Создаем файл на сервере
	targetPath := filepath.Join("upload", filepath.Base(header.Filename))
	dst, err := os.Create(targetPath)
	if err != nil {
		http.Error(w, "Ошибка создания файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Копируем содержимое
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Ошибка сохранения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Форматируем ответ
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head><meta charset="UTF-8"><title>Успешно!</title></head>
	<body>
		<h1>Файл успешно загружен!</h1>
		<p>Имя: %s</p>
		<p>Размер: %.2f KB</p>
		<a href="/">Загрузить ещё</a>
	</body>
	</html>
	`, header.Filename, float64(header.Size)/1024)
}