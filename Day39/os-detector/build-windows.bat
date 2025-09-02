@echo off
chcp 65001 > nul
echo ⚙ Сборка OS-Detector для Windows...
echo.

echo 🔍 Проверка установки Go...
go version
if errorlevel 1 (
    echo ❌ Ошибка: Go не установлен или не настроен
    pause
    exit /b 1
)

echo.
echo 📦 Подготовка папки для бинарников...
if not exist bin mkdir bin

echo.
echo 🔨 Сборка программы...
go build -o bin\os-detector.exe

echo.
echo ✅ Сборка завершена!
echo 🚀 Для запуска введите:
echo    .\bin\os-detector.exe
echo.
echo 📁 Файлы собраны в папку: %CD%\bin
echo.
pause