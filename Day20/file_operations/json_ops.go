package main

import (
    "encoding/json"
    "os"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

// Сохранить структуру в JSON
func saveToJSON(filename string, data interface{}) error {
    file, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename, file, 0644)
}

// Прочитать JSON в структуру
func loadFromJSON(filename string, target interface{}) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, target)
}

// Демо: работа с JSON
func runJSONDemo() {
    user := User{"Гоша Гофер", "gosha@gofer.com", 30}
    saveToJSON("user.json", user)
    
    var newUser User
    loadFromJSON("user.json", &newUser)
    println("\nJSON данные:\nИмя:", newUser.Name, "\nEmail:", newUser.Email)
}