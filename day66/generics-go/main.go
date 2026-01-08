package main

import "fmt"

// PrintSlice - универсальная функция для вывода любого слайса.
// T - type parameter (параметр типа), может быть любым.
func PrintSlice[T any](s []T) {
	fmt.Print("Универсальный слайс: ")
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

// Container - универсальная структура-контейнер для одного значения.
type Container[T any] struct {
	Value T
}

func main() {
	fmt.Println("=== Day 66: Generics in Go ===")
	fmt.Println("Изучаем универсальные решения для кода и жизни.")

	// Пример 1: Слайсы разных типов
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{1.1, 2.2, 3.3}
	stringSlice := []string{"метель", "код", "фокус"}

	PrintSlice(intSlice)
	PrintSlice(floatSlice)
	PrintSlice(stringSlice)

	// Пример 2: Универсальный контейнер
	intContainer := Container[int]{Value: 66}
	strContainer := Container[string]{Value: "День 66"}

	fmt.Printf("\nКонтейнер с числом: %v\n", intContainer.Value)
	fmt.Printf("Контейнер со строкой: %v\n", strContainer.Value)

	fmt.Println("\nДженерики — это как единый проездной в мир эффективного кода.")
}
