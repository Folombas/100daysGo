package main

import "fmt"

func demoBasics() {
    fmt.Println("\n1. ОСНОВЫ УКАЗАТЕЛЕЙ")
    fmt.Println("___________________")
    
    // Обычная переменная
    message := "Привет из Python!"
    fmt.Printf("message = %s\n", message)
    fmt.Printf("Адрес message: %p\n", &message)
    
    // Создаем указатель
    var pointer *string = &message
    fmt.Printf("\npointer = %p\n", pointer)
    fmt.Printf("Значение через pointer: %s\n", *pointer)
    
    // Меняем значение через указатель
    *pointer = "Теперь я использую Go!"
    fmt.Printf("\nПосле изменения через pointer:\n")
    fmt.Printf("message = %s\n", message)
    fmt.Printf("*pointer = %s\n", *pointer)
    
    // Новая переменная - новый адрес
    newMessage := message
    fmt.Printf("\nАдрес newMessage: %p\n", &newMessage)
}