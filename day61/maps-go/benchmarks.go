package main

import (
	"fmt"
	"time"
)

func demoMapBenchmarks() {
	printSeparator()

	fmt.Println("⚡ Сравнение производительности:")

	// Тест 1: Map vs Slice для поиска
	size := 100000
	testMap := make(map[int]string, size)
	testSlice := make([]string, size)

	// Заполнение данных
	for i := 0; i < size; i++ {
		value := fmt.Sprintf("value%d", i)
		testMap[i] = value
		testSlice[i] = value
	}

	// Поиск в Map
	start := time.Now()
	for i := 0; i < 1000; i++ {
		_ = testMap[i%size]
	}
	mapTime := time.Since(start)

	// Поиск в Slice (линейный)
	start = time.Now()
	for i := 0; i < 1000; i++ {
		target := i % size
		for j := 0; j < size; j++ {
			if j == target {
				_ = testSlice[j]
				break
			}
		}
	}
	sliceTime := time.Since(start)

	fmt.Printf("🔍 Поиск 1000 элементов:\n")
	fmt.Printf("  🗺️ Map: %v (O(1) - константное время)\n", mapTime)
	fmt.Printf("  📋 Slice: %v (O(n) - линейное время)\n", sliceTime)
	fmt.Printf("  📈 Map быстрее в %.0f раз!\n", float64(sliceTime)/float64(mapTime))

	// Тест 2: Итерация
	fmt.Println("\n🔄 Итерация по всем элементам:")

	start = time.Now()
	for range testMap {
		// Проход по всем элементам
	}
	mapIterTime := time.Since(start)

	start = time.Now()
	for range testSlice {
		// Проход по всем элементам
	}
	sliceIterTime := time.Since(start)

	fmt.Printf("  🗺️ Map итерация: %v\n", mapIterTime)
	fmt.Printf("  📋 Slice итерация: %v\n", sliceIterTime)
	fmt.Printf("  💡 Slice быстрее для итерации!\n")

	fmt.Println("\n🎯 Выводы:")
	fmt.Println("  ✅ Map отлично подходит для поиска по ключу")
	fmt.Println("  ✅ Slice лучше для последовательной обработки")
	fmt.Println("  💡 Выбирайте структуру данных по задаче!")
}
