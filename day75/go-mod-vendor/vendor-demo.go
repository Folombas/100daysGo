package main

import (
    "fmt"
     "github.com/fatih/color"
)

// VendorDemo демонстрирует использование зависимостей из vendor
func VendorDemo() {
    fmt.Println("\n🎬 Демонстрация vendor в действии:")

    // Используем библиотеку из vendor
    red := color.New(color.FgRed, color.Bold).SprintFunc()
    green := color.New(color.FgGreen, color.Bold).SprintFunc()
    yellow := color.New(color.FgYellow).SprintFunc()

    fmt.Println(red("   Внимание!") + " Без vendor:")
    fmt.Println("   │ При сборке: 'go get' скачивает зависимости из интернета")
    fmt.Println("   │ Риск: репозитории могут исчезнуть, измениться")
    fmt.Println("   │ Проблема: нет интернета = нет сборки")

    fmt.Println(green("   С vendor:"))
    fmt.Println("   │ Все зависимости в папке vendor/")
    fmt.Println("   │ Сборка работает даже без интернета")
    fmt.Println("   │ Гарантия: те же версии, что и при разработке")

    fmt.Println("\n" + yellow("   Команды:"))
    fmt.Println("   go mod vendor      - создать/обновить vendor")
    fmt.Println("   go build -mod=vendor - собрать с использованием vendor")
    fmt.Println("   go mod tidy        - очистить ненужные зависимости")
}
