#!/bin/bash

echo "🔨 Кросс-платформенная компиляция Go"
echo "====================================="

# Создаем папку для бинарных файлов
mkdir -p builds

# Компилируем для различных платформ
echo "Компиляция для Linux x64..."
GOOS=linux GOARCH=amd64 go build -o builds/app_linux main.go

echo "Компиляция для Windows x64..."
GOOS=windows GOARCH=amd64 go build -o builds/app_win.exe main.go

echo "Компиляция для macOS x64..."
GOOS=darwin GOARCH=amd64 go build -o builds/app_macos main.go

echo "Компиляция для Linux ARM..."
GOOS=linux GOARCH=arm go build -o builds/app_linux_arm main.go

echo "✅ Все бинарные файлы созданы в папке builds/"
echo "Размеры файлов:"
ls -lh builds/