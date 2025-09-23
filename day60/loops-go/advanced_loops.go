package main

import (
	"fmt"
)

func demoAdvancedLoops() {
	printSeparator()

	fmt.Println("📚 Цикл по массиву:")
	fruits := [5]string{"🍎", "🍌", "🍇", "🍊", "🍓"}
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("📦 fruits[%d] = %s\n", i, fruits[i])
	}

	fmt.Println("\n🎪 Цикл range по массиву:")
	for index, fruit := range fruits {
		fmt.Printf("📍 Индекс: %d, Фрукт: %s\n", index, fruit)
	}

	fmt.Println("\n🌊 Цикл range по слайсу:")
	numbers := []int{10, 20, 30, 40, 50}
	for i, num := range numbers {
		fmt.Printf("🔢 numbers[%d] = %d\n", i, num)
	}

	fmt.Println("\n🗺️ Цикл range по map:")
	capitals := map[string]string{
		"Россия":   "Москва 🏛️",
		"Франция":  "Париж 🗼",
		"Япония":   "Токио 🗾",
		"Бразилия": "Бразилиа 🌴",
	}
	for country, capital := range capitals {
		fmt.Printf("🌍 %s → %s\n", country, capital)
	}

	fmt.Println("\n📝 Цикл range по строке:")
	message := "Привет 🚀"
	for i, char := range message {
		fmt.Printf("🔡 Символ %d: %c (код: %d)\n", i, char, char)
	}
}
