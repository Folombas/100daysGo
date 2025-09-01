package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func startClient() {
	host := getEnv("HOST", "localhost")
	port := getEnv("PORT", "8080")
	url := fmt.Sprintf("http://%s:%s", host, port)

	fmt.Printf("🌐 Запуск клиента для %s\n", url)
	fmt.Println("")

	endpoints := []string{
		"/health",
		"/stats",
		"/api/users",
		"/api/network",
	}

	for _, endpoint := range endpoints {
		testEndpoint(url + endpoint)
	}
}

func testEndpoint(url string) {
	client := &http.Client{Timeout: 5 * time.Second}

	start := time.Now()
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		fmt.Printf("❌ %s: Ошибка - %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ %s: Ошибка чтения ответа - %v\n", url, err)
		return
	}

	var formattedBody string
	if resp.Header.Get("Content-Type") == "application/json" {
		var prettyJSON map[string]interface{}
		if json.Unmarshal(body, &prettyJSON) == nil {
			pretty, _ := json.MarshalIndent(prettyJSON, "", "  ")
			formattedBody = string(pretty)
		} else {
			formattedBody = string(body)
		}
	} else {
		formattedBody = string(body)
	}

	fmt.Printf("✅ %s\n", url)
	fmt.Printf("   Статус: %d\n", resp.StatusCode)
	fmt.Printf("   Время: %v\n", duration)
	fmt.Printf("   Ответ:\n%s\n", formattedBody)
	fmt.Println("---")
}