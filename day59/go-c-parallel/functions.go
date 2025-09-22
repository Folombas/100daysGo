package main

import "fmt"

// Демонстрация функций и их особенностей
func demoFunctions() {
	fmt.Println("Функции:")

	// Базовые функции (похоже на C)
	result := add(5, 3)
	fmt.Printf("add(5, 3) = %d\n", result)

	// Возврат нескольких значений (улучшение относительно C)
	sum, diff := calculate(10, 4)
	fmt.Printf("calculate(10, 4) = сумма: %d, разность: %d\n", sum, diff)

	// Функции как first-class citizens (лучше чем в C)
	operation := func(a, b int) int {
		return a * b
	}
	fmt.Printf("Лямбда: 4 * 5 = %d\n", operation(4, 5))

	// Передача указателей в функции (как в C)
	value := 10
	fmt.Printf("До modifyPointer: %d\n", value)
	modifyPointer(&value)
	fmt.Printf("После modifyPointer: %d\n", value)
}

// Простая функция как в C
func add(a, b int) int {
	return a + b
}

// Возврат нескольких значений (улучшение)
func calculate(a, b int) (int, int) {
	return a + b, a - b
}

// Работа с указателями в функциях (как в C)
func modifyPointer(ptr *int) {
	*ptr = *ptr * 2
}
