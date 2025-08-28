package main

import (
    "fmt"
    "log"
    
    "error_handling/handlers"
)

func main() {
    fmt.Println("🛡️ Веб-демонстрация обработки ошибок в Go")
    fmt.Println("========================================")
    
    // Инициализация дополнительных компонентов если нужно
    if err := handlers.StartWebServer(); err != nil {
        log.Fatalf("Ошибка запуска веб-сервера: %v", err)
    }
}