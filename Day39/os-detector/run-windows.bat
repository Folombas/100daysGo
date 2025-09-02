@echo off
chcp 65001 > nul
echo 🚀 Запуск OS Detector для Windows...

if not exist bin\os-detector.exe (
    echo ❌ Файл не собран. Сначала выполните build-windows.bat
    pause
    exit /b 1
)

echo 📦 Запуск программы...
bin\os-detector.exe

pause