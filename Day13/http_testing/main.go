package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Регистрируем обработчики
	http.HandleFunc("/", homeHandler) // Добавляем обработчик для корня
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/echo", echoHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Обработчик для корневого пути
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Добро пожаловать на Day13!</h1>
		<ul>
			<li><a href="/hello">Поприветствовать</a></li>
			<li>Отправить сообщение (POST на /echo): используй curl или Postman</li>
		</ul>`)
}

// Обработчик для /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем UTF-8
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	// Проверяем метод
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Параметр из query string
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Гость"
	}

	// Пишем ответ
	fmt.Fprintf(w, "Привет, %s! Добро пожаловать в Day13!", name)
}

// Обработчик для /echo
func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	// Только POST запросы
	if r.Method != http.MethodPost {
		http.Error(w, "Требуется POST-запрос", http.StatusMethodNotAllowed)
		return
	}

	// Читаем тело запроса
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Возвращаем прочитанные данные
	fmt.Fprintf(w, "Получено: %s", string(data))
}