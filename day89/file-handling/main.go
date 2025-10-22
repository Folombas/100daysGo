package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

// Skill представляет навык Гофера
type Skill struct {
	Name        string  `json:"name"`
	Level       int     `json:"level"`
	XP          float64 `json:"xp"`
	Description string  `json:"description"`
}

// GopherProfile представляет профиль Гофера
type GopherProfile struct {
	Name       string    `json:"name"`
	Level      int       `json:"level"`
	TotalXP    float64   `json:"total_xp"`
	Skills     []Skill   `json:"skills"`
	LastUpdate time.Time `json:"last_update"`
}

func main() {
	fmt.Println("🎯 День 89: File Handling в Go - Прокачка Гофера!")
	fmt.Println("🐹 Ты выбрал Гофера! Давай прокачаем его через силу файлов I/O!")

	// Создаем начального Гофера
	gopher := createInitialGopher()
	fmt.Printf("🎉 Создан новый Гофер: %s (Уровень: %d)\n", gopher.Name, gopher.Level)

	// Демонстрация различных операций с файлами
	demonstrateFileOperations(gopher)

	// Показываем финальный прогресс
	showFinalProgress()
}

func createInitialGopher() GopherProfile {
	skills := []Skill{
		{Name: "Concurrency", Level: 1, XP: 0, Description: "Параллельное выполнение задач"},
		{Name: "Channels", Level: 1, XP: 0, Description: "Коммуникация между горутинами"},
		{Name: "Interfaces", Level: 1, XP: 0, Description: "Полиморфизм в Go"},
		{Name: "Error Handling", Level: 1, XP: 0, Description: "Обработка ошибок"},
		{Name: "File I/O", Level: 1, XP: 0, Description: "Работа с файлами"},
	}

	return GopherProfile{
		Name:       "SuperGopher89",
		Level:      1,
		TotalXP:    0,
		Skills:     skills,
		LastUpdate: time.Now(),
	}
}

func demonstrateFileOperations(gopher GopherProfile) {
	fmt.Println("\n📁 ДЕМОНСТРАЦИЯ ОПЕРАЦИЙ С ФАЙЛАМИ:")
	fmt.Println("====================================")

	// 1. Запись в JSON файл
	fmt.Println("\n1. 💾 Сохраняем профиль Гофера в JSON...")
	saveGopherToJSON(gopher, "skills.json")

	// 2. Чтение из JSON файла
	fmt.Println("\n2. 📖 Читаем профиль из JSON...")
	loadedGopher := loadGopherFromJSON("skills.json")
	fmt.Printf("   Загружен Гофер: %s (Уровень: %d)\n", loadedGopher.Name, loadedGopher.Level)

	// 3. Запись прогресса в текстовый файл
	fmt.Println("\n3. 📝 Записываем прогресс в текстовый файл...")
	writeProgressToFile(loadedGopher, "progress.txt")

	// 4. Чтение и вывод содержимого текстового файла
	fmt.Println("\n4. 👀 Читаем и выводим прогресс из файла...")
	readAndDisplayProgress("progress.txt")

	// 5. Копирование файла (резервная копия)
	fmt.Println("\n5. 🗂️ Создаем резервную копию навыков...")
	copyFile("skills.json", "backup_skills.json")
	fmt.Println("   ✅ Резервная копия создана!")

	// 6. Получение информации о файле
	fmt.Println("\n6. 📊 Получаем информацию о файлах...")
	getFileInfo("skills.json")
	getFileInfo("progress.txt")

	// 7. Обновляем навыки и сохраняем
	fmt.Println("\n7. ⚡ Прокачиваем навыки Гофера...")
	updatedGopher := levelUpGopher(loadedGopher)
	saveGopherToJSON(updatedGopher, "skills.json")
	fmt.Println("   🎉 Гофер прокачан! Проверь файл skills.json")
}

