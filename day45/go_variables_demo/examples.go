package main

import "fmt"

// Демонстрация 6 способов объявления переменных в Go

// 1. Объявление на уровне пакета
var packageLevelVar string = "Переменная на уровне пакета"

// 2. Объявление блока переменных
var (
    globalName    string = "Алексей"
    globalAge     int    = 30
    globalIsAdmin bool   = true
)

// DemonstrateVarDeclaration демонстрирует различные способы объявления переменных
func DemonstrateVarDeclaration() {
    fmt.Println("🎯 6 способов объявления переменных в Go")
    fmt.Println("==========================================")

    // 1. Полное объявление с указанием типа
    var name string
    name = "Иван"
    fmt.Printf("1. Полное объявление: var name string = %q\n", name)

    // 2. Объявление с инициализацией
    var age int = 25
    fmt.Printf("2. Объявление с инициализацией: var age int = %d\n", age)

    // 3. Сокращенное объявление (тип определяется автоматически)
    country := "Россия"
    fmt.Printf("3. Сокращенное объявление: country := %q\n", country)

    // 4. Объявление нескольких переменных одного типа
    var x, y, z int
    x, y, z = 10, 20, 30
    fmt.Printf("4. Несколько переменных: x=%d, y=%d, z=%d\n", x, y, z)

    // 5. Объявление нескольких переменных с инициализацией
    var a, b, c = 1, 2.5, "три"
    fmt.Printf("5. Несколько с инициализацией: a=%d, b=%.1f, c=%q\n", a, b, c)

    // 6. Сокращенное объявление нескольких переменных
    firstName, lastName := "Анна", "Петрова"
    fmt.Printf("6. Сокращенное для нескольких: %s %s\n", firstName, lastName)

    fmt.Println("\n📦 Глобальные переменные:")
    fmt.Printf("   packageLevelVar: %s\n", packageLevelVar)
    fmt.Printf("   globalName: %s, globalAge: %d, globalIsAdmin: %t\n", 
        globalName, globalAge, globalIsAdmin)
}

// DemonstrateZeroValues демонстрирует нулевые значения переменных
func DemonstrateZeroValues() {
    fmt.Println("\n🔍 Нулевые значения в Go:")
    
    var i int
    var f float64
    var b bool
    var s string
    var arr [3]int
    var sl []string
    
    fmt.Printf("   int: %d\n", i)
    fmt.Printf("   float64: %.1f\n", f)
    fmt.Printf("   bool: %t\n", b)
    fmt.Printf("   string: %q\n", s)
    fmt.Printf("   array: %v\n", arr)
    fmt.Printf("   slice: %v (nil=%t)\n", sl, sl == nil)
}

// DemonstrateTypeInference демонстрирует определение типа
func DemonstrateTypeInference() {
    fmt.Println("\n🧠 Определение типа (type inference):")
    
    // Go автоматически определяет тип
    value := 42          // int
    message := "Привет"  // string
    ratio := 3.14        // float64
    enabled := true      // bool
    
    fmt.Printf("   value := 42        → %T\n", value)
    fmt.Printf("   message := 'Привет' → %T\n", message)
    fmt.Printf("   ratio := 3.14      → %T\n", ratio)
    fmt.Printf("   enabled := true    → %T\n", enabled)
}

// DemonstrateRedeclaration демонстрирует переобъявление переменных
func DemonstrateRedeclaration() {
    fmt.Println("\n🔄 Переобъявление и тень переменных (shadowing):")
    
    x := 10
    fmt.Printf("   Вне блока: x = %d\n", x)
    
    {
        x := 20 // Создает новую переменную, затеняющую внешнюю
        fmt.Printf("   Внутри блока: x = %d\n", x)
    }
    
    fmt.Printf("   Снова вне блока: x = %d\n", x)
    
    // Переприсваивание (не переобъявление)
    x = 30
    fmt.Printf("   После присваивания: x = %d\n", x)
}