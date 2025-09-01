package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func runLoadTest() {
	host := getEnv("HOST", "localhost")
	port := getEnv("PORT", "8080")
	url := fmt.Sprintf("http://%s:%s/health", host, port)
	requests := 100
	concurrency := 10

	fmt.Printf("🔥 Запуск нагрузочного теста\n")
	fmt.Printf("   URL: %s\n", url)
	fmt.Printf("   Запросов: %d\n", requests)
	fmt.Printf("   Конкурентность: %d\n", concurrency)
	fmt.Println("")

	var wg sync.WaitGroup
	var mu sync.Mutex
	successful := 0
	failed := 0
	totalTime := time.Duration(0)

	startTest := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			client := &http.Client{Timeout: 5 * time.Second}
			requestsPerWorker := requests / concurrency

			for j := 0; j < requestsPerWorker; j++ {
				start := time.Now()
				resp, err := client.Get(url)
				duration := time.Since(start)

				mu.Lock()
				totalTime += duration
				if err != nil || resp.StatusCode != 200 {
					failed++
				} else {
					successful++
				}
				mu.Unlock()

				if resp != nil {
					resp.Body.Close()
				}

				if workerID == 0 && j%10 == 0 {
					fmt.Printf("⏳ Выполнено запросов: %d/%d\n", successful+failed, requests)
				}
			}
		}(i)
	}

	wg.Wait()

	testDuration := time.Since(startTest)
	avgTime := totalTime / time.Duration(requests)

	fmt.Println("")
	fmt.Printf("📊 Результаты нагрузочного теста:\n")
	fmt.Printf("   Успешных: %d\n", successful)
	fmt.Printf("   Неудачных: %d\n", failed)
	fmt.Printf("   Общее время: %v\n", testDuration)
	fmt.Printf("   Среднее время: %v\n", avgTime)
	fmt.Printf("   Запросов в секунду: %.2f\n", float64(requests)/testDuration.Seconds())
	
	if failed > 0 {
		os.Exit(1)
	}
}