package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// EnvironmentChecker проверяет и отображает настройки окружения Go
type EnvironmentChecker struct {
	Name     string
	Version  string
	IsReady  bool
	Problems []string
}

func main() {
	fmt.Println("🎮 День 93: 'Гоша в стране Go-чудес' 🎮")
	fmt.Println("========================================")

	checker := &EnvironmentChecker{
		Name:    "Гоша",
		Version: "1.21",
	}

	checker.WelcomeMessage()
	checker.CheckGoInstallation()
	checker.CheckEnvironmentVariables()
	checker.CheckProjectStructure()
	checker.CheckDevelopmentTools()
	checker.ShowResults()
	checker.MotivationalMessage()
}

func (e *EnvironmentChecker) WelcomeMessage() {
	fmt.Printf("\n👋 Привет, я %s!\n", e.Name)
	fmt.Println("После 10 лет поисков я нашел свой язык - Go! 🚀")
	fmt.Println("Сегодня мы настроим идеальное окружение для программирования!")
	time.Sleep(2 * time.Second)
}

func (e *EnvironmentChecker) CheckGoInstallation() {
	fmt.Println("\n🔍 Проверяем установку Go...")

	// Проверяем версию Go
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()

	if err != nil {
		e.Problems = append(e.Problems, "❌ Go не установлен или не настроен PATH")
		e.IsReady = false
		return
	}

	fmt.Printf("✅ Go установлен: %s", string(output))
	e.IsReady = true

	// Проверяем минимальную версию
	if strings.Contains(string(output), "go1.21") {
		fmt.Println("✅ Версия Go соответствует требованиям (1.21+)")
	} else {
		fmt.Println("⚠️  Рекомендуется обновить Go до версии 1.21 или выше")
	}
}

func (e *EnvironmentChecker) CheckEnvironmentVariables() {
	fmt.Println("\n🌍 Проверяем переменные окружения...")

	// Проверяем GOPATH
	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		fmt.Printf("✅ GOPATH: %s\n", gopath)
	} else {
		fmt.Println("✅ GOPATH не установлен (используется стандартный)")
	}

	// Проверяем GOROOT
	goroot := os.Getenv("GOROOT")
	if goroot != "" {
		fmt.Printf("✅ GOROOT: %s\n", goroot)
	} else {
		// Автоматически определяем GOROOT
		cmd := exec.Command("go", "env", "GOROOT")
		output, err := cmd.Output()
		if err == nil {
			fmt.Printf("✅ GOROOT: %s", string(output))
		}
	}

	// Показываем другие важные переменные
	fmt.Println("\n📊 Другие настройки:")
	cmd := exec.Command("go", "env", "GOOS")
	osType, _ := cmd.Output()
	fmt.Printf("   GOOS: %s", osType)

	cmd = exec.Command("go", "env", "GOARCH")
	arch, _ := cmd.Output()
	fmt.Printf("   GOARCH: %s", arch)
}

func (e *EnvironmentChecker) CheckProjectStructure() {
	fmt.Println("\n📁 Проверяем структуру проекта...")

	// Текущая рабочая директория
	wd, _ := os.Getwd()
	fmt.Printf("📂 Текущая директория: %s\n", wd)

	// Рекомендуемая структура
	recommendedStructure := []string{
		"cmd/",
		"internal/",
		"pkg/",
		"api/",
		"web/",
		"configs/",
		"scripts/",
		"build/",
		"deployments/",
		"test/",
	}

	fmt.Println("\n💡 Рекомендуемая структура проекта Go:")
	for _, dir := range recommendedStructure {
		fmt.Printf("   📁 %s\n", dir)
	}

	// Проверяем go.mod
	if _, err := os.Stat("go.mod"); err == nil {
		fmt.Println("✅ go.mod найден - проект инициализирован правильно")
	} else {
		fmt.Println("❌ go.mod не найден - запустите: go mod init <module-name>")
		e.Problems = append(e.Problems, "Отсутствует go.mod файл")
	}
}

func (e *EnvironmentChecker) CheckDevelopmentTools() {
	fmt.Println("\n🛠️  Проверяем инструменты разработки...")

	tools := map[string]string{
		"gopls":    "Language Server Protocol для Go",
		"staticcheck": "Статический анализатор кода",
		"golangci-lint": "Мульти-линтер для Go",
		"dlv":      "Отладчик Delve",
	}

	for tool, description := range tools {
		cmd := exec.Command("which", tool)
		if runtime.GOOS == "windows" {
			cmd = exec.Command("where", tool)
		}

		if err := cmd.Run(); err == nil {
			fmt.Printf("✅ %s: %s\n", tool, description)
		} else {
			fmt.Printf("⚠️  %s: не установлен (%s)\n", tool, description)
		}
	}

	// Проверяем IDE/редакторы
	fmt.Println("\n💻 Рекомендуемые редакторы:")
	editors := []string{"VSCode", "GoLand", "Vim с vim-go", "Neovim"}
	for _, editor := range editors {
		fmt.Printf("   ✨ %s\n", editor)
	}
}

func (e *EnvironmentChecker) ShowResults() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📊 РЕЗУЛЬТАТЫ ПРОВЕРКИ ОКРУЖЕНИЯ")
	fmt.Println(strings.Repeat("=", 50))

	if e.IsReady && len(e.Problems) == 0 {
		fmt.Println("🎉 ВАУ! Окружение настроено ИДЕАЛЬНО! 🎉")
		fmt.Println("🚀 Гоша готов покорять Go и найти работу мечты!")
	} else if e.IsReady {
		fmt.Printf("✅ Основное окружение готово, но есть %d проблем(ы)\n", len(e.Problems))
		for _, problem := range e.Problems {
			fmt.Printf("   🔧 %s\n", problem)
		}
	} else {
		fmt.Println("❌ Требуется дополнительная настройка окружения")
	}
}

func (e *EnvironmentChecker) MotivationalMessage() {
	messages := []string{
		"\n💫 Помни, Гоша: последовательность важнее перфекционизма!",
		"📚 Изучай по чуть-чуть каждый день - это лучше, чем ничего!",
		"🎯 С фокусом на Go ты обязательно достигнешь успеха!",
		"🤝 Твое СДВГ - это суперсила, а не ограничение!",
		"🚀 Всего 7 дней до 100-дневного финиша - ты молодец!",
	}

	fmt.Println("\n" + strings.Repeat("✨", 25))
	fmt.Println(messages[time.Now().Day()%len(messages)])
	fmt.Println(strings.Repeat("✨", 25))

	fmt.Printf("\n⏰ Время проверки: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("🎯 Следующая цель: Day94 - Продолжаем изучение!")
}

// Вспомогательная функция для создания структуры проекта
func CreateProjectStructure(projectName string) error {
	dirs := []string{
		"cmd/" + projectName,
		"internal/app",
		"pkg/utils",
		"api",
		"web/static",
		"web/templates",
		"configs",
		"scripts",
		"build",
		"deployments",
		"test",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		fmt.Printf("📁 Создана папка: %s\n", dir)
	}

	return nil
}
