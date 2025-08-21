package main

import "fmt"

func demoAdvanced() {
    fmt.Println("\n3. ПРОДВИНУТЫЕ КОНЦЕПЦИИ")
    fmt.Println("_______________________")
    
    // Указатель на указатель
    value := "hello"
    p1 := &value
    p2 := &p1
    
    fmt.Printf("value: %s\n", value)
    fmt.Printf("p1: %p -> %s\n", p1, *p1)
    fmt.Printf("p2: %p -> %p -> %s\n", p2, *p2, **p2)
    
    // Методы с pointer receiver
    counter := Counter{count: 0}
    counter.Increment()
    counter.Increment()
    fmt.Printf("\nСчетчик: %d\n", counter.Value())
    
    // Nil указатели
    var nilPointer *string
    fmt.Printf("\nNil pointer: %p\n", nilPointer)
    if nilPointer == nil {
        fmt.Println("Указатель не инициализирован!")
    }
    
    // Безопасное использование
    safeUsage(nilPointer)
}

type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++
}

func (c *Counter) Value() int {
    return c.count
}

func safeUsage(ptr *string) {
    if ptr != nil {
        fmt.Println("Значение:", *ptr)
    } else {
        fmt.Println("Указатель nil, пропускаем")
    }
}