package main

import (
	"fmt"
	"modules-dependencies-go/internal/calculator"
	"modules-dependencies-go/pkg/book"
	"modules-dependencies-go/pkg/delivery"
	"modules-dependencies-go/pkg/motivation"
	"modules-dependencies-go/pkg/train"
	"strings"
	"time"
)

func main() {
	fmt.Println("🚂 DAY 72: Code Organization - Modules & Dependencies")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("Сюжет: Бизнес-Путешествие в Апрелевку, Docker на Go и организация кода")

	// Инициализируем зависимости
	fmt.Println("📦 ИНИЦИАЛИЗАЦИЯ ЗАВИСИМОСТЕЙ:")
	fmt.Println(strings.Repeat("-", 40))

	// Создаем экземпляры с использованием зависимостей
	trainJourney := train.NewJourney("Верхние Лихоборы", "Апрелевка", 9, 30)
	dockerBook := book.NewDockerBook("Docker: Полное руководство", 450, true)
	courierDelivery := delivery.NewDelivery("Север Москвы", "Апрелевка", 1, "жирный")
	motivator := motivation.NewMotivator()
	scoreCalc := calculator.NewScoreCalculator()

	// Демонстрация работы зависимостей
	fmt.Printf("   ✅ Поездка: %s\n", trainJourney.Route())
	fmt.Printf("   ✅ Книга: %s\n", dockerBook.Info())
	fmt.Printf("   ✅ Доставка: %s\n", courierDelivery.Details())
	fmt.Println("   ✅ Мотиватор: инициализирован")
	fmt.Println("   ✅ Калькулятор очков: готов")

	// Симуляция дня
	simulateDay(trainJourney, dockerBook, courierDelivery, motivator, scoreCalc)

	// Вывод итогов
	printSummary(trainJourney, dockerBook, courierDelivery, motivator, scoreCalc)
}

func simulateDay(t *train.Journey, b *book.DockerBook, d *delivery.Delivery, m *motivation.Motivator, c *calculator.ScoreCalculator) {
	fmt.Println("🌅 НАЧАЛО ДНЯ:")
	fmt.Println(strings.Repeat("-", 40))

	// Утренний ритуал
	fmt.Println("   🕤 09:30 - Подъём")
	fmt.Println("   🪒 Побрился, помылся в душе")
	fmt.Println("   🍳 Покушал вкусно")
	time.Sleep(300 * time.Millisecond)

	// Запуск доставки (зависимость от пакета delivery)
	fmt.Printf("\n   🚚 %s\n", d.Start())
	deliveryScore := d.CalculateScore()
	c.AddScore("доставка", deliveryScore)
	fmt.Printf("   🏆 Очки за доставку: +%d\n", deliveryScore)
	time.Sleep(400 * time.Millisecond)

	// Поездка в поезде (зависимость от пакета train)
	fmt.Printf("\n   🚂 %s\n", t.StartJourney())

	// Чтение книги в поезде (зависимость от пакета book)
	fmt.Printf("\n   📖 В поезде: %s\n", b.Open())
	pagesRead := 52
	fmt.Printf("   📄 Прочитано страниц: %d\n", pagesRead)
	readingScore := b.CalculateReadingScore(pagesRead)
	c.AddScore("чтение", readingScore)
	fmt.Printf("   🏆 Очки за чтение: +%d\n", readingScore)
	time.Sleep(400 * time.Millisecond)

	// Наблюдения в вагоне
	fmt.Println("\n   👀 Наблюдения в вагоне:")
	fmt.Println("      • Один бомж лежит на сидениях в центре вагона")
	fmt.Println("      • Под стук колёс читается особенно хорошо")
	fmt.Println("      • Электричка ритмично покачивается")

	// Осознание про Docker и Go
	fmt.Println("\n   💡 ОСОЗНАНИЕ:")
	fmt.Println("      Docker написан на Go!")
	fmt.Println("      Значит, изучая Go, я изучаю фундамент Docker")

	// Использование мотивационных фраз (зависимость от пакета motivation)
	fmt.Println("\n   🎯 МОТИВАЦИОННЫЕ ФРАЗЫ:")
	motivations := m.GetMotivationalPhrases(5)
	for i, phrase := range motivations {
		fmt.Printf("      %d. %s\n", i+1, phrase)
		motivationScore := 15
		c.AddScore("мотивация", motivationScore)
		time.Sleep(200 * time.Millisecond)
	}

	// Завершение поездки
	fmt.Printf("\n   🏁 %s\n", t.EndJourney())
	trainScore := t.CalculateJourneyScore()
	c.AddScore("поездка", trainScore)
	fmt.Printf("   🏆 Очки за поездку: +%d\n", trainScore)
}

