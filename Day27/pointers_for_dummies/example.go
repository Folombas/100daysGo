package main

import "fmt"

func demoExample() {
    fmt.Println("\n2. ПРАКТИЧЕСКИЙ ПРИМЕР")
    fmt.Println("_____________________")
    
    // 1. Создаем переменную (дом)
    message := "Я люблю Go!"
    fmt.Printf("1. message = %s\n", message)
    
    // 2. Создаем указатель (адрес дома)
    pointer := &message
    fmt.Printf("2. pointer = %p (адрес message)\n", pointer)
    
    // 3. Смотрим что внутри указателя
    fmt.Printf("3. *pointer = %s (значение по адресу)\n", *pointer)
    
    // 4. Меняем значение через указатель
    *pointer = "Я понимаю указатели!"
    fmt.Printf("4. message = %s (изменилось через указатель)\n", message)
    
    // 5. Сравниваем адреса
    fmt.Printf("5. &message = %p (адрес оригинала)\n", &message)
    fmt.Printf("6. pointer  = %p (тот же адрес)\n", pointer)
}
