package main

import (
    "fmt"
)

func main() {
    fmt.Println("День 36: Односвязные списки в Go")
    fmt.Println("=================================")
    
    // Создаем новый список
    list := NewSinglyLinkedList()
    
    // Добавляем элементы в список
    list.AddToEnd("Первый")
    list.AddToEnd("Второй")
    list.AddToFront("Новый первый")
    list.AddToEnd("Третий")
    
    // Выводим список
    fmt.Printf("Список после добавления элементов: %s\n", list.String())
    fmt.Printf("Размер списка: %d\n", list.GetSize())
    
    // Поиск элемента
    if node, found := list.Find("Второй"); found {
        fmt.Printf("Найден элемент: %v\n", node.Data)
    }
    
    // Реверс списка
    list.Reverse()
    fmt.Printf("Список после реверса: %s\n", list.String())
    
    // Удаление элементов
    if data, ok := list.RemoveFromFront(); ok {
        fmt.Printf("Удален первый элемент: %v\n", data)
    }
    
    fmt.Printf("Финальный список: %s\n", list.String())
    fmt.Printf("Финальный размер: %d\n", list.GetSize())
    
    // Демонстрация с числами
    fmt.Println("\n--- Работа со списком чисел ---")
    numList := NewSinglyLinkedList()
    for i := 1; i <= 5; i++ {
        numList.AddToEnd(i * 10)
    }
    
    fmt.Printf("Числовой список: %s\n", numList.String())
    numList.Reverse()
    fmt.Printf("После реверса: %s\n", numList.String())
}