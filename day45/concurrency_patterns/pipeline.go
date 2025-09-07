package main

import (
	"fmt"
	"math/rand"
)

// Pipeline Pattern
func PipelineDemo() {
	fmt.Println("\n=== Pipeline Pattern ===")
	
	// Генерация чисел
	numbers := generateNumbers(10)
	
	// Конвейер обработки
	squared := square(numbers)
	cubed := cube(squared)
	
	// Вывод результатов
	for result := range cubed {
		fmt.Printf("Результат: %d\n", result)
	}
}

func generateNumbers(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- rand.Intn(10) + 1
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func cube(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num * num
		}
		close(out)
	}()
	return out
}