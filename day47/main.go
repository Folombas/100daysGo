package main

import (
	"day47/examples"
	"day47/factory"
	"fmt"
	"time"
)

func main() {
	fmt.Println("🏭 Day 47: Горутины и каналы - Фабрика параллелизма")
	fmt.Println("==================================================")

	// Демонстрация работы фабрики
	fmt.Println("\n1. Работа фабрики с конвейером:")
	factory.DemoFactory()

	// Небольшая пауза для наглядности
	time.Sleep(1 * time.Second)

	// Демонстрация конвейерной обработки
	fmt.Println("\n2. Конвейерная обработка данных:")
	examples.DemoPipeline()

	// Демонстрация паттернов
	fmt.Println("\n3. Паттерны работы с горутинами:")
	examples.DemoPatterns()

	fmt.Println("\n🎉 Демонстрация завершена! Фабрика работает исправно!")
}