package main

import (
	"fmt"
	"math"
)

// Интерфейс с методом для вычисления площади
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Структура Прямоугольник
type Rectangle struct {
	Width, Height float64
}

// Реализация методов интерфейса для Прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Структура Круг
type Circle struct {
	Radius float64
}

// Реализация методов интерфейса для Круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Функция, принимающая интерфейс Shape
func printShapeInfo(s Shape) {
	fmt.Printf("Площадь: %.2f, Периметр: %.2f\n", s.Area(), s.Perimeter())
}

// Демонстрация работы с интерфейсами
func demonstrateInterfaces() {
	fmt.Println("=== Демонстрация интерфейсов ===")
	
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2.5}
	
	fmt.Print("Прямоугольник: ")
	printShapeInfo(rect)
	
	fmt.Print("Круг: ")
	printShapeInfo(circle)
	
	// Пустой интерфейс (может содержать любое значение)
	fmt.Println("\n--- Пустой интерфейс ---")
	var emptyInterface interface{}
	
	emptyInterface = 42
	fmt.Printf("Значение: %v, Тип: %T\n", emptyInterface, emptyInterface)
	
	emptyInterface = "Строка"
	fmt.Printf("Значение: %v, Тип: %T\n", emptyInterface, emptyInterface)
}