func printSummary(t *train.Journey, b *book.DockerBook, d *delivery.Delivery, m *motivation.Motivator, c *calculator.ScoreCalculator) {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("📊 ИТОГИ ДНЯ 72:")
	fmt.Println(strings.Repeat("-", 70))

	// Информация о поездке
	fmt.Println("   🚂 ПОЕЗДКА:")
	fmt.Printf("      Маршрут: %s → %s\n", t.From, t.To)
	fmt.Printf("      Время: %d:%02d\n", t.DepartureHour, t.DepartureMinute)
	fmt.Printf("      Длительность: %.1f часа\n", t.Duration)

	// Информация о доставке
	fmt.Println("\n   🚚 ДОСТАВКА:")
	fmt.Printf("      Откуда: %s\n", d.From)
	fmt.Printf("      Куда: %s\n", d.To)
	fmt.Printf("      Количество: %d заказ\n", d.Count)
	fmt.Printf("      Тип: %s\n", d.Type)

	// Информация о книге
	fmt.Println("\n   📖 ОБУЧЕНИЕ:")
	fmt.Printf("      Книга: %s\n", b.Title)
	fmt.Printf("      Страниц: %d\n", b.TotalPages)
	fmt.Printf("      Формат: %s\n", func() string {
		if b.IsPhysical { return "бумажная" }
		return "электронная"
	}())
	fmt.Println("      Тема: Docker (написан на Go)")

	// Мотивационные изречения
	fmt.Println("\n   💪 5 ВДОХНОВЛЯЮЩИХ ИЗРЕЧЕНИЙ:")
	inspirations := m.GetInspirationalQuotes(5)
	for i, quote := range inspirations {
		fmt.Printf("      %d. %s\n", i+1, quote)
	}

	// 10 мотивационных фраз (ещё 5 дополнительных)
	fmt.Println("\n   🎯 ЕЩЁ 5 МОТИВАЦИОННЫХ ФРАЗ:")
	extraMotivations := m.GetExtraMotivationalPhrases(5)
	for i, phrase := range extraMotivations {
		fmt.Printf("      %d. %s\n", i+6, phrase)
	}

	// Очки и геймификация
	totalScore := c.GetTotalScore()
	level := c.CalculateLevel(totalScore)

	fmt.Println("\n   🏆 СИСТЕМА ОЧКОВ:")
	categories := map[string]string{
		"доставка": "Доставка заказа",
		"чтение":   "Чтение книги",
		"поездка":  "Поездка в электричке",
		"мотивация": "Мотивационные фразы",
	}

	for category, description := range categories {
		score := c.GetCategoryScore(category)
		fmt.Printf("      %s: +%d очков\n", description, score)
	}

	fmt.Printf("\n   🎯 ОБЩИЙ СЧЁТ: %d очков\n", totalScore)
	fmt.Printf("   📈 УРОВЕНЬ: %s\n", level)

	// Заключение
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("💡 ВЫВОД О CODE ORGANIZATION:")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("   Сегодняшний день демонстрирует важность организации кода:")
	fmt.Println("   1. Модули (Go modules) - как отдельные вагоны поезда")
	fmt.Println("   2. Зависимости (Dependencies) - как связи между вагонами")
	fmt.Println("   3. Пакеты (Packages) - как отделения в вагонах")
	fmt.Println("   4. Импорты (Imports) - как билеты для перемещения между вагонами")
	fmt.Println("\n   Каждый пакет выполняет свою роль:")
	fmt.Println("   • pkg/delivery - логика доставки")
	fmt.Println("   • pkg/train - логика поездки")
	fmt.Println("   • pkg/book - работа с книгами")
	fmt.Println("   • pkg/motivation - мотивационная система")
	fmt.Println("   • internal/calculator - внутренняя логика подсчёта")
	fmt.Println("\n   Вместе они создают целостную систему!")
	fmt.Println(strings.Repeat("=", 70))

	// Финальная мотивация
	fmt.Println("\n🚀 ФИНАЛЬНЫЙ СТИМУЛ:")
	fmt.Println("   Нужно продолжать изучение, нужно!")
	fmt.Println("   Не отвлекайся на видеомонтаж!")
	fmt.Println("   Потом на выходных будешь монтировать свои видео,")
	fmt.Println("   когда устроишься на нормальную работу в офис!")
	fmt.Println(strings.Repeat("=", 70))
}
