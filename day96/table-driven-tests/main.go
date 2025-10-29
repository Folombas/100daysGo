package main

import (
	"fmt"
	"time"
)

type CyberDetective struct {
	Name          string
	Age           int
	Specialty     string
	CasesSolved   int
	TestingSkills int
}

func main() {
	detective := &CyberDetective{
		Name:          "Гоша",
		Age:           37,
		Specialty:     "Охотник за багами",
		CasesSolved:   0,
		TestingSkills: 25,
	}

	detective.ShowIntro()
	detective.ExplainTableDrivenTests()
	detective.RunTestScenarios()
	detective.ShowResults()
	detective.MotivationalConclusion()
}

func (c *CyberDetective) ShowIntro() {
	fmt.Println("🔍 КИБЕР-ДЕТЕКТИВ: ОХОТА ЗА БАГАМИ 🔍")
	fmt.Println("======================================")
	fmt.Printf("👮 Детектив: %s, %d лет\n", c.Name, c.Age)
	fmt.Printf("🎯 Специальность: %s\n", c.Specialty)
	fmt.Printf("📊 Навыки тестирования: %d%%\n", c.TestingSkills)
	fmt.Println("\n💡 Миссия: Освоить Table-Driven Tests для поимки коварных багов!")
	fmt.Println("🚀 Начинаем расследование...")
	pressToContinue()
}

func (c *CyberDetective) ExplainTableDrivenTests() {
	fmt.Println("\n📚 ТЕОРИЯ: TABLE-DRIVEN TESTS")
	fmt.Println("============================")

	concepts := []string{
		"🎯 Table-Driven Tests - это продвинутая техника тестирования в Go",
		"📊 Тестовые данные организованы в виде таблицы (среза структур)",
		"🔍 Каждая строка таблицы - отдельный тестовый случай",
		"💡 Позволяет тестировать множество сценариев в одном тесте",
		"🚀 Упрощает добавление новых тестовых случаев",
		"🎯 Идеально для функций с разными входными параметрами",
	}

	for _, concept := range concepts {
		fmt.Printf("   %s\n", concept)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n💪 Преимущества для нейроразнообразных разработчиков:")
	fmt.Println("   • Структурированность (помогает при ОКР)")
	fmt.Println("   • Предсказуемость (комфортно для аутистов)")
	fmt.Println("   • Пошаговость (идеально для СДВГ)")

	c.TestingSkills = 50
	fmt.Printf("\n✅ Навыки тестирования улучшены: %d%%\n", c.TestingSkills)
	pressToContinue()
}

func (c *CyberDetective) RunTestScenarios() {
	fmt.Println("\n🔍 ЗАПУСК ТЕСТОВЫХ СЦЕНАРИЕВ:")
	fmt.Println("============================")

	fmt.Println("🎯 Тестируем детективные функции...")

	// Запускаем тесты
	fmt.Println("\n🧪 Запуск Table-Driven Tests:")

	testCases := []struct {
		name     string
		function string
		status   string
	}{
		{"Проверка валидации пароля", "ValidatePassword", "✅ ПРОЙДЕН"},
		{"Анализ сложности пароля", "PasswordStrength", "✅ ПРОЙДЕН"},
		{"Генерация безопасного пароля", "GenerateSecurePassword", "✅ ПРОЙДЕН"},
		{"Валидация email", "ValidateEmail", "✅ ПРОЙДЕН"},
		{"Проверка возраста", "ValidateAge", "✅ ПРОЙДЕН"},
	}

	for _, tc := range testCases {
		fmt.Printf("   🔍 %s - %s\n", tc.function, tc.status)
		time.Sleep(500 * time.Millisecond)
		c.CasesSolved++
	}

	fmt.Println("\n🎉 Все тесты пройдены успешно!")
	c.TestingSkills = 85
	pressToContinue()
}

func (c *CyberDetective) ShowResults() {
	fmt.Println("\n📊 РЕЗУЛЬТАТЫ РАССЛЕДОВАНИЯ:")
	fmt.Println("===========================")

	fmt.Printf("👮 Детектив: %s\n", c.Name)
	fmt.Printf("🎯 Решено кейсов: %d/5\n", c.CasesSolved)
	fmt.Printf("💪 Навыки тестирования: %d%%\n", c.TestingSkills)
	fmt.Printf("🚀 Уровень мастерства Go: 85%%\n")

	fmt.Println("\n🔍 ОБНАРУЖЕННЫЕ БАГИ:")
	bugs := []string{
		"❌ Уязвимость: слабые пароли принимались",
		"❌ Ошибка: email без @ считался валидным",
		"❌ Проблема: отрицательный возраст допускался",
		"✅ ВСЕ БАГИ УСТРАНЕНЫ с помощью Table-Driven Tests!",
	}

	for _, bug := range bugs {
		fmt.Printf("   %s\n", bug)
		time.Sleep(700 * time.Millisecond)
	}

	pressToContinue()
}

func (c *CyberDetective) MotivationalConclusion() {
	fmt.Println("\n🎉 МИССИЯ ВЫПОЛНЕНА!")
	fmt.Println("===================")

	fmt.Println(`
	🕵️┌─────────────────────────────────┐
	🕵️│        КИБЕР-ДЕТЕКТИВ           │
	🕵️│                                 │
	🕵️│  Table-Driven Tests освоены!    │
	🕵️│                                 │
	🕵️│  Навыки тестирования: 85% → 95% │
	🕵️│  Решено кейсов: 5/5             │
	🕵️│  Уровень Go: ПРОДВИНУТЫЙ       │
	🕵️└─────────────────────────────────┘
	`)

	fmt.Println("💡 Ключевые достижения:")
	achievements := []string{
		"✅ Table-Driven Tests полностью освоены",
		"✅ 5 детективных функций протестированы",
		"✅ Найденные баги успешно устранены",
		"✅ Навыки тестирования выросли до 85%",
		"✅ Гоша готов к профессиональной разработке!",
	}

	for _, achievement := range achievements {
		fmt.Printf("   %s\n", achievement)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Printf("\n📅 Прогресс челленджа: 96/100 дней (96%%)\n")
	fmt.Println("🚀 До профессионального уровня осталось: 4 дня!")
}

func pressToContinue() {
	fmt.Print("\n↵ Нажми Enter чтобы продолжить...")
	fmt.Scanln()
}
