package main

import (
	"fmt"
)

type LanguageHistory struct {
	Name      string
	Born      int
	Creators  []string
	Motto     string
	Strength  string
}

func main() {
	fmt.Println("📖 История волшебного зоомагазина: Рождение Гофера")
	fmt.Println("==================================================")

	histories := []LanguageHistory{
		{
			Name:     "🐍 Питон Python",
			Born:     1991,
			Creators: []string{"Гвидо ван Россум"},
			Motto:    "Простота — залог надежности",
			Strength: "Универсальность",
		},
		{
			Name:     "🐘 Слоник PHP",
			Born:     1995,
			Creators: []string{"Расмус Лердорф"},
			Motto:    "Для веба и не только",
			Strength: "Веб-разработка",
		},
		{
			Name:     "🦀 Крабик Rust",
			Born:     2010,
			Creators: []string{"Грейдон Хор"},
			Motto:    "Безопасность без сборщика мусора",
			Strength: "Безопасность памяти",
		},
		{
			Name:     "🐫 Старый верблюд Perl",
			Born:     1987,
			Creators: []string{"Ларри Уолл"},
			Motto:    "Есть несколько способов сделать это",
			Strength: "Текстовые обработки",
		},
		{
			Name:     "🐹 Гофер Golang",
			Born:     2007,
			Creators: []string{"Роб Пайк", "Кен Томпсон", "Роберт Гризмер"},
			Motto:    "Простота, эффективность, надежность",
			Strength: "ПРОДУКТИВНОСТЬ",
		},
	}

	fmt.Println("\n🌟 История создания Гофера:")
	fmt.Println("================================")
	
	gopher := histories[4]
	
	milestones := []struct {
		year    int
		event   string
		details string
	}{
		{2007, "Зарождение идеи", "Разочарование в сложности C++\nТри титана Google объединились"},
		{2009, "Публичный анонс", "Открытый исходный код\nСообщество сразу оценило простоту"},
		{2012, "Версия 1.0", "Стабильность и гарантии\nГотов к продакшену"},
		{2015, "Go 1.5", "Самодостаточность\nКомпилятор написан на Go"},
		{2020, "Go 1.15", "Современные фичи\nМодули, дженерики в пути"},
		{2023, "Доминирование", "Docker, Kubernetes, Cloud\nФундамент современного облака"},
	}

	fmt.Printf("🐹 %s - ребенок великих родителей:\n", gopher.Name)
	for _, creator := range gopher.Creators {
		fmt.Printf("   👨‍💻 %s\n", creator)
	}
	fmt.Printf("🎯 Девиз: \"%s\"\n", gopher.Motto)

	fmt.Println("\n📅 Хронология великого пути:")
	for _, milestone := range milestones {
		fmt.Printf("%d: %s\n", milestone.year, milestone.event)
		fmt.Printf("   %s\n", milestone.details)
	}

	fmt.Println("\n💡 Почему Гоша выбрал ИМЕННО ГОФЕРА:")
	reasons := []string{
		"🎯 **ЦЕЛЕПОЛАГАНИЕ**: Создан для решения реальных проблем Google",
		"⚡ **ПРОСТОТА**: Минималистичный синтаксис без лишней магии",
		"🏃 **СКОРОСТЬ**: Компиляция за секунды, выполнение как у C++",
		"🔧 **ПРАКТИЧНОСТЬ**: Встроенные инструменты для настоящей работы",
		"🌐 **БУДУЩЕЕ**: Язык облачных технологий и микросервисов",
	}

	for i, reason := range reasons {
		fmt.Printf("%d. %s\n", i+1, reason)
	}

	fmt.Println("\n✨ МОРАЛЬ ДЛЯ ГОШИ:")
	fmt.Println("Гофер рожден великими умами, прошел путь от идеи до лидера.")
	fmt.Println("Его история — гарантия, что ты инвестируешь время в ПЕРСПЕКТИВУ!")
	fmt.Println("Не меняй питомца — расти вместе с ним до мастера!")
	
	fmt.Println("\n🚀 Следующий шаг: установка Go и первый 'Hello, Gopher!'")
}
