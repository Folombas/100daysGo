// main.go
package main

import (
	"fmt"
	"log"

	"packages_modules_demo/calculator"
	"packages_modules_demo/formatter"
	"packages_modules_demo/web"
)

func main() {
	fmt.Println("🚀 Day 54: Пакеты, модули и импорты — Архитектура твоего кода!")
	fmt.Println("📚 Демонстрация работы с внутренними пакетами и модулями...")

	// Используем пакет calculator
	a, b := 15.5, 7.3
	sum := calculator.Add(a, b)
	diff := calculator.Subtract(a, b)
	product := calculator.Multiply(a, b)

	// Исправлено: принимаем оба значения от Divide
	quotient, ok := calculator.Divide(a, b)
	if !ok {
		fmt.Printf("🔹 Деление: невозможно (деление на ноль)")
	} else {
		formatter.PrintResult("Деление", a, b, quotient)
	}

	// Выводим остальные операции
	formatter.PrintResult("Сложение", a, b, sum)
	formatter.PrintResult("Вычитание", a, b, diff)
	formatter.PrintResult("Умножение", a, b, product)

	fmt.Println("\n🌐 Запускаем веб-демо на http://localhost:8080")
	fmt.Println("   Нажмите Ctrl+C для выхода.")

	// Запуск веб-сервера
	if err := web.StartServer(); err != nil {
		log.Fatal("❌ Ошибка запуска веб-сервера:", err)
	}
}
