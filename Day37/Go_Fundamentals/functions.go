package main

import "fmt"

// Простая функция
func Add(a, b int) int {
    return a + b
}

// Функция с несколькими возвращаемыми значениями
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("деление на ноль")
    }
    return a / b, nil
}

// Функция с именованными возвращаемыми значениями
func Power(base, exponent int) (result int) {
    result = 1
    for i := 0; i < exponent; i++ {
        result *= base
    }
    return
}

// Функция с переменным числом аргументов
func Sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Рекурсивная функция
func Factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * Factorial(n-1)
}

// Демонстрация функций
func DemonstrateFunctions() {
    fmt.Println("=== Функции ===")
    
    fmt.Printf("Add(5, 3) = %d\n", Add(5, 3))
    
    result, err := Divide(10, 2)
    if err != nil {
        fmt.Printf("Ошибка: %v\n", err)
    } else {
        fmt.Printf("Divide(10, 2) = %.1f\n", result)
    }
    
    fmt.Printf("Power(2, 8) = %d\n", Power(2, 8))
    fmt.Printf("Sum(1, 2, 3, 4, 5) = %d\n", Sum(1, 2, 3, 4, 5))
    fmt.Printf("Factorial(5) = %d\n", Factorial(5))
    
    fmt.Println()
}