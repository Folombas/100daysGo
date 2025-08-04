package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Обработчик поиска (query parameters)
func searchHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем query-параметры
	query := r.URL.Query()

	// Извлекаем параметры
	q := query.Get("q")
	category := query.Get("category")
	page := query.Get("page")
	if page == "" {
		page = "1"
	}

	// Экранируем пользовательский ввод
	safeQ := url.QueryEscape(q)

	// Формируем ответ
	response := map[string]interface{}{
		"status":   "success",
		"query":    q,
		"safe":     safeQ,
		"category": category,
		"page":     page,
		"results": []map[string]string{
			{"id": "1", "title": "Результат 1 по запросу " + q},
			{"id": "2", "title": "Еще один результат в категории " + category},
		},
		"message": "Поисковый запрос обработан успешно!",
	}

	// Отправляем JSON
	json.NewEncoder(w).Encode(response)
}

// Обработчик пользователя (path parameters)
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр из пути
	id := r.PathValue("id")

	// Эмуляция данных пользователя
	user := map[string]interface{}{
		"id":       id,
		"name":     "Гоша Гофер",
		"email":    fmt.Sprintf("user%s@example.com", id),
		"role":     "admin",
		"reg_date": "2025-01-15",
	}

	// Формируем ответ
	response := map[string]interface{}{
		"status": "success",
		"user":   user,
		"links": map[string]string{
			"self":    r.URL.String(),
			"profile": "/user/" + id + "/profile",
		},
	}

	json.NewEncoder(w).Encode(response)
}

// Обработчик регистрации (form data)
func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим форму
	err := r.ParseForm()
	if err != nil {
		http.Error(w, `{"status":"error","message":"Неверный формат данных"}`, http.StatusBadRequest)
		return
	}

	// Получаем данные из формы
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Простая валидация
	if name == "" || email == "" || password == "" {
		http.Error(w, `{"status":"error","message":"Все поля обязательны для заполнения"}`, http.StatusBadRequest)
		return
	}

	// Формируем ответ
	response := map[string]interface{}{
		"status":  "success",
		"message": "Пользователь зарегистрирован!",
		"data": map[string]string{
			"name":  name,
			"email": email,
			"id":    "1001",
		},
	}

	json.NewEncoder(w).Encode(response)
}

// Обработчик JSON данных
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем Content-Type
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, `{"status":"error","message":"Требуется application/json"}`, http.StatusUnsupportedMediaType)
		return
	}

	// Декодируем JSON
	var requestData struct {
		Product  string  `json:"product"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, `{"status":"error","message":"Неверный формат JSON"}`, http.StatusBadRequest)
		return
	}

	// Обработка данных
	total := requestData.Price * float64(requestData.Quantity)

	// Формируем ответ
	response := map[string]interface{}{
		"status":   "success",
		"product":  requestData.Product,
		"price":    requestData.Price,
		"quantity": requestData.Quantity,
		"total":    total,
		"message":  "Данные успешно обработаны!",
	}

	json.NewEncoder(w).Encode(response)
}
