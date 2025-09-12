package examples

import (
    "fmt"
    "sync"
    "time"
)

// Фабрика функций
func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// Функция с отложенным выполнением
func deferredExecution() {
    fmt.Println("Начало выполнения...")
    defer fmt.Println("Завершение выполнения (отложенная функция)")

    fmt.Println("Основная логика...")
}

// Работа с замыканиями
func closureExample() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func DemoAdvancedPatterns() {
    fmt.Println("🎩 Продвинутые паттерны в Go")
    fmt.Println("---------------------------")

    // Фабрика функций
    double := createMultiplier(2)
    triple := createMultiplier(3)

    fmt.Printf("Удвоение 5: %d\n", double(5))
    fmt.Printf("Утроение 5: %d\n", triple(5))

    // Отложенное выполнение
    deferredExecution()

    // Замыкания
    counter := closureExample()
    fmt.Println("Счетчик замыканий:")
    for i := 0; i < 3; i++ {
        fmt.Printf("  %d\n", counter())
    }

    // Синглтон с once
    var once sync.Once
    var instance *string

    getInstance := func() *string {
        once.Do(func() {
            s := "единственный экземпляр"
            instance = &s
            fmt.Println("Создан синглтон")
        })
        return instance
    }

    for i := 0; i < 3; i++ {
        go func() {
            inst := getInstance()
            fmt.Printf("Получен экземпляр: %s\n", *inst)
        }()
    }

    time.Sleep(100 * time.Millisecond)
    fmt.Println()
}
