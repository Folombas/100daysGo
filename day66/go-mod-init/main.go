package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("🚀 GO MODULES & DEPENDENCIES DEMO")
	fmt.Println("=================================")

	// Текущая директория
	wd, _ := os.Getwd()
	fmt.Printf("📁 Текущая директория: %s\n\n", wd)

	// 1. Демонстрация создания модуля
	fmt.Println("🎯 1. СОЗДАНИЕ МОДУЛЯ:")
	fmt.Println("   go mod init go-mod-init")
	fmt.Println("   -> Создает go.mod файл")

	// 2. Основные команды go mod
	fmt.Println("🛠️  2. ОСНОВНЫЕ КОМАНДЫ:")
	commands := []struct {
		cmd string
		desc string
	}{
		{"go mod init", "Создает новый модуль"},
		{"go mod tidy", "Очищает зависимости"},
		{"go mod download", "Скачивает зависимости"},
		{"go mod vendor", "Создает vendor папку"},
		{"go mod graph", "Показывает граф зависимостей"},
		{"go mod why", "Объясняет зачем зависимость"},
	}

	for _, cmd := range commands {
		fmt.Printf("   %s - %s\n", cmd.cmd, cmd.desc)
	}
	fmt.Println()

	// 3. Пример go.mod файла
	fmt.Println("📄 3. ПРИМЕР GO.MOD ФАЙЛА:")
	goModExample := `module go-mod-init

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/sirupsen/logrus v1.9.0
)

replace example.com/local => ../local-package`
	fmt.Println(goModExample)
	fmt.Println()

	// 4. Работа с зависимостями
	fmt.Println("📦 4. РАБОТА С ЗАВИСИМОСТЯМИ:")
	fmt.Println("   go get github.com/package/name@v1.2.3")
	fmt.Println("   go get -u github.com/package/name")
	fmt.Println("   go mod tidy")
	fmt.Println()

	// 5. Практическая демонстрация
	fmt.Println("🔧 5. ПРАКТИЧЕСКАЯ ДЕМОНСТРАЦИЯ:")

	// Проверяем наличие go.mod
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println("   ❌ go.mod не найден - создаем...")
		// Создаем go.mod
		createGoMod()
	} else {
		fmt.Println("   ✅ go.mod найден")
	}

	// Показываем содержимое go.mod
	showGoMod()

	fmt.Println("\n🎯 ВЫВОДЫ:")
	fmt.Println("✅ Go modules - стандартная система управления зависимостями")
	fmt.Println("✅ go.mod - файл описания модуля")
	fmt.Println("✅ go.sum - файл с хешами для безопасности")
	fmt.Println("✅ Используйте go mod tidy для очистки зависимостей")
	fmt.Println("✅ Используйте семантическое версионирование")
}

func createGoMod() {
	content := `module example.com/demo

go 1.21

require (
    golang.org/x/example v0.0.0-20210811190340-787a929d5a0d
)

replace example.com/local => ./local
`

	err := os.WriteFile("go.mod", []byte(content), 0644)
	if err != nil {
		fmt.Printf("   ❌ Ошибка создания go.mod: %v\n", err)
		return
	}
	fmt.Println("   ✅ go.mod успешно создан")
}

func showGoMod() {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		fmt.Printf("   ❌ Ошибка чтения go.mod: %v\n", err)
		return
	}

	fmt.Println("   📋 Содержимое go.mod:")
	fmt.Printf("%s\n", string(data))
}
