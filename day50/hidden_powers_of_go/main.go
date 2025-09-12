package main

import (
    "fmt"
    "hidden_powers_of_go/features"
    "hidden_powers_of_go/examples"
    "runtime"
)

func main() {
    fmt.Println("🎉 50 оттенков Go: скрытое богатство за минимализмом")
    fmt.Println("=====================================================")

    // Демонстрация возможностей
    fmt.Printf("Версия Go: %s\n", runtime.Version())
    fmt.Printf("ОС: %s, Архитектура: %s\n\n", runtime.GOOS, runtime.GOARCH)

    features.DemoConcurrency()
    features.DemoInterfaces()
    features.DemoReflection()
    features.DemoGenerics()
    features.DemoErrorHandling()

    examples.DemoAdvancedPatterns()

    fmt.Println("\n✨ И это лишь малая часть возможностей Go!")
}
