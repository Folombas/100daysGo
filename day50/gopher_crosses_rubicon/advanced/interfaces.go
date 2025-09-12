package advanced

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

type Stringer interface {
    String() string
}

func (c Circle) String() string {
    return fmt.Sprintf("Круг радиусом %.2f", c.Radius)
}

func DemoInterfaces() {
    fmt.Println("\n2. Продвинутые интерфейсы")
    fmt.Println("------------------------")

    // Пустой интерфейс и type assertion
    var i interface{} = Circle{Radius: 5}

    if s, ok := i.(Shape); ok {
        fmt.Printf("Площадь: %.2f\n", s.Area())
        fmt.Printf("Периметр: %.2f\n", s.Perimeter())
    }

    if s, ok := i.(Stringer); ok {
        fmt.Println("Строковое представление:", s.String())
    }

    // Type switch
    switch v := i.(type) {
    case Shape:
        fmt.Printf("Это фигура с площадью %.2f\n", v.Area())
    case Stringer:
        fmt.Println("Это объект с строковым представлением:", v.String())
    default:
        fmt.Println("Неизвестный тип")
    }

    // Вложенные интерфейсы
    var complex interface{} = Circle{Radius: 7}
    if shape, ok := complex.(Shape); ok {
        if stringer, ok := complex.(Stringer); ok {
            fmt.Printf("Комплексный объект: %s, площадь=%.2f\n",
                stringer.String(), shape.Area())
        }
    }
}
