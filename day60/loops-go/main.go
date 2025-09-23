package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🎯 День 60: Циклы в Go - Вращайся и Повторяй! 🔄")
	fmt.Println("==============================================")
	fmt.Println()

	// Анимация загрузки
	fmt.Print("🔄 Загрузка модуля циклов ")
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("🚀 1. Базовые циклы:")
	demoBasicLoops()

	fmt.Println("\n🔮 2. Продвинутые циклы:")
	demoAdvancedLoops()

	fmt.Println("\n🎨 3. Паттерны с циклами:")
	demoPatterns()

	fmt.Println("\n⚡ 4. Бенчмарки производительности:")
	demoBenchmarks()

	fmt.Println("\n✨ Демонстрация завершена! Теперь ты мастер циклов! 🎓")
	fmt.Println("💫 Поздравляю с 60-м днем марафона! Ты великолепен! 🌟")
}

// Простая функция для создания разделителя
func printSeparator() {
	fmt.Println("——————————————————————————————————————————————")
}
