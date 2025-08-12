package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func updateUser(user *User) {
    user.Name = "Гоша Гофер"
    user.Age = 37
}

func demoStruct() {
    u := User{"Женя", 30}
    fmt.Printf("\n3. Структура до:\n%+v\n", u)
    
    updateUser(&u)
    fmt.Printf("Структура после:\n%+v\n", u)
}