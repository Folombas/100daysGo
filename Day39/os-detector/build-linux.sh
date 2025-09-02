#!/bin/bash

echo "🐧 Сборка OS Detector для Linux (Ubuntu 24.04, Debian 12, Kali Linux)..."
echo "📂 Текущая директория: $(pwd)"
echo "🔄 Версия Go: $(go version)"

# Создаем папку для бинарников
mkdir -p bin

# Очистка предыдущих сборок
echo "🧹 Очистка предыдущих сборок..."
rm -f bin/os-detector-*

# Сборка для Ubuntu 24.04 (amd64)
echo "🔨 Сборка для Ubuntu 24.04 (amd64)..."
GOOS=linux GOARCH=amd64 go build -o bin/os-detector-ubuntu-24.04

# Сборка для Debian 12 (amd64)
echo "🔨 Сборка для Debian 12 (amd64)..."
GOOS=linux GOARCH=amd64 go build -o bin/os-detector-debian-12

# Сборка для Kali Linux (amd64)
echo "🔨 Сборка для Kali Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -o bin/os-detector-kali-linux

# Делаем файлы исполняемыми
chmod +x bin/os-detector*

echo "✅ Сборка завершена!"
echo "📦 Собранные файлы:"
ls -la bin/os-detector*

echo ""
echo "🚀 Для запуска выберите соответствующий бинарник:"
echo "   Для Ubuntu 24.04:   ./bin/os-detector-ubuntu-24.04"
echo "   Для Debian 12:      ./bin/os-detector-debian-12"
echo "   Для Kali Linux:     ./bin/os-detector-kali-linux"