package main

import (
	"fmt"
	"os"
)

// 🚷 Демонстрация БАЗОВОГО использования goto
func demonstrateBasicGoto() {
	fmt.Println("\n🎯 1. БАЗОВОЕ ИСПОЛЬЗОВАНИЕ GOTO")
	fmt.Println("===============================")

	fmt.Println("🚦 Начало функции...")

	goto skipCode // 🎪 Перепрыгиваем блок кода

	// Этот код будет пропущен
	fmt.Println("❌ Этот текст никогда не увидим!")

skipCode:
	fmt.Println("✅ Перепрыгнули с помощью goto!")
}

// ⚠️ Демонстрация ПЛОХОГО использования goto (спагетти-код)
func demonstrateSpaghettiGoto() {
	fmt.Println("\n🍝 2. ПЛОХОЙ ПРИМЕР: SPAGHETTI-КОД С GOTO")
	fmt.Println("========================================")

	i := 0

start:
	fmt.Printf("🔁 Итерация %d\n", i)
	i++

	if i < 3 {
		goto middle
	}

	if i == 3 {
		goto end
	}

middle:
	fmt.Println("🔄 В середине...")
	goto start

end:
	fmt.Println("🏁 Конец спагетти-кода!")
}

// ✅ Демонстрация ПРАВИЛЬНОГО использования goto (обработка ошибок)
func demonstrateGoodGoto() error {
	fmt.Println("\n✅ 3. ПРАВИЛЬНОЕ ИСПОЛЬЗОВАНИЕ: ОБРАБОТКА ОШИБОК")
	fmt.Println("==============================================")

	file1, err := os.Create("temp1.txt")
	if err != nil {
		return err
	}

	file2, err := os.Create("temp2.txt")
	if err != nil {
		file1.Close() // ❌ Дублирование кода очистки
		return err
	}

	file3, err := os.Create("temp3.txt")
	if err != nil {
		file1.Close() // ❌ Ещё больше дублирования
		file2.Close()
		return err
	}

	// Работа с файлами...
	fmt.Println("📁 Файлы созданы успешно!")

	// Очистка ресурсов
	file1.Close()
	file2.Close()
	file3.Close()

	return nil
}

// ✅ УЛУЧШЕННАЯ версия с goto для очистки ресурсов (ИСПРАВЛЕННАЯ)
func demonstrateGoodGotoImproved() error {
	fmt.Println("\n🎪 4. УЛУЧШЕННАЯ ВЕРСИЯ С GOTO")
	fmt.Println("==============================")

	var err error
	var file1, file2, file3 *os.File

	file1, err = os.Create("temp1.txt")
	if err != nil {
		return err
	}

	file2, err = os.Create("temp2.txt")
	if err != nil {
		goto cleanupFile1
	}

	file3, err = os.Create("temp3.txt")
	if err != nil {
		goto cleanupFiles
	}

	fmt.Println("📁 Файлы созданы успешно!")

	file3.Close()
	file2.Close()
	file1.Close()
	return nil

cleanupFiles:
	if file2 != nil {
		file2.Close()
	}
cleanupFile1:
	if file1 != nil {
		file1.Close()
	}
	return err
}

// ✅ СОВРЕМЕННАЯ версия с defer (лучшая практика)
func demonstrateModernApproach() error {
	fmt.Println("\n🌟 5. СОВРЕМЕННЫЙ ПОДХОД С DEFER")
	fmt.Println("===============================")

	file1, err := os.Create("temp1.txt")
	if err != nil {
		return err
	}
	defer file1.Close()

	file2, err := os.Create("temp2.txt")
	if err != nil {
		return err
	}
	defer file2.Close()

	file3, err := os.Create("temp3.txt")
	if err != nil {
		return err
	}
	defer file3.Close()

	fmt.Println("📁 Файлы созданы успешно с использованием defer!")
	return nil
}

// 🔄 АЛЬТЕРНАТИВЫ GOTO с современными конструкциями Go
func demonstrateAlternatives() {
	fmt.Println("\n🔄 6. АЛЬТЕРНАТИВЫ GOTO В СОВРЕМЕННОМ GO")
	fmt.Println("=======================================")

	fmt.Println("🎯 Альтернатива 1: Использование defer")
	if err := demonstrateModernApproach(); err != nil {
		fmt.Println("❌ Ошибка:", err)
	}

	fmt.Println("\n🎯 Альтернатива 2: Разделение на функции")
	if err := createAndProcessFiles(); err != nil {
		fmt.Println("❌ Ошибка:", err)
	} else {
		fmt.Println("✅ Файлы обработаны успешно!")
	}
}

func createAndProcessFiles() error {
	file1, err := os.Create("temp1.txt")
	if err != nil {
		return err
	}
	defer file1.Close()

	file2, err := os.Create("temp2.txt")
	if err != nil {
		return err
	}
	defer file2.Close()

	return nil
}

func main() {
	fmt.Println("🎪 ДЕМОНСТРАЦИЯ OPERATORA GOTO В GO!")
	fmt.Println("====================================")

	demonstrateBasicGoto()
	demonstrateSpaghettiGoto()

	_ = demonstrateGoodGoto()
	_ = demonstrateGoodGotoImproved()
	_ = demonstrateModernApproach()

	demonstrateAlternatives()

	// Очищаем временные файлы
	cleanupTempFiles()

	fmt.Println("\n🎯 ВЫВОДЫ: Используйте defer вместо goto в 99% случаев!")
}

func cleanupTempFiles() {
	files := []string{"temp1.txt", "temp2.txt", "temp3.txt"}
	for _, file := range files {
		os.Remove(file)
	}
}
