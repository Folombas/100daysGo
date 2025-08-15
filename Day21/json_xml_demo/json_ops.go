package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID      	int      `json:"user_id"`
	Name    	string   `json:"name"`
	Email   	string   `json:"email,omitempty"`
	Password	string   `json:"-"`
}

// Сериализация в JSON
func toJSON(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// Десериализация из JSON
func fromJSON(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// JSON демо
func jsonDemo() {
    // Создаём объект
    user := User{1, "Гоша Гошник", "gosha@golang_programmer.ru", "secret"}
    
    // Сериализуем
    jsonData, _ := toJSON(user)
    writeFile("user.json", jsonData)
    
    // Читаем и десериализуем
    data, _ := readFile("user.json")
    var newUser User
    fromJSON(data, &newUser)
    
    fmt.Println("JSON Demo:")
    fmt.Println("Исходный:", user)
    fmt.Println("Из файла:", newUser)
    fmt.Println("Сырой JSON:", string(data))
}