package main

import "fmt"

func main() {
    fmt.Println("🐹 Мастерство указателей в Go")
    fmt.Println("=============================")
    
    demoBasics()
    demoPractical()
    demoAdvanced()
    
    fmt.Println("\n🎯 Запомни: Указатели - это просто адреса в памяти!")
    fmt.Println("   * - получить значение по адресу")
    fmt.Println("   & - получить адрес переменной")
}