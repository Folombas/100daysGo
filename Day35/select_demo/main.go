package main

import (
	"fmt"
)

func main() {
	fmt.Println("🎯 Демонстрация select и default в Go")
	fmt.Println("====================================")
	fmt.Println()

	// Базовые примеры
	fmt.Println("1. Базовые примеры select:")
	fmt.Println("-------------------------")
	basicSelectExamples()
	fmt.Println()

	// Пример с default
	fmt.Println("2. Примеры с default:")
	fmt.Println("---------------------")
	defaultExamples()
	fmt.Println()

	// Таймауты
	fmt.Println("3. Таймауты с select:")
	fmt.Println("---------------------")
	timeoutExamples()
	fmt.Println()

	// Мультиплексирование
	fmt.Println("4. Мультиплексирование каналов:")
	fmt.Println("-------------------------------")
	multiplexingExamples()
	fmt.Println()

	// Бесконечные циклы
	fmt.Println("5. Бесконечные циклы с select:")
	fmt.Println("------------------------------")
	infiniteLoopExamples()
	fmt.Println()

	// Запуск веб-демо
	fmt.Println("6. Запуск веб-демонстрации:")
	fmt.Println("---------------------------")
	fmt.Println("Запустите отдельно: go run demo/web_demo.go")
	fmt.Println("И откройте http://localhost:8080 в браузере")
	fmt.Println()
}