#!/bin/bash

# Скрипт сборки OS Detector для Linux
# Разместите этот файл в корне проекта и сделайте исполняемым: chmod +x build-linux.sh

set -e  # Прерывать выполнение при ошибках

echo "🐧 Сборка OS Detector для Linux..."
echo "📂 Текущая директория: $(pwd)"
echo "🔄 Версия Go: $(go version)"

# Создаем папку для бинарников, если её нет
mkdir -p bin

# Очистка предыдущих сборок
echo "🧹 Очистка предыдущих сборок..."
rm -f os-detector os-detector-* bin/os-detector-*

# Сборка для текущей архитектуры
echo "🔨 Сборка для $(uname -m)..."
go build -o bin/os-detector

# Сборка для других архитектур Linux
echo "🔨 Сборка для AMD64..."
GOOS=linux GOARCH=amd64 go build -o bin/os-detector-amd64

echo "🔨 Сборка для ARM64..."
GOOS=linux GOARCH=arm64 go build -o bin/os-detector-arm64

echo "🔨 Сборка для 386..."
GOOS=linux GOARCH=386 go build -o bin/os-detector-386

# Делаем файлы исполняемыми
chmod +x bin/os-detector*

echo "✅ Сборка завершена!"
echo "📦 Собранные файлы:"
ls -la bin/os-detector*

echo ""
echo "🚀 Для запуска выполните:"
echo "   ./bin/os-detector"