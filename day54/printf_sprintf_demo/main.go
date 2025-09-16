// main.go
package main

import (
	"fmt"
	"log"

	"printf_sprintf_demo/web"
)

func main() {
	fmt.Println("🚀 Day 54: Printf & Sprintf — Мастерство форматированного вывода в Go!")
	fmt.Println("📚 Запускаем демонстрацию в консоли и веб-интерфейсе...\n")

	demoConsole()
	fmt.Println("\n🌐 Запускаем веб-сервер на http://localhost:8080")
	fmt.Println("   Нажмите Ctrl+C для выхода.\n")

	// Запуск веб-сервера
	if err := web.StartServer(); err != nil {
		log.Fatal("❌ Ошибка запуска веб-сервера:", err)
	}
}

func demoConsole() {
	fmt.Println("=== 🖥️  Демонстрация в консоли ===")

	// Базовые примеры Printf
	fmt.Printf("Привет, %s! Сегодня %d-й день марафона.\n", "друг", 54)
	fmt.Printf("Точность числа Пи: %.5f\n", 3.1415926535)
	fmt.Printf("Целое число с ведущими нулями: %05d\n", 42)

	// Выравнивание
	fmt.Printf("Выравнивание вправо: %10s\n", "Go")
	fmt.Printf("Выравнивание влево: %-10s\n", "Go")

	// Sprintf — форматирование в строку
	name := "Алексей"
	age := 28
	message := fmt.Sprintf("Меня зовут %s, мне %d лет.", name, age)
	fmt.Println("📩 Сформированное сообщение:", message)

	// Комплексный пример
	product := "Ноутбук"
	price := 89999.99
	quantity := 3
	total := price * float64(quantity)
	fmt.Printf("🛒 %s x%d = %.2f руб.\n", product, quantity, total)

	// Логические и специальные символы
	isReady := true
	fmt.Printf("Готово: %t\n", isReady)
	fmt.Printf("Проценты: %d%%\n", 99)

	// Unicode и кириллица — всё работает!
	city := "Москва"
	fmt.Printf("📍 Город: %s, длина строки: %d символов\n", city, len(city))

	fmt.Println("✅ Консольная демонстрация завершена!")
}
