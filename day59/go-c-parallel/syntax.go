package main

import "fmt"

// Демонстрация синтаксических сходств
func demoSyntax() {
	// Похожие конструкции циклов
	fmt.Println("Циклы (for):")

	// Как в C: for (i = 0; i < 5; i++)
	fmt.Println("Go вариант C-стиля:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Условные операторы
	fmt.Println("\nУсловные операторы (if/else):")
	x := 10
	if x > 5 {
		fmt.Println("x больше 5")
	} else {
		fmt.Println("x меньше или равно 5")
	}

	// Switch statement (похож на C, но с улучшениями)
	fmt.Println("\nSwitch statement:")
	switch x {
	case 5:
		fmt.Println("x равен 5")
	case 10:
		fmt.Println("x равен 10")
	default:
		fmt.Println("x другое значение")
	}
}