func saveGopherToJSON(gopher GopherProfile, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("❌ Ошибка создания файла: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(gopher); err != nil {
		fmt.Printf("❌ Ошибка кодирования JSON: %v\n", err)
		return
	}
	fmt.Printf("   ✅ Профиль сохранен в %s\n", filename)
}

func loadGopherFromJSON(filename string) GopherProfile {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("❌ Ошибка открытия файла: %v\n", err)
		return GopherProfile{}
	}
	defer file.Close()

	var gopher GopherProfile
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&gopher); err != nil {
		fmt.Printf("❌ Ошибка декодирования JSON: %v\n", err)
		return GopherProfile{}
	}

	gopher.LastUpdate = time.Now()
	return gopher
}

func writeProgressToFile(gopher GopherProfile, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("❌ Ошибка создания файла: %v\n", err)
		return
	}
	defer file.Close()

	content := fmt.Sprintf(`ПРОГРЕСС ГОФЕРА: %s
Уровень: %d
Общий опыт: %.1f
Последнее обновление: %s

НАВЫКИ:
`, gopher.Name, gopher.Level, gopher.TotalXP, gopher.LastUpdate.Format("2006-01-02 15:04:05"))

	for _, skill := range gopher.Skills {
		content += fmt.Sprintf("- %s: Уровень %d (Опыт: %.1f)\n  %s\n\n",
			skill.Name, skill.Level, skill.XP, skill.Description)
	}

	content += fmt.Sprintf("🎯 СДВГ-суперсила: Фокус на одном языке (%s) приносит результаты!\n", gopher.Name)

	if _, err := file.WriteString(content); err != nil {
		fmt.Printf("❌ Ошибка записи в файл: %v\n", err)
		return
	}
	fmt.Printf("   ✅ Прогресс записан в %s\n", filename)
}

func readAndDisplayProgress(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("❌ Ошибка открытия файла: %v\n", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("❌ Ошибка чтения файла: %v\n", err)
		return
	}

	fmt.Println("   📄 Содержимое файла progress.txt:")
	fmt.Println("   " + string(content))
}

func copyFile(src, dst string) {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("❌ Ошибка открытия исходного файла: %v\n", err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		fmt.Printf("❌ Ошибка создания файла назначения: %v\n", err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Printf("❌ Ошибка копирования файла: %v\n", err)
		return
	}
}

func getFileInfo(filename string) {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("❌ Ошибка получения информации о файле: %v\n", err)
		return
	}

	fmt.Printf("   📋 %s:\n", filename)
	fmt.Printf("     Размер: %d байт\n", info.Size())
	fmt.Printf("     Модифицирован: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
	fmt.Printf("     Режим: %s\n", info.Mode())
}

func levelUpGopher(gopher GopherProfile) GopherProfile {
	// Прокачиваем навыки
	for i := range gopher.Skills {
		gopher.Skills[i].Level += 1
		gopher.Skills[i].XP += 100
	}

	gopher.Level += 1
	gopher.TotalXP += 500
	gopher.LastUpdate = time.Now()

	fmt.Printf("   🚀 %s достиг уровня %d!\n", gopher.Name, gopher.Level)
	fmt.Println("   📈 Все навыки улучшены!")

	return gopher
}

func showFinalProgress() {
	fmt.Println("\n🎊 ФИНАЛЬНЫЙ ПРОГРЕСС:")
	fmt.Println("====================")
	fmt.Println("✅ Создан и сохранен профиль Гофера в JSON")
	fmt.Println("✅ Записан подробный прогресс в текстовый файл")
	fmt.Println("✅ Создана резервная копия навыков")
	fmt.Println("✅ Освоены основные операции File I/O в Go")
	fmt.Println("\n🎯 СДВГ-победа: Ты сфокусировался на Go и достиг результатов!")
	fmt.Println("🐹 Твой Гофер теперь сильнее, чем Слоник PHP, Питон Python и Крабик Rust!")
	fmt.Println("\n💪 Продолжай в том же духе! Уровень 100 уже близко!")
}
