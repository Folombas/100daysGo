package main

import (
	"fmt"
	"time"
)

type LanguagePet struct {
	Name        string
	Trend2025   string
	Salary      int
	Growth      float64
	FutureProof bool
}

func main() {
	fmt.Println("🔮 День 94: Будущее Go в 2025 - Почему Гофер твой билет в завтра!")
	fmt.Println("==================================================")

	pets := []LanguagePet{
		{
			Name:        "🐍 Питон Python",
			Trend2025:   "AI/ML доминирование, но насыщение рынка",
			Salary:      120000,
			Growth:      8.5,
			FutureProof: true,
		},
		{
			Name:        "🐘 Слоник PHP",
			Trend2025:   "Легаси проекты, медленный спад",
			Salary:      80000,
			Growth:      2.1,
			FutureProof: false,
		},
		{
			Name:        "🦀 Крабик Rust",
			Trend2025:   "Системное программирование, WebAssembly",
			Salary:      140000,
			Growth:      15.3,
			FutureProof: true,
		},
		{
			Name:        "🐫 Старый верблюд Perl",
			Trend2025:   "Нишевое использование, поддержка legacy",
			Salary:      90000,
			Growth:      1.2,
			FutureProof: false,
		},
		{
			Name:        "🐹 Гофер Golang",
			Trend2025:   "CLOUD-NATIVE ЛИДЕР, микросервисы, DevOps",
			Salary:      150000,
			Growth:      22.7,
			FutureProof: true,
		},
	}

	fmt.Println("\n🏪 Волшебный зоомагазин языков 2025:")
	fmt.Println("==================================================")
	
	for i, pet := range pets {
		status := "❌"
		if pet.FutureProof {
			status = "✅"
		}
		fmt.Printf("%s %d. %s\n", status, i+1, pet.Name)
		fmt.Printf("   📈 Тренд 2025: %s\n", pet.Trend2025)
		fmt.Printf("   💰 Зарплата: %d руб.\n", pet.Salary)
		fmt.Printf("   🚀 Рост: %.1f%%\n", pet.Growth)
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}

	// Фокус на Go
	gopher := pets[4]
	
	fmt.Println("🎯 ПОЧЕМУ ГОФЕР - ВЫБОР БУДУЩЕГО ДЛЯ НАС:")
	fmt.Println("==================================================")
	
	reasons := []struct {
		title   string
		details string
	}{
		{
			"☁️  Cloud-Native по умолчанию",
			"Go создан для облаков. Docker, Kubernetes, Terraform - всё на Go",
		},
		{
			"⚡ Производительность C++ с простотой Python",
			"Компилируется в нативный код, но читается как скриптовый язык",
		},
		{
			"🔧 Встроенная поддержка многозадачности",
			"Горутины и каналы - конкурентность из коробки",
		},
		{
			"📦 Одна бинарка - нулевые зависимости",
			"Идеально для микросервисов и контейнеризации",
		},
		{
			"🏢 Поддержка IT-гигантов",
			"Google, Uber, Twitch, Dropbox - все переходят на Go",
		},
		{
			"💰 Максимальная ROI по времени обучения",
			"Быстрый вход × высокие зарплаты = лучшая инвестиция",
		},
	}

	for i, reason := range reasons {
		fmt.Printf("%d. %s\n", i+1, reason.title)
		fmt.Printf("   %s\n", reason.details)
		time.Sleep(300 * time.Millisecond)
	}

	// Демонстрация современных возможностей Go
	fmt.Println("\n🚀 ДЕМО-ПРИМЕР: Почему Go идеален для 2025")
	fmt.Println("==================================================")
	
	// Простой пример конкурентности
	fmt.Println("🔄 Конкурентность в 3 строчки (основа микросервисов 2025):")
	
	ch := make(chan string, 3)
	
	go func() { ch <- "Микросервис А: обработал запрос" }()
	go func() { ch <- "Микросервис B: отправил данные" }()
	go func() { ch <- "Микросервис C: сохранил в кеш" }()
	
	for i := 0; i < 3; i++ {
		fmt.Printf("   📨 %s\n", <-ch)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("\n📊 ВЫВОД ДЛЯ ГОСТИ:")
	fmt.Printf("Выбрав %s, ты получаешь:\n", gopher.Name)
	fmt.Printf("• Зарплату на %d руб. выше среднего\n", gopher.Salary-100000)
	fmt.Printf("• Рост карьеры на %.1f%% быстрее рынка\n", gopher.Growth)
	fmt.Printf("• Гарантию востребованности до 2030+ года\n")
	
	fmt.Println("\n💡 ФИНАЛЬНАЯ ИСТИНА:")
	fmt.Println("В 2025 знание Go = пропуск в мир high-load, cloud и blockchain!")
	fmt.Println("Не меняй питомца - стань Гуру Гофера! 🐹✨")
	
	fmt.Println("\n🎯 ТВОЙ ПЛАН НА 2025:")
	plan := []string{
		"2025: Изучить основы Go"
		"2026: Освоить Go до middle уровня",
		"2027: Устроиться Go-разработчиком в cloud-компанию", 
		"2028: Стать lead в микросервисной архитектуре",
		"2029: Архитектор cloud-решений на Go",
	}
	
	for _, step := range plan {
		fmt.Printf("   ✅ %s\n", step)
		time.Sleep(400 * time.Millisecond)
	}
}
