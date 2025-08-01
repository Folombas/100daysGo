package main

import (
    "fmt"
    "log"
)

func main() {
    // Демонстрация работы калькулятора
    fmt.Println("Калькулятор:")
    fmt.Printf("5 + 3 = %d\n", Add(5, 3))
    fmt.Printf("10 - 7 = %d\n", Subtract(10, 7))
    fmt.Printf("4 * 6 = %d\n", Multiply(4, 6))
    
    if res, ok := Divide(8, 4); ok {
        fmt.Printf("8 / 4 = %.1f\n", res)
    } else {
        fmt.Println("Деление на ноль!")
    }

    // Проверка простых чисел
    fmt.Println("\nПроверка простых чисел:")
    primes := []int{2, 3, 4, 5, 29, 30}
    for _, n := range primes {
        fmt.Printf("%d - простoe: %t\n", n, IsPrime(n))
    }

    // Генерация отчета о покрытии
    fmt.Println("\nГенерация отчета о покрытии...")
    if err := GenerateCoverageReport(); err != nil {
        log.Fatalf("Ошибка генерации отчета: %v", err)
    }
    fmt.Println("Отчет coverage.html успешно создан!")
}