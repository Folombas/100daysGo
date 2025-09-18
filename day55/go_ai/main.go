// main.go
package main

import (
	"fmt"
	"log"

	"go_ai/ai_simulator"
	"go_ai/system_checker"
	"go_ai/web"
)

func main() {
	fmt.Println("🚀 Day 55: Go & ИИ — Кодируй будущее, даже если железо из прошлого!")
	fmt.Println("📚 Демонстрация: как Go помогает в мире ИИ без мощного GPU...")

	// Проверка системы
	system_checker.PrintSystemInfo()

	// Симуляция "лёгкой ИИ-задачи" — генерация текста через шаблон (без ML!)
	fmt.Println("\n🧪 Симуляция генерации ИИ-описания...")
	prompt := "расскажи о человеке, который учит Go, чтобы купить новую видеокарту"
	result := ai_simulator.GenerateText(prompt)
	fmt.Printf("💬 Результат симуляции:\n%s\n", result)

	fmt.Println("\n🌐 Запускаем веб-демо на http://localhost:8080")
	fmt.Println("   Нажмите Ctrl+C для выхода.")

	// Запуск веб-сервера
	if err := web.StartServer(); err != nil {
		log.Fatal("❌ Ошибка запуска веб-сервера:", err)
	}
}
