package main

import "fmt"

func main() {
    fmt.Println("Демонстрация тестирования в Go")
    fmt.Println("==============================")
    
    // Простой пример использования калькулятора
    calc := Calculator{}
    result := calc.Add(10, 5)
    fmt.Printf("10 + 5 = %d\n", result)
    
    result = calc.Subtract(10, 5)
    fmt.Printf("10 - 5 = %d\n", result)
}