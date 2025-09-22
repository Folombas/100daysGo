package main

import "fmt"

// Демонстрация работы с указателями
func demoPointers() {
	fmt.Println("Работа с указателями:")

	// Объявление переменной и указателя (как в C)
	var value int = 42
	var pointer *int = &value // & - взятие адреса (как в C)

	fmt.Printf("Значение: %d\n", value)
	fmt.Printf("Адрес: %p\n", pointer)
	fmt.Printf("Значение через указатель: %d\n", *pointer) // * - разыменование

	// Изменение значения через указатель
	*pointer = 100
	fmt.Printf("Новое значение: %d\n", value)

	// Указатели в структурах (как в C)
	type Point struct {
		X, Y int
	}

	p := Point{10, 20}
	pp := &p
	pp.X = 30 // Автоматическое разыменование (удобнее чем в C)
	fmt.Printf("Точка: (%d, %d)\n", p.X, p.Y)
}
