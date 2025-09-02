package main

import (
    "os"
    "os-detector/utils"
)

func main() {
    if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
        printUsage()
        return
    }
    
    // Всегда запускаем детектор ОС, если не запрошена помощь
    RunOSDetector()
}

func printUsage() {
    utils.PrintCyrillic("🖥️  OS Detector - Утилита определения операционной системы")
    utils.PrintCyrillic("")
    utils.PrintCyrillic("Использование: os-detector")
    utils.PrintCyrillic("")
    utils.PrintCyrillic("Опции:")
    utils.PrintCyrillic("  -h, --help    Показать эту справку")
    utils.PrintCyrillic("")
    utils.PrintCyrillic("Утилита определяет и отображает информацию об операционной системе,")
    utils.PrintCyrillic("на которой она запущена, включая тип ОС и версию дистрибутива Linux.")
}