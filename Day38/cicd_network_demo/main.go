package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "server":
		startServer()
	case "client":
		startClient()
	case "loadtest":
		runLoadTest()
	case "benchmark":
		runBenchmark() // Добавляем вызов функции
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("CI/CD Network Demo - Использование:")
	fmt.Println("  server    - запустить HTTP/TCP сервер")
	fmt.Println("  client    - запустить HTTP клиент")
	fmt.Println("  loadtest  - запустить нагрузочное тестирование")
	fmt.Println("  benchmark - запустить бенчмарк производительности")
	fmt.Println("")
	fmt.Println("Переменные окружения:")
	fmt.Println("  PORT=8080     - порт сервера")
	fmt.Println("  HOST=localhost - хост сервера")
	fmt.Println("  CI=true       - режим CI/CD")
}

// Добавляем недостающую функцию runBenchmark
func runBenchmark() {
	fmt.Println("🏃 Запуск бенчмарк-тестов...")
	fmt.Println("Реализация бенчмарк-тестов находится в файле benchmark.go")
	fmt.Println("Для запуска используйте: go run benchmark.go")
}