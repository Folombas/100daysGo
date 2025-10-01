package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("📦 GO MOD VENDOR - ДЕМОНСТРАЦИЯ ВЕНДОРИНГА")
	fmt.Println("==========================================")

	// Показываем текущую директорию
	wd, _ := os.Getwd()
	fmt.Printf("📁 Рабочая директория: %s\n\n", wd)

	// 1. Что такое вендоринг
	fmt.Println("🎯 ЧТО ТАКОЕ VENDOR В GO:")
	concepts := []string{
		"✅ vendor/ - папка с локальными копиями зависимостей",
		"✅ go mod vendor - создает vendor папку",
		"✅ -mod=vendor - флаг для использования vendor",
		"✅ Изоляция - проект не зависит от внешних репозиториев",
		"✅ Воспроизводимость - гарантия одинаковых версий",
	}

	for _, concept := range concepts {
		fmt.Println("   ", concept)
	}
	fmt.Println()

	// 2. Создаем демонстрационный проект с зависимостями
	fmt.Println("🔧 СОЗДАЕМ ДЕМО-ПРОЕКТ С ЗАВИСИМОСТЯМИ:")
	setupDemoProject()

	// 3. Показываем зависимости до вендоринга
	fmt.Println("📊 ЗАВИСИМОСТИ ДО VENDOR:")
	showDependencies()

	// 4. Выполняем go mod vendor
	fmt.Println("\n🏗️  ВЫПОЛНЯЕМ GO MOD VENDOR:")
	runGoModVendor()

	// 5. Показываем структуру vendor
	fmt.Println("\n📁 СТРУКТУРА VENDOR ПАПКИ:")
	showVendorStructure()

	// 6. Демонстрация сборки с vendor
	fmt.Println("\n🔨 ДЕМОНСТРАЦИЯ СБОРКИ С VENDOR:")
	demonstrateVendorBuild()

	// 7. Практические сценарии использования
	fmt.Println("\n💡 ПРАКТИЧЕСКИЕ СЦЕНАРИИ:")
	scenarios := []struct {
		scenario string
		benefit  string
	}{
		{"CI/CD без интернета", "Сборка без доступа к внешним репозиториям"},
		{"Гарантия версий", "Исключает проблемы с удалением пакетов"},
		{"Стабильные билды", "Одинаковые зависимости в разных средах"},
		{"Аудит безопасности", "Полный контроль над используемым кодом"},
		{"Оффлайн разработка", "Работа без постоянного доступа к интернету"},
	}

	for i, scenario := range scenarios {
		fmt.Printf("   %d. %s - %s\n", i+1, scenario.scenario, scenario.benefit)
	}

	fmt.Println("\n🎯 ВЫВОДЫ:")
	fmt.Println("✅ vendor/ - мощный инструмент для изоляции зависимостей")
	fmt.Println("✅ Обязателен для enterprise и security-critical проектов")
	fmt.Println("✅ Гарантирует воспроизводимость сборок")
	fmt.Println("✅ Требует больше места и управления версиями")
	fmt.Println("✅ -mod=vendor обеспечивает использование локальных копий")
}

func setupDemoProject() {
	// Создаем основной файл с использованием внешних зависимостей
	mainContent := `package main

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("📦 Демонстрация работы с зависимостями")

	// Используем golang.org/x/text
	caser := cases.Title(language.Russian)
	title := caser.String("привет мир")
	fmt.Printf("Заголовок: %s\n", title)

	// Используем github.com/pkg/errors
	err := processData()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println("✅ Все зависимости работают!")
}

func processData() error {
	// Создаем ошибку с stack trace
	return errors.Wrap(
		fmt.Errorf("данные невалидны"),
		"processData failed",
	)
}
`

	err := os.WriteFile("main.go", []byte(mainContent), 0644)
	if err != nil {
		fmt.Printf("❌ Ошибка создания main.go: %v\n", err)
		return
	}

	// Добавляем зависимости
	fmt.Println("   📥 Добавляем демонстрационные зависимости...")
	exec.Command("go", "get", "golang.org/x/text").Run()
	exec.Command("go", "get", "github.com/pkg/errors").Run()

	fmt.Println("✅ Демо-проект создан с внешними зависимостями")
}

func showDependencies() {
	// Показываем go.mod
	data, _ := os.ReadFile("go.mod")
	fmt.Printf("go.mod:\n%s\n", string(data))

	// Показываем граф зависимостей
	cmd := exec.Command("go", "mod", "graph")
	output, _ := cmd.Output()

	fmt.Println("Граф зависимостей:")
	lines := strings.Split(string(output), "\n")
	for i, line := range lines {
		if i < 5 && line != "" { // Показываем первые 5 зависимостей
			fmt.Printf("   %s\n", line)
		}
	}
	if len(lines) > 5 {
		fmt.Printf("   ... и ещё %d зависимостей\n", len(lines)-5)
	}
}

func runGoModVendor() {
	// Выполняем go mod vendor
	cmd := exec.Command("go", "mod", "vendor")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("❌ Ошибка выполнения go mod vendor: %v\n", err)
		return
	}

	fmt.Println("✅ vendor/ папка создана успешно!")

	// Показываем статистику
	vendorPath := "vendor"
	if info, err := os.Stat(vendorPath); err == nil && info.IsDir() {
		size := getDirSize(vendorPath)
		fmt.Printf("📊 Размер vendor папки: %.2f MB\n", float64(size)/(1024*1024))
	}
}

func showVendorStructure() {
	vendorPath := "vendor"

	// Показываем содержимое vendor
	cmd := exec.Command("find", vendorPath, "-type", "d", "-maxdepth", "2")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("❌ Ошибка анализа vendor: %v\n", err)
		return
	}

	lines := strings.Split(string(output), "\n")
	fmt.Println("Структура vendor/:")
	for i, line := range lines {
		if i < 10 && line != "" { // Показываем первые 10 элементов
			fmt.Printf("   %s\n", line)
		}
	}
	if len(lines) > 10 {
		fmt.Printf("   ... и ещё %d папок/файлов\n", len(lines)-10)
	}
}

func demonstrateVendorBuild() {
	fmt.Println("\n   🔨 Сборка с использованием vendor...")

	// Сборка с -mod=vendor
	buildCmd := exec.Command("go", "build", "-mod=vendor", "-o", "demo-app", ".")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr

	err := buildCmd.Run()
	if err != nil {
		fmt.Printf("   ❌ Ошибка сборки с vendor: %v\n", err)
		return
	}

	fmt.Println("   ✅ Успешная сборка с -mod=vendor!")

	// Проверяем что бинарник работает
	if _, err := os.Stat("demo-app"); err == nil {
		fmt.Println("   🚀 Запускаем собранное приложение...")
		runCmd := exec.Command("./demo-app")
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr
		runCmd.Run()

		// Убираем временный файл
		os.Remove("demo-app")
	}
}

func getDirSize(path string) int64 {
	var size int64
	filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size
}
