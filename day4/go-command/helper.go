package main

import "fmt"

// HelperFunction демонстрирует работу с несколькими файлами
func HelperFunction() string {
	return "🛠️ Это функция из helper.go!"
}

// Calculate простой калькулятор для демонстрации
func Calculate(a, b int) int {
	return a * b
}

// PrintMessage печатает сообщение
func PrintMessage(msg string) {
	fmt.Printf("📢 Сообщение: %s\n", msg)
}

