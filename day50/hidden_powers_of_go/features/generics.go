package features

import (
    "fmt"
)


// Обобщенная структура
type Pair[K, V any] struct {
    Key   K
    Value V
}

func (p Pair[K, V]) String() string {
    return fmt.Sprintf("(%v: %v)", p.Key, p.Value)
}

func DemoGenerics() {
    fmt.Println("Генерики в Go (наконец-то!)")
    fmt.Println("-----------------------------")


// Обобщенные структуры
    pairs := []Pair[string, int]{
        {"яблоки", 5},
        {"бананы", 3},
        {"апельсины", 7},
    }

    fmt.Println("Пары ключ-значение:")
    for _, pair := range pairs {
        fmt.Printf("  - %s\n", pair)
    }

    fmt.Println()
}
