#!/bin/bash

echo "🐧 Сборка OS Detector для Linux..."
echo "Текущая версия Go: $(go version)"

# Очистка предыдущих сборок
echo "🧹 Очистка предыдущих сборок..."
rm -f os-detector os-detector-*

# Сборка для текущей архитектуры
echo "🔨 Сборка для $(uname -m)..."
go build -o os-detector

# Сборка для других архитектур
echo "🔨 Сборка для AMD64..."
GOOS=linux GOARCH=amd64 go build -o os-detector-amd64

echo "🔨 Сборка для ARM64..."
GOOS=linux GOARCH=arm64 go build -o os-detector-arm64

# Делаем файлы исполняемыми
chmod +x os-detector os-detector-*

echo "✅ Сборка завершена!"
echo "📦 Собранные файлы:"
ls -la os-detector*