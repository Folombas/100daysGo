package main

import "fmt"

func demoBasicMaps() {
	printSeparator()

	fmt.Println("🆕 Создание Map разными способами:")

	// Способ 1: make
	ages := make(map[string]int)
	ages["Анна"] = 25
	ages["Петр"] = 30
	fmt.Printf("📊 Map через make: %v\n", ages)

	// Способ 2: литерал
	capitals := map[string]string{
		"Россия": "Москва 🏛️",
		"Франция": "Париж 🗼",
		"Япония": "Токио 🗾",
	}
	fmt.Printf("🌍 Map через литерал: %v\n", capitals)

	// Способ 3: пустая Map
	emptyMap := map[string]bool{}
	fmt.Printf("📦 Пустая Map: %v\n", emptyMap)

	fmt.Println("\n🔍 Доступ к элементам:")
	fmt.Printf("Столица России: %s\n", capitals["Россия"])
	fmt.Printf("Возраст Анны: %d лет\n", ages["Анна"])

	// Проверка существования ключа
	if capital, exists := capitals["Германия"]; exists {
		fmt.Printf("Столица Германии: %s\n", capital)
	} else {
		fmt.Println("❌ Германия не найдена в нашей мапе")
	}
}
