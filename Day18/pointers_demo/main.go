package main

import "fmt"

func main() {
    fmt.Println("Демонстрация указателей в Go")
    fmt.Println("============================")
    
    // Базовый пример
    var num int = 42
    ptr := &num
    
    fmt.Printf("\n1. Базовый пример:\nЗначение: %d\nАдрес: %p\nЧерез указатель: %d\n", 
        num, ptr, *ptr)
    
    // Демонстрации из других файлов
    demoSwap()
    demoStruct()
    demoMethods()
}