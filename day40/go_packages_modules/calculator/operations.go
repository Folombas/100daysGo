package calculator

import "fmt"

// Add возвращает сумму двух чисел
func Add(a, b float64) float64 {
    return a + b
}

// Subtract возвращает разность двух чисел
func Subtract(a, b float64) float64 {
    return a - b
}

// Multiply возвращает произведение двух чисел
func Multiply(a, b float64) float64 {
    return a * b
}

// Divide возвращает результат деления a на b
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль невозможно")
    }
    return a / b, nil
}

// init функция вызывается автоматически при импорте пакета
func init() {
    fmt.Println("Пакет calculator инициализирован!")
}