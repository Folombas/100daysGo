package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

// Добавляем функцию getEnv
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func runBenchmark() {
	fmt.Println("🏃 Запуск бенчмарк-тестов...")
	fmt.Println("")

	// Benchmark HTTP requests
	benchmarkHTTP()
}

func benchmarkHTTP() {
	host := getEnv("HOST", "localhost")
	port := getEnv("PORT", "8080")
	url := fmt.Sprintf("http://%s:%s/health", host, port)

	client := &http.Client{Timeout: 5 * time.Second}

	// Warm up
	for i := 0; i < 10; i++ {
		client.Get(url)
	}

	// Benchmark
	start := time.Now()
	requests := 1000

	for i := 0; i < requests; i++ {
		resp, err := client.Get(url)
		if err == nil {
			resp.Body.Close()
		}
	}

	duration := time.Since(start)
	rps := float64(requests) / duration.Seconds()

	fmt.Printf("📊 HTTP Бенчмарк:\n")
	fmt.Printf("   Запросов: %d\n", requests)
	fmt.Printf("   Время: %v\n", duration)
	fmt.Printf("   Запросов в секунду: %.2f\n", rps)
	fmt.Printf("   Среднее время запроса: %.2f ms\n", duration.Seconds()/float64(requests)*1000)
}

// Для интеграции с testing package
func BenchmarkHealthEndpoint(b *testing.B) {
	host := getEnv("HOST", "localhost")
	port := getEnv("PORT", "8080")
	url := fmt.Sprintf("http://%s:%s/health", host, port)

	client := &http.Client{Timeout: 5 * time.Second}

	for i := 0; i < b.N; i++ {
		resp, err := client.Get(url)
		if err == nil {
			resp.Body.Close()
		}
	}
}