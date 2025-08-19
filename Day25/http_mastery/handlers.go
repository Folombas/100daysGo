package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Базовый обработчик
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}
	
	response := map[string]string{
		"message": "Привет, мир! 🌍",
		"status":  "success",
		"time":    time.Now().Format(time.RFC3339),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Возвращает текущее время
func timeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"time":    time.Now().Format("2006-01-02 15:04:05"),
		"timezone": "Europe/Moscow",
		"timestamp": fmt.Sprintf("%d", time.Now().Unix()),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Эхо-ответ с полученными данными
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST запросы", http.StatusMethodNotAllowed)
		return
	}
	
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Неверный JSON", http.StatusBadRequest)
		return
	}
	
	response := map[string]interface{}{
		"received": data,
		"metadata": map[string]interface{}{
			"headers": r.Header,
			"method":  r.Method,
			"url":     r.URL.String(),
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}