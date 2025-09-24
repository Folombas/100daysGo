package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🎯 День 61: Map в Go - Ключи к данным! 🗺️")
	fmt.Println("========================================")
	fmt.Println()

	// Анимация загрузки
	fmt.Print("🗺️ Загрузка карт данных ")
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("🚀 1. Базовые операции с Map:")
	demoBasicMaps()

	fmt.Println("\n🌈 2. Продвинутые Map:")
	demoAdvancedMaps()

	fmt.Println("\n⚡ 3. Операции с Map:")
	demoMapOperations()

	fmt.Println("\n🎨 4. Паттерны использования:")
	demoMapPatterns()

	fmt.Println("\n📊 5. Бенчмарки и сравнения:")
	demoMapBenchmarks()

	fmt.Println("\n✨ Демонстрация завершена! Теперь ты мастер Map! 🎓")
	fmt.Println("💫 Map - это мощный инструмент для работы с данными! 🌟")
}

func printSeparator() {
	fmt.Println("——————————————————————————————————————————————")
}
