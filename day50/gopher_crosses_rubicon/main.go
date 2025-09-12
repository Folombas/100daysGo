package main

import (
    "fmt"
    "runtime"
		"gopher_crosses_rubicon/advanced"
    "gopher_crosses_rubicon/internals"
    "gopher_crosses_rubicon/performance"
)

func main() {
    fmt.Println("🎉 50 дней айти-марафона: Гофер переходит Рубикон")
    fmt.Println("==================================================")
    fmt.Println("Изучаем подкапотные тонкости языка Go")
    fmt.Printf("Версия Go: %s\n", runtime.Version())
    fmt.Printf("ОС: %s, Архитектура: %s\n\n", runtime.GOOS, runtime.GOARCH)

    // Демонстрация продвинутых возможностей
    fmt.Println("🚀 ПРОДВИНУТЫЕ ВОЗМОЖНОСТИ")
    fmt.Println("==========================")
    advanced.DemoChannels()
    advanced.DemoInterfaces()
    advanced.DemoReflection()

    // Демонстрация внутренних механизмов
    fmt.Println("🔧 ВНУТРЕННИЕ МЕХАНИЗМЫ")
    fmt.Println("=======================")
    internals.DemoMemoryManagement()
    internals.DemoScheduler()

    // Демонстрация оптимизаций
    fmt.Println("⚡ ОПТИМИЗАЦИЯ ПРОИЗВОДИТЕЛЬНОСТИ")
    fmt.Println("================================")
    performance.DemoOptimizations()

    fmt.Println("\n🎯 Гофер успешно перешел Рубикон и готов к новым вызовам!")
    fmt.Println("Следующие 50 дней будут еще более захватывающими!")
}
