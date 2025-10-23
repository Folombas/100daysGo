package main

import (
	"fmt"
	"time"
)

type LanguagePet struct {
	Name     string
	CareTime time.Duration
	Complexity int
	JobDemand  int
}

func main() {
	fmt.Println("🐾 Добро пожаловать в волшебный зоомагазин языков программирования!")
	fmt.Println("================================================")

	pets := []LanguagePet{
		{"🐍 Питон Python", 4 * time.Hour, 8, 7},
		{"🐘 Слоник PHP", 3 * time.Hour, 6, 5},
		{"🦀 Крабик Rust", 5 * time.Hour, 9, 8},
		{"🐫 Старый верблюд Perl", 4 * time.Hour, 7, 3},
		{"🐹 Гофер Golang", 2 * time.Hour, 4, 9},
	}

	fmt.Println("\nДоступные питомцы-языки:")
	for i, pet := range pets {
		fmt.Printf("%d. %s\n", i+1, pet.Name)
	}

	fmt.Println("\n🤔 Почему Гоша выбирает Гофера?")
	fmt.Println("================================================")

	gopher := pets[4]
	
	reasons := []string{
		"🎯 **Фокус вместо метаний**: Go имеет чёткую спецификацию и минимум способов сделать одно действие",
		"⚡ **Скорость обучения**: Простой синтаксис позволяет быстро достичь продуктивности",
		"🏃 **Производительность**: Быстрая компиляция и выполнение дают мгновенную обратную связь",
		"💼 **Востребованность**: Высокий спрос на рынке труда с достойными зарплатами",
		"🧠 **Психологический комфорт**: Статическая типизация и явные ошибки снижают тревожность",
	}

	for i, reason := range reasons {
		fmt.Printf("%d. %s\n", i+1, reason)
	}

	fmt.Println("\n📊 Сравнительные характеристики:")
	fmt.Println("Язык       | Время ухода | Сложность | Востребованность")
	fmt.Println("-----------|-------------|-----------|-----------------")
	for _, pet := range pets {
		fmt.Printf("%-10s | %-11v | %-9d | %-15d\n", 
			pet.Name, pet.CareTime, pet.Complexity, pet.JobDemand)
	}

	fmt.Println("\n🎯 Заключение для Гоши:")
	fmt.Printf("Выбрав %s, ты получаешь:\n", gopher.Name)
	fmt.Printf("- Всего %v ежедневных занятий вместо распыления\n", gopher.CareTime)
	fmt.Printf("- Сложность всего %d/10 вместо постоянной перегрузки\n", gopher.Complexity)
	fmt.Printf("- Востребованность %d/10 для быстрого трудоустройства\n", gopher.JobDemand)

	fmt.Println("\n✨ Гоша, твой путь к успеху начинается с фокуса на Go!")
	fmt.Println("Доведи уход за гофером до автоматизма и работа не заставит себя ждать!")
}
