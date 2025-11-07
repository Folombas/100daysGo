package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("🚀 Демонстрация команды 'go' - вашего швейцарского ножа в Go!")
	fmt.Println("==============================================")

	// Демонстрация различных подкоманд
	commands := []struct{
		name string
		desc string
		cmd  string
	}{
		{"go build", "Компиляция программы", "go build -o demo-app"},
		{"go run", "Запуск без компиляции", "go run main.go helper.go"},
		{"go test", "Запуск тестов", "go test -v"},
		{"go fmt", "Форматирование кода", "go fmt ."},
		{"go mod", "Управление зависимостями", "go mod tidy"},
		{"go vet", "Статический анализ", "go vet ."},
	}

	for i, item := range commands {
		fmt.Printf("\n%d. %s: %s\n", i+1, item.name, item.desc)
	}

	fmt.Println("\n🎯 Выберите команду для демонстрации (1-6) или 0 для выхода:")

	var choice int
	fmt.Scan(&choice)

	if choice > 0 && choice <= len(commands) {
		demoCommand(commands[choice-1])
	} else {
		fmt.Println("👋 До свидания! Продолжайте изучать Go!")
	}
}

func demoCommand(cmdInfo struct{name string; desc string; cmd string}) {
	fmt.Printf("\n🎬 Демонстрация: %s\n", cmdInfo.name)
	fmt.Printf("💡 Команда: %s\n", cmdInfo.cmd)

	cmd := exec.Command("bash", "-c", cmdInfo.cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("⚠️ Ошибка: %v\n", err)
	} else {
		fmt.Printf("✅ %s выполнена успешно!\n", cmdInfo.name)
	}
}
