package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("🚀 Day 41: Особенности языка Go")
	fmt.Println(strings.Repeat("=", 50))
	
	// Демонстрация конкурентности
	demonstrateGoroutines()
	demonstrateMutex()
	demonstrateChannels()
	
	fmt.Println("\n" + strings.Repeat("=", 50))
	
	// Демонстрация интерфейсов
	demonstrateInterfaces()
	
	fmt.Println("\n" + strings.Repeat("=", 50))
	
	// Демонстрация обработки ошибок
	demonstrateErrorHandling()
	
	fmt.Println("\n" + strings.Repeat("=", 50))
	
	// Другие особенности Go
	demonstrateOtherFeatures()
}

func demonstrateOtherFeatures() {
	fmt.Println("=== Другие особенности Go ===")
	
	// Множественное возвращение значений
	fmt.Println("--- Множественное возвращение значений ---")
	a, b := multipleReturn()
	fmt.Printf("Возвращенные значения: %d, %s\n", a, b)
	
	// Короткое объявление переменных
	fmt.Println("\n--- Короткое объявление переменных ---")
	x := 42
	y := "Привет"
	fmt.Printf("x = %d, y = %s\n", x, y)
	
	// Range по коллекциям
	fmt.Println("\n--- Range по коллекциям ---")
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Индекс: %d, Значение: %d\n", index, value)
	}
	
	// Замыкания
	fmt.Println("\n--- Замыкания ---")
	counter := createCounter()
	fmt.Println("Счетчик:", counter())
	fmt.Println("Счетчик:", counter())
	fmt.Println("Счетчик:", counter())
}

func multipleReturn() (int, string) {
	return 42, "ответ"
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}