#!/bin/bash

# Скрипт сборки OS Detector для Windows
# Запускается из WSL2 для сборки Windows-версии

set -e  # Прерывать выполнение при ошибках

echo "🪟 Сборка OS Detector для Windows..."
echo "📂 Текущая директория: $(pwd)"
echo "🔄 Версия Go: $(go version)"

# Создаем папку для бинарников, если её нет
mkdir -p bin

# Очистка предыдущих сборок Windows
echo "🧹 Очистка предыдущих сборок Windows..."
rm -f os-detector.exe bin/os-detector-windows-*

# Сборка для Windows (64-bit)
echo "🔨 Сборка для Windows AMD64..."
GOOS=windows GOARCH=amd64 go build -o bin/os-detector-windows-amd64.exe

# Сборка для Windows (32-bit)
echo "🔨 Сборка для Windows 386..."
GOOS=windows GOARCH=386 go build -o bin/os-detector-windows-386.exe

# Сборка для Windows ARM64
echo "🔨 Сборка для Windows ARM64..."
GOOS=windows GOARCH=arm64 go build -o bin/os-detector-windows-arm64.exe

echo "✅ Сборка для Windows завершена!"
echo "📦 Собранные файлы:"
ls -la bin/os-detector-windows-*

echo ""
echo "💡 Для копирования файлов в Windows используйте:"
echo "   cp bin/os-detector-windows-amd64.exe /mnt/c/Users/ВашеИмя/Desktop/"
echo ""
echo "🚀 Для запуска в Windows откройте командную строку и выполните:"
echo "   os-detector-windows-amd64.exe"