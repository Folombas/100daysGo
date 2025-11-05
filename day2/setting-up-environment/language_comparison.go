package main

import "fmt"

func CompareWithOtherLanguages() {
    fmt.Println("\n🆚 Отличия Go от других языков:")
    fmt.Println("-------------------------------")

    features := []struct{
        Feature string
        Go      string
        Others  string
    }{
        {"Сборка мусора", "✅ Встроенный GC", "❌ C/C++: ручное управление"},
        {"Параллельность", "✅ Горутины + каналы", "⚠️ Python/JS: GIL, callback hell"},
        {"Зависимости", "✅ Go Modules", "⚠️ Python: pip, Node: npm"},
        {"Компиляция", "✅ Один бинарный файл", "⚠️ Python/JS: интерпретатор"},
        {"Типизация", "✅ Статическая + строгая", "⚠️ Python/JS: динамическая"},
    }

    for _, f := range features {
        fmt.Printf("\n%s:\n", f.Feature)
        fmt.Printf("   Go: %s\n", f.Go)
        fmt.Printf("   Другие: %s\n", f.Others)
    }
}
