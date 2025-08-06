package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Структура для декодирования JSON
type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func main() {
	// Создаем HTTP-клиент с таймаутом
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 1. GET-запрос к публичному API
	fmt.Println("▶️ GET-запрос к JSONPlaceholder")
	getExample(client)

	// 2. POST-запрос с созданием нового поста
	fmt.Println("\n▶️ POST-запрос с созданием ресурса")
	postExample(client)
}

func getExample(client *http.Client) {
	// Отправляем GET-запрос
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalf("Ошибка GET-запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Неверный статус: %s", resp.Status)
	}

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка чтения ответа: %v", err)
	}

	// Декодируем JSON
	var post Post
	if err := json.Unmarshal(body, &post); err != nil {
		log.Fatalf("Ошибка декодирования JSON: %v", err)
	}

	// Выводим результат
	fmt.Printf("📨 Получен пост #%d:\n", post.ID)
	fmt.Printf("Заголовок: %s\n", post.Title)
	fmt.Printf("Текст: %s\n", post.Body)
}

func postExample(client *http.Client) {
	// Создаем данные для отправки
	newPost := Post{
		Title:  "Мой пост",
		Body:   "Создано в Day12 марафона Go!",
		UserID: 1,
	}

	// Кодируем в JSON
	jsonData, err := json.Marshal(newPost)
	if err != nil {
		log.Fatalf("Ошибка кодирования JSON: %v", err)
	}

	// Отправляем POST-запрос
	resp, err := client.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatalf("Ошибка POST-запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем статус
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Неверный статус: %s", resp.Status)
	}

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка чтения ответа: %v", err)
	}

	// Декодируем ответ
	var createdPost Post
	if err := json.Unmarshal(body, &createdPost); err != nil {
		log.Fatalf("Ошибка декодирования JSON: %v", err)
	}

	// Выводим результат
	fmt.Printf("✅ Создан новый пост!\n")
	fmt.Printf("ID: %d\n", createdPost.ID)
	fmt.Printf("Статус: %s\n", resp.Status)
}