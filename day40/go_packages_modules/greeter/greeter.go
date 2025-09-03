package greeter

import (
    "fmt"
    "time"
)

// Greet возвращает приветствие с текущим временем
func Greet(name string) string {
    currentTime := time.Now().Format("15:04:05")
    return fmt.Sprintf("Привет, %s! Сейчас %s.", name, currentTime)
}

// GreetWithColor возвращает цветное приветствие
func GreetWithColor(name, color string) string {
    return fmt.Sprintf("[%s]Добро пожаловать, %s![/%s]", color, name, color)
}