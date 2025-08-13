package main

import "fmt"

func main() {
    // Инициализация структур
    ivan := NewPerson("Гоша", 37)
    emp := NewEmployee("Ася", 23, 150000.50, "Яндекс")

    // Демонстрация методов
    PrintToStdout(
        ivan.Introduce(),
        emp.Introduce(),
        emp.WorkInfo(),
    )

    // Полиморфизм через интерфейсы
    var speaker Speaker
    speaker = ivan
    PrintToStdout(speaker.Speak())

    speaker = Animal{"Собака"}
    PrintToStdout(MakeSound(speaker))

    // Изменение состояния
    ivan.Birthday()
    fmt.Printf("После дня рождения: %s\n", ivan.Introduce())
}