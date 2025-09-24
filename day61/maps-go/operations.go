package main

import "fmt"

func demoMapOperations() {
	printSeparator()

	fmt.Println("🔄 Операции с Map:")

	products := map[string]float64{
		"🍎 Яблоки":  150.50,
		"🍌 Бананы":  89.90,
		"🥛 Молоко":  75.00,
		"🍞 Хлеб":    45.30,
	}

	fmt.Println("🛒 Исходный список продуктов:")
	for product, price := range products {
		fmt.Printf("  %s: %.2f руб.\n", product, price)
	}

	// Добавление элемента
	products["🧀 Сыр"] = 320.00
	fmt.Println("\n✅ Добавили сыр:", products["🧀 Сыр"])

	// Обновление элемента
	products["🍞 Хлеб"] = 42.50
	fmt.Println("✏️ Обновили цену хлеба:", products["🍞 Хлеб"])

	// Удаление элемента
	delete(products, "🍌 Бананы")
	fmt.Println("❌ Удалили бананы")

	// Проверка длины
	fmt.Printf("📏 Теперь в карте %d элементов\n", len(products))

	fmt.Println("\n🛍️ Обновленный список:")
	total := 0.0
	for product, price := range products {
		fmt.Printf("  %s: %.2f руб.\n", product, price)
		total += price
	}
	fmt.Printf("💰 Общая стоимость: %.2f руб.\n", total)

	// Очистка Map
	clear(products)
	fmt.Printf("🧹 Очистили карту. Теперь элементов: %d\n", len(products))
}
