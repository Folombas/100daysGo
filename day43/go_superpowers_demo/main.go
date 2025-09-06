package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("🚀 Day 43: Почему Go - Суперсилы современного разработчика")
	fmt.Println("==========================================================")
	
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	
	switch os.Args[1] {
	case "web":
		fmt.Println("🌐 Запуск веб-сервиса...")
		StartWebServer()
	case "cli":
		fmt.Println("🐚 Запуск CLI инструмента...")
		// Пропускаем первый аргумент (название команды)
		os.Args = os.Args[1:]
		StartCLI()
	case "concurrency":
		fmt.Println("⚡ Демонстрация конкурентности...")
		StartWorkerPool()
	case "benchmark":
		fmt.Println("🏎️  Запуск бенчмарков...")
		printBenchmarkInfo()
	default:
		fmt.Printf("❌ Неизвестная команда: %s\n", os.Args[1])
		printUsage()
	}
}

func printBenchmarkInfo() {
	fmt.Println("🏎️  Сравнение производительности:")
	fmt.Println("   Конкурентность vs Последовательность")
	
	// Запуск бенчмарков
	fmt.Println("Запустите для просмотра результатов:")
	fmt.Println("   go test -bench=. -benchmem")
}

func printUsage() {
	fmt.Println("Использование:")
	fmt.Println("  go run . web     - Запуск веб-сервиса")
	fmt.Println("  go run . cli     - Запуск CLI инструмента")
	fmt.Println("  go run . concurrency - Демонстрация конкурентности")
	fmt.Println("  go run . benchmark  - Информация о бенчмарках")
	fmt.Println()
	fmt.Println("Примеры CLI:")
	fmt.Println("  go run . cli --text=\"Привет мир\" --op=upper")
	fmt.Println("  go run . cli --text=\"ПРИВЕТ МИР\" --op=lower")
	fmt.Println("  go run . cli --text=\"привет мир\" --op=title")
	fmt.Println("  go run . cli --text=\"привет\" --op=reverse")
	fmt.Println()
	fmt.Println("Для запуска бенчмарков выполните:")
	fmt.Println("   go test -bench=. -benchmem")
}