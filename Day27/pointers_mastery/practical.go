package main

import "fmt"

func demoPractical() {
    fmt.Println("\n2. ПРАКТИЧЕСКОЕ ПРИМЕНЕНИЕ")
    fmt.Println("_________________________")
    
    // Пример с числами
    x, y := 10, 20
    fmt.Printf("До swap: x=%d, y=%d\n", x, y)
    swap(&x, &y)
    fmt.Printf("После swap: x=%d, y=%d\n", x, y)
    
    // Пример со структурой
    user := User{"Анна", 25}
    fmt.Printf("\nДо изменения: %+v\n", user)
    updateUser(&user)
    fmt.Printf("После изменения: %+v\n", user)
    
    // Эффективность с большими данными
    bigData := make([]int, 1000000)
    processData(&bigData)
    fmt.Printf("\nРазмер bigData: %d элементов\n", len(bigData))
}

func swap(a, b *int) {
    *a, *b = *b, *a
}

type User struct {
    Name string
    Age  int
}

func updateUser(u *User) {
    u.Age++
    u.Name = "Обновленная " + u.Name
}

func processData(data *[]int) {
    // Работаем с данными без копирования
    (*data)[0] = 100
}