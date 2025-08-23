package main

import "fmt"

// Calculator предоставляет базовые арифметические операции
type Calculator struct{}

// Add возвращает сумму двух чисел
func (c Calculator) Add(a, b int) int {
    return a + b
}

// Subtract возвращает разность двух чисел
func (c Calculator) Subtract(a, b int) int {
    return a - b
}

// Multiply возвращает произведение двух чисел
func (c Calculator) Multiply(a, b int) int {
    return a * b
}

// Divide возвращает результат деления a на b
func (c Calculator) Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль")
    }
    return a / b, nil
}