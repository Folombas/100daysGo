@echo off
chcp 65001 > nul
echo 🪟 Сборка OS Detector для Windows...
echo 📂 Текущая директория: %CD%
echo 🔄 Проверка установки Go...

:: Проверяем, установлен ли Go
go version >nul 2>&1
if errorlevel 1 (
    echo ❌ Go не установлен или не добавлен в PATH
    echo 💻 Скачайте и установите Go с https://golang.org/dl/
    pause
    exit /b 1
)

echo ✅ Версия Go: %GOVERSION%
for /f "tokens=*" %%i in ('go version') do set GOVERSION=%%i
echo ✅ Версия Go: %GOVERSION%

:: Создаем папку для бинарников, если её нет
if not exist bin mkdir bin

:: Очистка предыдущих сборок
echo 🧹 Очистка предыдущих сборок...
del /q bin\os-detector-*.exe 2>nul

:: Сборка для Windows (64-bit)
echo 🔨 Сборка для Windows AMD64...
go build -o bin\os-detector-windows-amd64.exe

:: Сборка для Windows (32-bit)
echo 🔨 Сборка для Windows 386...
set GOARCH=386
go build -o bin\os-detector-windows-386.exe

:: Возвращаем архитектуру по умолчанию
set GOARCH=amd64

echo ✅ Сборка для Windows завершена!
echo 📦 Собранные файлы:
dir /b bin\os-detector-*.exe

echo.
echo 🚀 Для запуска выполните:
echo    bin\os-detector-windows-amd64.exe
echo.
pause