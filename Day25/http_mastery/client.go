package main

import (
	"encoding/json"
	//"fmt"
	"io"
	"net/http"
	"time"
)

// Структура для внешнего API
type ExternalAPIResponse struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

// Обработчик внешнего API
func externalAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Создаем HTTP-клиент с таймаутами
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Делаем запрос к внешнему API
	resp, err := client.Get("http://ip-api.com/json/")
	if err != nil {
		http.Error(w, "Ошибка внешнего API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения ответа", http.StatusInternalServerError)
		return
	}

	// Парсим JSON
	var apiResponse ExternalAPIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		http.Error(w, "Ошибка парсинга JSON", http.StatusInternalServerError)
		return
	}

	// Формируем ответ
	response := map[string]interface{}{
		"external_api": "ip-api.com",
		"your_ip_info": apiResponse,
		"request_time": time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
