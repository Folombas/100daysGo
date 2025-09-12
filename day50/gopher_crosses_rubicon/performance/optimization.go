package performance

import (
    "fmt"
    "time"
		"strings"
)

func DemoOptimizations() {
    fmt.Println("\n1. Методы оптимизации в Go")
    fmt.Println("--------------------------")

    // Бенчмарк разных подходов
    fmt.Println("Сравнение производительности:")

    // Тест 1: Аллокация памяти
    testAllocation()

    // Тест 2: Конкатенация строк
    testStringConcatenation()

    // Тест 3: Использование интерфейсов
    testInterfacePerformance()

    // Профилирование
    fmt.Println("\nСоветы по оптимизации:")
    fmt.Println("  - Используйте sync.Pool для частых аллокаций")
    fmt.Println("  - Избегайте лишних преобразований типов")
    fmt.Println("  - Используйте буферизированные каналы при необходимости")
    fmt.Println("  - Минимизируйте использование рефлексии")
    fmt.Println("  - Используйте strings.Builder для конкатенации строк")
}

func testAllocation() {
    start := time.Now()
    for i := 0; i < 100000; i++ {
        _ = make([]byte, 1024)
    }
    elapsed1 := time.Since(start)

    start = time.Now()
    pool := make([][]byte, 100000)
    for i := 0; i < 100000; i++ {
        pool[i] = make([]byte, 1024)
    }
    elapsed2 := time.Since(start)

    fmt.Printf("  Аллокация: без пула=%v, с предварительным выделением=%v\n",
        elapsed1, elapsed2)
}

func testStringConcatenation() {
    start := time.Now()
    s := ""
    for i := 0; i < 10000; i++ {
        s += "a"
    }
    elapsed1 := time.Since(start)

    start = time.Now()
    var builder strings.Builder
    for i := 0; i < 10000; i++ {
        builder.WriteString("a")
    }
    s2 := builder.String()
    elapsed2 := time.Since(start)

    fmt.Printf("  Конкатенация строк: оператор +=%v, Builder=%v\n",
        elapsed1, elapsed2)
    _ = s2 // Используем переменную
}

func testInterfacePerformance() {
    start := time.Now()
    var val interface{}
    for i := 0; i < 1000000; i++ {
        val = i
        _ = val.(int)
    }
    elapsed1 := time.Since(start)

    start = time.Now()
    for i := 0; i < 1000000; i++ {
        val := i
        _ = val
    }
    elapsed2 := time.Since(start)

    fmt.Printf("  Интерфейсы: с интерфейсом=%v, без интерфейса=%v\n",
        elapsed1, elapsed2)
}
