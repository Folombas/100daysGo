package features

import (
    "fmt"
    "math"
)

// Базовый интерфейс
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Реализации интерфейса
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func DemoInterfaces() {
    fmt.Println("🎯 Мощь интерфейсов в Go")
    fmt.Println("------------------------")

    shapes := []Shape{
        Circle{Radius: 5},
        Rectangle{Width: 3, Height: 4},
    }

    for _, shape := range shapes {
        fmt.Printf("Фигура: %T\n", shape)
        fmt.Printf("  Площадь: %.2f\n", shape.Area())
        fmt.Printf("  Периметр: %.2f\n", shape.Perimeter())
    }

    // Пустой интерфейс и type assertion
    var anything interface{} = "Привет, Go!"
    fmt.Printf("\nПустой интерфейс: %v\n", anything)

    if str, ok := anything.(string); ok {
        fmt.Printf("Это строка: %s\n", str)
    }

    // Type switch
    switch v := anything.(type) {
    case string:
        fmt.Printf("Переменная типа string: %s\n", v)
    case int:
        fmt.Printf("Переменная типа int: %d\n", v)
    default:
        fmt.Printf("Неизвестный тип: %T\n", v)
    }

    fmt.Println()
}
