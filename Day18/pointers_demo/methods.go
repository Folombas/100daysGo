package main

import "fmt"

type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++
}

func (c Counter) CurrentValue() int {
    return c.count
}

func demoMethods() {
    c := Counter{}
    fmt.Println("\n4. Счетчик до:", c.CurrentValue())
    
    for i := 0; i < 3; i++ {
        c.Increment()
    }
    fmt.Println("Счетчик после:", c.CurrentValue())
    
    // Важное замечание
    fmt.Println("\n⚠️ Методы с pointer receiver позволяют:")
    fmt.Println("- Изменять состояние структуры")
    fmt.Println("- Работать с nil (предотвращать панику)")
}