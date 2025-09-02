package utils

import (
    "fmt"
    "os"
)

// EnsureUnicodeSupport обеспечивает поддержку Unicode
func EnsureUnicodeSupport() {
    // Устанавливаем переменную окружения LANG, если не установлена
    if os.Getenv("LANG") == "" {
        os.Setenv("LANG", "en_US.UTF-8")
    }
}

// PrintCyrillic печатает кириллический текст с проверкой поддержки
func PrintCyrillic(text string) {
    EnsureUnicodeSupport()
    fmt.Println(text)
}

// PrintBoldCyrillic печатает жирный кириллический текст
func PrintBoldCyrillic(text string) {
    EnsureUnicodeSupport()
    fmt.Printf("\x1b[1m%s\x1b[0m\n", text) // Жирный текст
}

// PrintColorCyrillic печатает цветной кириллический текст
func PrintColorCyrillic(text string, colorCode string) {
    EnsureUnicodeSupport()
    fmt.Printf("%s%s\x1b[0m\n", colorCode, text)
}