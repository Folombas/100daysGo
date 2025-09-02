@echo off
chcp 65001 > nul
echo 🚀 Запуск OS Detector для Windows...

:: Проверяем, собран ли файл
if not exist bin\os-detector-windows-amd64.exe (
    echo ❌ Файл не собран. Сначала выполните build-windows.bat
    pause
    exit /b 1
)

:: Запускаем программу
echo 📦 Запуск программы...
bin\os-detector-windows-amd64.exe

pause