package main

import "fmt"

// Демонстрация управляющих конструкций
func DemonstrateControlStructures() {
    fmt.Println("=== Управляющие конструкции ===")
    
    // Условные операторы
    fmt.Println("--- Условные операторы ---")
    x := 10
    
    if x > 5 {
        fmt.Println("x больше 5")
    } else {
        fmt.Println("x меньше или равно 5")
    }
    
    // Короткая запись условия
    if y := 20; y > 15 {
        fmt.Println("y больше 15")
    }
    
    // Switch
    fmt.Println("--- Оператор switch ---")
    day := "среда"
    
    switch day {
    case "понедельник":
        fmt.Println("Начало недели")
    case "вторник", "среда", "четверг":
        fmt.Println("Середина недели")
    case "пятница":
        fmt.Println("Конец рабочей недели")
    default:
        fmt.Println("Выходной день")
    }
    
    // Циклы
    fmt.Println("--- Циклы ---")
    
    // Классический for
    fmt.Print("Классический for: ")
    for i := 0; i < 5; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    
    // While-подобный цикл
    fmt.Print("While-подобный: ")
    j := 0
    for j < 5 {
        fmt.Printf("%d ", j)
        j++
    }
    fmt.Println()
    
    // Бесконечный цикл с break
    fmt.Print("С break: ")
    k := 0
    for {
        if k >= 5 {
            break
        }
        fmt.Printf("%d ", k)
        k++
    }
    fmt.Println()
    
    // Range для итерации по коллекциям
    fmt.Println("Range по слайсу:")
    numbers := []int{10, 20, 30, 40, 50}
    for index, value := range numbers {
        fmt.Printf("numbers[%d] = %d\n", index, value)
    }
    
    fmt.Println("Range по карте:")
    fruits := map[string]int{"яблоко": 5, "апельсин": 3, "банан": 7}
    for key, value := range fruits {
        fmt.Printf("%s: %d\n", key, value)
    }
    
    fmt.Println()
}