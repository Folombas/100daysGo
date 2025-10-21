package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// PetStore представляет зоомагазин языков программирования
type PetStore struct {
	pets []*ProgrammingPet
}

type ProgrammingPet struct {
	Name     string    `json:"name"`
	Species  string    `json:"species"`
	Focus    bool      `json:"focus"`
	Progress int       `json:"progress"`
	LastFed  time.Time `json:"last_fed"`
}

type StudyPlan struct {
	Topics    []string  `json:"topics"`
	StartDate time.Time `json:"start_date"`
	Deadline  time.Time `json:"deadline"`
}

func main() {
	fmt.Println("🐹 ФОКУС НА GO: История Гоши и его гофера!")
	fmt.Println("===========================================")

	store := &PetStore{
		pets: []*ProgrammingPet{
			{Name: "Питоша", Species: "Python", Focus: false},
			{Name: "Слоник", Species: "PHP", Focus: false},
			{Name: "Крабик", Species: "Rust", Focus: false},
			{Name: "Гофер", Species: "Golang", Focus: true},
			{Name: "Перлуша", Species: "Perl", Focus: false},
		},
	}

	plan := &StudyPlan{
		Topics: []string{
			"Синтаксис и основы", "Структуры и интерфейсы", "Конкурентность",
			"Стандартная библиотека", "Тестирование", "Web-разработка",
		},
		StartDate: time.Now(),
		Deadline:  time.Now().Add(100 * 24 * time.Hour),
	}

	fmt.Println("🎯 СИТУАЦИЯ: Гоша с СДВГ и ОКР 5 лет метался между языками...")
	fmt.Println("💡 РЕШЕНИЕ: Выбрать ОДНОГО питомца - Гофера - и сфокусироваться!")

	fmt.Println("\n📚 СТАНДАРТНАЯ БИБЛИОТЕКА GO В ДЕЙСТВИИ:")
	fmt.Println("=======================================")

	store.demoFileOperations()
	store.demoJSONOperations(plan)
	store.demoStringOperations()
	store.demoTimeOperations(plan)
	store.demoLogging()

	fmt.Println("\n🎉 РЕЗУЛЬТАТ: Гоша наконец-то сфокусировался и изучает Go!")
	fmt.Println("🏆 Стандартная библиотека - его верный помощник в этом пути!")
}

func (p *PetStore) demoFileOperations() {
	fmt.Println("\n1. 📁 РАБОТА С ФАЙЛАМИ (пакет os):")

	dataDir := "study_data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Printf("Ошибка создания директории: %v", err)
		return
	}

	filePath := filepath.Join(dataDir, "study_plan.txt")
	content := "Фокус на Go: Изучать только одного гофера!\n"

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		log.Printf("Ошибка записи файла: %v", err)
		return
	}

	if data, err := os.ReadFile(filePath); err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
	} else {
		fmt.Printf("   ✅ Создан и прочитан файл: %s\n", string(data))
	}
}

func (p *PetStore) demoJSONOperations(plan *StudyPlan) {
	fmt.Println("\n2. 📊 РАБОТА С JSON (пакет encoding/json):")

	jsonData, err := json.MarshalIndent(plan, "", "  ")
	if err != nil {
		log.Printf("Ошибка маршалинга JSON: %v", err)
		return
	}

	jsonPath := filepath.Join("study_data", "plan.json")
	if err := os.WriteFile(jsonPath, jsonData, 0644); err != nil {
		log.Printf("Ошибка сохранения JSON: %v", err)
	} else {
		fmt.Printf("   ✅ JSON сохранен: %s\n", jsonPath)
	}
}

func (p *PetStore) demoStringOperations() {
	fmt.Println("\n3. 🔤 РАБОТА СО СТРОКАМИ (пакет strings):")

	distractions := "Python,PHP,Rust,Perl,JavaScript,Java,C#,Ruby"
	focusPet := "Golang"

	languages := strings.Split(distractions, ",")
	cleaned := strings.ReplaceAll(distractions, "Python", "ИГНОР")
	cleaned = strings.ReplaceAll(cleaned, "PHP", "ИГНОР")

	fmt.Printf("   🎯 Фокус-язык: %s\n", focusPet)
	fmt.Printf("   🚫 Отвлекающие: %s\n", cleaned)
	fmt.Printf("   📊 Всего языков: %d\n", len(languages))
}

func (p *PetStore) demoTimeOperations(plan *StudyPlan) {
	fmt.Println("\n4. ⏰ РАБОТА СО ВРЕМЕНЕМ (пакет time):")

	now := time.Now()
	daysStudied := int(now.Sub(plan.StartDate).Hours() / 24)
	daysRemaining := int(plan.Deadline.Sub(now).Hours() / 24)

	fmt.Printf("   🗓️  Начало: %s\n", plan.StartDate.Format("02.01.2006"))
	fmt.Printf("   🎯 Дедлайн: %s\n", plan.Deadline.Format("02.01.2006"))
	fmt.Printf("   📅 Дней изучения: %d\n", daysStudied)
	fmt.Printf("   ⏳ Осталось дней: %d\n", daysRemaining)
}

func (p *PetStore) demoLogging() {
	fmt.Println("\n5. 📝 ЛОГИРОВАНИЕ (пакет log):")

	logFile, err := os.OpenFile("study_data/progress.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Ошибка открытия лог-файла: %v", err)
		return
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := log.New(multiWriter, "GO-FOCUS: ", log.Ldate|log.Ltime)

	logger.Println("Старт программы фокусировки на Go")
	logger.Println("Выбран питомец: Гофер (Golang)")
	logger.Println("Фокус установлен на одного питомца!")

	fmt.Printf("   ✅ Логи записаны в study_data/progress.log\n")
}
