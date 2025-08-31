package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func StartHTTPClient() {
	fmt.Println("🌐 Запуск HTTP клиента...")
	fmt.Println()

	// Тестирование основных эндпоинтов
	testEndpoint("http://localhost:8080/api/health", "GET", nil)
	testEndpoint("http://localhost:8080/api/time", "GET", nil)
	testEndpoint("http://localhost:8080/api/users", "GET", nil)
	testEndpoint("http://localhost:8080/network/test", "GET", nil)

	// Тестирование защищенных эндпоинтов
	fmt.Println("🔐 Тестирование защищенных эндпоинтов:")
	testProtectedEndpoint("http://localhost:8080/admin")
	testProtectedEndpoint("http://localhost:8080/admin/dashboard")

	// Тестирование с правильным токеном
	fmt.Println("✅ Тестирование с правильным токеном:")
	testWithToken("http://localhost:8080/admin", "Bearer secret-token-123")
}

func testEndpoint(url, method string, body []byte) {
	fmt.Printf("📤 Запрос: %s %s\n", method, url)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		fmt.Printf("❌ Ошибка создания запроса: %v\n", err)
		return
	}

	start := time.Now()
	resp, err := client.Do(req)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("❌ Ошибка запроса: %v\n", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Ошибка чтения ответа: %v\n", err)
		return
	}

	fmt.Printf("📥 Ответ: %d (%s)\n", resp.StatusCode, duration)
	fmt.Printf("📦 Тело ответа: %s\n\n", string(responseBody))
}

func testProtectedEndpoint(url string) {
	fmt.Printf("📤 Запрос к защищенному эндпоинту: %s\n", url)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)

	if err != nil {
		fmt.Printf("❌ Ошибка запроса: %v\n", err)
		return
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("📥 Ответ: %d\n", resp.StatusCode)
	fmt.Printf("📦 Тело ответа: %s\n\n", string(responseBody))
}

func testWithToken(url, token string) {
	fmt.Printf("📤 Запрос с токеном: %s\n", url)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("❌ Ошибка создания запроса: %v\n", err)
		return
	}

	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Ошибка запроса: %v\n", err)
		return
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("📥 Ответ: %d\n", resp.StatusCode)
	fmt.Printf("📦 Тело ответа: %s\n\n", string(responseBody))
}