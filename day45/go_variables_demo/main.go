package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("🚀 Day 45: 6 способов объявления переменных в Go")
    fmt.Println("================================================")
    
    // Демонстрация различных способов объявления переменных
    DemonstrateVarDeclaration()
    
    // Демонстрация нулевых значений
    DemonstrateZeroValues()
    
    // Демонстрация определения типа
    DemonstrateTypeInference()
    
    // Демонстрация переобъявления и затенения
    DemonstrateRedeclaration()
    
    // Дополнительные примеры
    fmt.Println("\n💡 Дополнительные примеры:")
    demonstrateAdditionalExamples()
    
    fmt.Println("\n🎉 Изучение способов объявления переменных завершено!")
    fmt.Println("Теперь вы знаете все тонкости работы с переменными в Go!")
}

// demonstrateAdditionalExamples показывает дополнительные примеры
func demonstrateAdditionalExamples() {
    // Объявление с использованием функций
    length := len("Привет")
    fmt.Printf("   Длина строки: %d\n", length)
    
    // Множественное присваивание
    x, y := 10, 20
    x, y = y, x // Обмен значений
    fmt.Printf("   После обмена: x=%d, y=%d\n", x, y)
    
    // Использование _ для игнорирования значений
    firstName, _, lastName := getPerson()
    fmt.Printf("   Имя: %s, Фамилия: %s\n", firstName, lastName)
    
    // Короткое объявление в условии
    if temperature := getTemperature(); temperature > 25 {
        fmt.Printf("   Температура %d°C - жарко!\n", temperature)
    } else {
        fmt.Printf("   Температура %d°C - прохладно!\n", temperature)
    }
}

// getTemperature возвращает случайную температуру
func getTemperature() int {
    return 15 + int(time.Now().UnixNano()%25)
}

// getPerson возвращает данные о человеке
func getPerson() (string, int, string) {
    return "Мария", 28, "Иванова"
}