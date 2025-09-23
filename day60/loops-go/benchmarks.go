package main

import (
	"fmt"
	"time"
)

func demoBenchmarks() {
	printSeparator()

	fmt.Println("⚡ Сравнение производительности циклов:")

	// Подготовка данных
	size := 1000000
	slice := make([]int, size)
	for i := range slice {
		slice[i] = i
	}

	// Тест 1: Классический for
	start := time.Now()
	sum1 := 0
	for i := 0; i < len(slice); i++ {
		sum1 += slice[i]
	}
	time1 := time.Since(start)

	// Тест 2: Range
	start = time.Now()
	sum2 := 0
	for _, value := range slice {
		sum2 += value
	}
	time2 := time.Since(start)

	// Тест 3: Range с индексом
	start = time.Now()
	sum3 := 0
	for i := range slice {
		sum3 += slice[i]
	}
	time3 := time.Since(start)

	fmt.Printf("📊 Результаты (размер данных: %d):\n", size)
	fmt.Printf("🔹 Классический for: %v, сумма: %d\n", time1, sum1)
	fmt.Printf("🔹 Range по значению: %v, сумма: %d\n", time2, sum2)
	fmt.Printf("🔹 Range по индексу: %v, сумма: %d\n", time3, sum3)

	fmt.Println("\n💡 Вывод: Range по индексу обычно быстрее для больших массивов!")

	// Демонстрация вложенных циклов
	fmt.Println("\n🎯 Вложенные циклы (матрица):")
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("📍 matrix[%d][%d] = %d\n", i, j, value)
		}
	}
}
