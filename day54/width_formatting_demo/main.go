// main.go
package main

import (
	"fmt"
	"log"

	"width_formatting_demo/web"
)

func main() {
	fmt.Println("🚀 Day 54x: Ширина под контролем — Printf как типограф в Go!")
	fmt.Println("📚 Демонстрация форматирования ширины в консоли и вебе.\n")

	demoConsoleWidthFormatting()
	fmt.Println("\n🌐 Запускаем веб-сервер на http://localhost:8080")
	fmt.Println("   Нажмите Ctrl+C для выхода.\n")

	if err := web.StartServer(); err != nil {
		log.Fatal("❌ Ошибка запуска веб-сервера:", err)
	}
}

func demoConsoleWidthFormatting() {
	fmt.Println("=== 🖥️  Демонстрация форматирования ширины в консоли ===\n")

	// Пример 1: Фиксированная ширина — выравнивание по правому краю
	fmt.Println("🔹 Пример 1: Выравнивание по правому краю (ширина 10)")
	fmt.Printf("|%10s|\n", "Go")
	fmt.Printf("|%10s|\n", "Golang")
	fmt.Printf("|%10s|\n", "Привет")
	fmt.Println()

	// Пример 2: Выравнивание по левому краю
	fmt.Println("🔹 Пример 2: Выравнивание по левому краю (ширина 10)")
	fmt.Printf("|%-10s|\n", "Go")
	fmt.Printf("|%-10s|\n", "Golang")
	fmt.Printf("|%-10s|\n", "Привет")
	fmt.Println()

	// Пример 3: Заполнение нулями
	fmt.Println("🔹 Пример 3: Заполнение числовой ширины нулями")
	fmt.Printf("ID: %06d\n", 42)
	fmt.Printf("Код: %08d\n", 123)
	fmt.Printf("Счёт: %05d\n", 7)
	fmt.Println()

	// Пример 4: Динамическая ширина через *
	width := 15
	fmt.Println("🔹 Пример 4: Динамическая ширина (через *)")
	fmt.Printf("|%*s|\n", width, "Динамически")
	fmt.Printf("|%*s|\n", width, "Go")
	fmt.Printf("|%*s|\n", width, "Форматирование")
	fmt.Println()

	// Пример 5: Ширина для чисел с плавающей точкой
	fmt.Println("🔹 Пример 5: Ширина + точность для float")
	fmt.Printf("|%10.2f|\n", 3.14)
	fmt.Printf("|%10.2f|\n", 123.456)
	fmt.Printf("|%10.2f|\n", 9999.999)
	fmt.Println()

	// Пример 6: Комбинирование ширины, выравнивания и кириллицы
	fmt.Println("🔹 Пример 6: Таблица с кириллицей и выравниванием")
	headers := []string{"Имя", "Возраст", "Город"}
	data := [][]string{
		{"Алексей", "28", "Москва"},
		{"Мария", "32", "Санкт-Петербург"},
		{"Дмитрий", "25", "Новосибирск"},
	}

	// Печать заголовков
	fmt.Printf("| %-10s | %6s | %-15s |\n", headers[0], headers[1], headers[2])
	fmt.Println("|------------|--------|-----------------|")

	// Печать данных
	for _, row := range data {
		fmt.Printf("| %-10s | %6s | %-15s |\n", row[0], row[1], row[2])
	}
	fmt.Println()

	fmt.Println("✅ Консольная демонстрация завершена!")
}
