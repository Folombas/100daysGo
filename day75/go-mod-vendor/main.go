package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"go-mod-vendor/internal/vendorLogic"
)

func main() {
	fmt.Println("📦 DAY 75: GO MOD VENDOR - КОПИЛКА ЗАВИСИМОСТЕЙ")
	fmt.Println("================================================")

	fmt.Println("📖 Легенда дня:")
	fmt.Println("   Гоша копит 600 рублей в копилку, как и мы копим зависимости в vendor.")
	fmt.Println("   Vendor — это локальная копилка всех зависимостей проекта.")
	fmt.Println("   Позволяет работать без интернета и гарантирует повторяемость сборок.")

	// Проверяем наличие vendor
	if _, err := os.Stat("vendor"); os.IsNotExist(err) {
		fmt.Println("⚠  Папка vendor не найдена!")
		fmt.Println("   Создаем копилку зависимостей...")

		// Создаем vendor
		cmd := exec.Command("go", "mod", "vendor")
		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ Ошибка: %v\n", err)
			return
		}

		fmt.Println("✅ Vendor создан! Все зависимости в локальной копилке.")
	} else {
		fmt.Println("✅ Vendor уже существует! Копилка зависимостей готова.")
	}

	// Демонстрация работы с vendor
	vendorLogic.ShowVendorBenefits()

	// Проверяем содержимое vendor
	checkVendorContents()

	fmt.Println("\n🎯 Геймификация:")
	fmt.Println("   Уровень: 'Вендор-копилка' достигнут!")
	fmt.Println("   +100 XP за создание локальной копилки зависимостей")
	fmt.Println("   Следующий уровень: 'Мастер изоляции зависимостей'")
}

func checkVendorContents() {
	fmt.Println("\n🔍 Содержимое vendor/")

	// Простой обход vendor
	vendorPath := "vendor"

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(vendorPath, path)
		depth := len(filepath.SplitList(relPath))

		if info.IsDir() && depth == 1 {
			fmt.Printf("   ├── 📁 %s/\n", filepath.Base(path))
		} else if !info.IsDir() && depth == 2 && filepath.Ext(path) == ".go" {
			fmt.Printf("   │   └── 📄 %s\n", filepath.Base(path))
		}

		return nil
	}

	filepath.Walk(vendorPath, walkFunc)
	fmt.Println("   └── ... (все зависимости сохранены локально)")
}
