package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🚀 100aysG0: Reboot - Day 1")
	fmt.Println("📚 История языка Go: от 2007 до 2025")
	fmt.Println("=====================================")

	// Анимированное приветствие
	for i := 0; i < 3; i++ {
		fmt.Print("Загрузка истории Go")
		for j := 0; j < 3; j++ {
			fmt.Print(".")
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Print("\r")
	}

	fmt.Println("\n🎯 Добро пожаловать в путешествие во времени!")

	// Запуск компонентов
	DisplayTimeline()
	fmt.Println()
	StartInteractiveQuiz()
}
