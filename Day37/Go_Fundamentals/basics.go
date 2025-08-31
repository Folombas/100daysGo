package main

import "fmt"

// Демонстрация базовых концепций Go
func DemonstrateBasics() {
    fmt.Println("=== Базовые концепции Go ===")
    
    // Объявление переменных
    var explicitType string = "явный тип"
    implicitType := "неявный тип"
    fmt.Printf("Переменные: %s, %s\n", explicitType, implicitType)
    
    // Константы
    const pi = 3.14159
    fmt.Printf("Константа π: %.5f\n", pi)
    
    // Нулевые значения
    var zeroInt int
    var zeroString string
    var zeroBool bool
    fmt.Printf("Нулевые значения: int=%d, string='%s', bool=%t\n", 
        zeroInt, zeroString, zeroBool)
    
    // Преобразование типов
    var x int32 = 100
    var y int64 = int64(x)
    fmt.Printf("Преобразование типов: int32(%d) → int64(%d)\n", x, y)
    
    fmt.Println()
}
