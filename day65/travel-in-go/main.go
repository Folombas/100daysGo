package main

import (
	"fmt"
	"strings"
	"time"
)

// Experience - интерфейс для любого опыта (путешествие или учеба)
type Experience interface {
	Describe() string
	GetValue() float64
}

// TravelPhoto - фото из путешествия (соцсети)
type TravelPhoto struct {
	Location string
	Likes    int
	Cost     float64
}

func (tp TravelPhoto) Describe() string {
	return fmt.Sprintf("Фото из %s (%d лайков)", tp.Location, tp.Likes)
}

func (tp TravelPhoto) GetValue() float64 {
	return -tp.Cost // Траты на поездку
}

// GoChallenge - задание по Go
type GoChallenge struct {
	Topic      string
	HoursSpent float64
	Completed  bool
}

func (gc GoChallenge) Describe() string {
	status := "в процессе"
	if gc.Completed {
		status = "завершено"
	}
	return fmt.Sprintf("Задание по Go: %s (%s, %.1f ч.)", gc.Topic, status, gc.HoursSpent)
}

func (gc GoChallenge) GetValue() float64 {
	value := gc.HoursSpent * 1000 // Каждый час Go стоит 1000 у.е. будущего дохода
	if gc.Completed {
		value *= 2
	}
	return value
}

// DailyRoutine - ежедневная рутина Гоши
type DailyRoutine struct {
	Activity     string
	Location     string
	Satisfaction int
}

func (dr DailyRoutine) Describe() string {
	return fmt.Sprintf("%s в %s (удовлетворение: %d/10)", dr.Activity, dr.Location, dr.Satisfaction)
}

func (dr DailyRoutine) GetValue() float64 {
	return float64(dr.Satisfaction) * 50
}

// TravelDream - мечта о путешествии
type TravelDream struct {
	Destination string
	Price       float64
	Saved       float64
}

func (td TravelDream) Describe() string {
	progress := td.Saved / td.Price * 100
	return fmt.Sprintf("Мечта: %s (накоплено %.1f%%)", td.Destination, progress)
}

func (td TravelDream) GetValue() float64 {
	return td.Saved * 10 // Мотивационная ценность
}

// Type Switch обработчик
func processExperience(exp Experience) {
	fmt.Println("\n🔍 Анализирую опыт...")
	fmt.Printf("   Описание: %s\n", exp.Describe())

	// TYPE SWITCH - здесь мы определяем конкретный тип
	switch v := exp.(type) {
	case TravelPhoto:
		fmt.Println("   🚫 Обнаружено: Фото из соцсетей")
		fmt.Println("   💡 Совет Гоше: Зависть к чужим путешествиям не принесет тебе билеты")
		fmt.Printf("   💸 Энергозатраты: %.0f у.е.\n", -v.GetValue())

	case GoChallenge:
		fmt.Println("   ✅ Обнаружено: Изучение Go")
		fmt.Println("   💡 Совет Гоше: Каждый час кода приближает тебя к работе в айти")
		fmt.Printf("   💰 Будущая ценность: %.0f у.е.\n", v.GetValue())
		if !v.Completed {
			fmt.Println("   ⚡ Доделай задание, и ценность удвоится!")
		}

	case DailyRoutine:
		fmt.Println("   🏡 Обнаружено: Ежедневная рутина")
		if v.Satisfaction > 7 {
			fmt.Println("   👍 Отличный день! Ты на правильном пути")
		} else {
			fmt.Println("   💪 Не унывай! Завтра будет лучше")
		}

	case TravelDream:
		fmt.Println("   🌍 Обнаружено: Мечта о путешествии")
		progress := v.Saved / v.Price * 100
		if progress > 50 {
			fmt.Println("   🎯 Уже больше половины! Продолжай в том же духе")
		} else if progress > 0 {
			fmt.Printf("   📈 Прогресс: %.1f%%. Каждый день Go — шаг к мечте\n", progress)
		} else {
			fmt.Println("   💡 Начни с малого. Первый коммит — первый шаг")
		}

	default:
		fmt.Println("   ❓ Неизвестный тип опыта")
	}

	// Type assertion с проверкой (альтернативный способ)
	fmt.Print("\n   Проверка через assertion: ")
	if tp, ok := exp.(TravelPhoto); ok {
		fmt.Printf("Это фото из %s, не корми зависть!\n", tp.Location)
	} else if _, ok := exp.(GoChallenge); ok {
		fmt.Println("Это Go! Самый ценный актив!")
	} else {
		fmt.Println("Ценный опыт, продолжай!")
	}
}

// Путешествие по типам данных (вместо реальных путешествий)
func typeSwitchJourney() {
	fmt.Println("================================")
	fmt.Println("   TRAVEL IN GO: Type Switch Journey")
	fmt.Println("   День 65: Вместо Беловежской Пущи — мир интерфейсов")
	fmt.Println("================================")

	// Создаем слайс разных "опытов" Гоши
	experiences := []Experience{
		TravelPhoto{Location: "Беловежская Пуща", Likes: 150, Cost: 50000},
		GoChallenge{Topic: "Type Switch", HoursSpent: 3.5, Completed: true},
		DailyRoutine{Activity: "Изучение Go", Location: "Химки", Satisfaction: 9},
		TravelDream{Destination: "Зимовка на Бали с MacBook", Price: 500000, Saved: 75000},
		GoChallenge{Topic: "Интерфейсы", HoursSpent: 2.0, Completed: false},
		TravelPhoto{Location: "Курт", Likes: 87, Cost: 15000},
		DailyRoutine{Activity: "Уборка", Location: "Кухня", Satisfaction: 6},
	}

	// Счетчики для статистики
	var totalValue float64
	var goChallenges, travelPhotos int

	for i, exp := range experiences {
		fmt.Printf("\n[Опыт %d/%d]\n", i+1, len(experiences))
		processExperience(exp)
		totalValue += exp.GetValue()

		// Считаем типы для статистики
		switch exp.(type) {
		case GoChallenge:
			goChallenges++
		case TravelPhoto:
			travelPhotos++
		}
	}

	// Итоги дня
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📊 ИТОГИ ПУТЕШЕСТВИЯ ПО ТИПАМ:")
	fmt.Printf("   Всего опытов: %d\n", len(experiences))
	fmt.Printf("   Заданий по Go: %d\n", goChallenges)
	fmt.Printf("   Фото из путешествий: %d\n", travelPhotos)
	fmt.Printf("   Общая ценность дня: %.0f у.е.\n", totalValue)

	if totalValue > 0 {
		fmt.Println("\n💡 ВЫВОД ДНЯ:")
		fmt.Println("   Гоша, твои часы с Go стоят больше,")
		fmt.Println("   чем все лайки под фото с зубрами из Беловежской Пущи!")
		fmt.Println("   Каждый type switch — это твой шаг к мечте.")
	} else {
		fmt.Println("\n⚠️  Внимание: слишком много времени на соцсети!")
	}
}

func main() {
	// Легенда дня
	fmt.Println("🌅 УТРО 07.01.2026:")
	fmt.Println("7:00 - Будильник на Honor 10x Lite")
	fmt.Println("7:05 - Теплый душ, бритьё")
	fmt.Println("7:30 - Цикорий с медом, сухарики с изюмом")
	fmt.Println("8:00 - Уборка, помощь матушке")
	fmt.Println("9:00 - Включаю комп...")

	time.Sleep(1 * time.Second)

	// Вместо соцсетей - путешествие по типам
	fmt.Println("\n📱 Вместо соцсетей (фото зубров в Беловежской Пуще)...")
	time.Sleep(1 * time.Second)
	fmt.Println("💡 Решаю: лучше изучу Type Switch в Go!")

	typeSwitchJourney()

	// Мотивационное сообщение
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("💌 ПИСЬМО СЕБЕ ИЗ БУДУЩЕГО:")
	fmt.Println("   'Привет, Гоша из 2026 года!")
	fmt.Println("   Ты помнишь, как завидовал фото из Беловежской Пущи?")
	fmt.Println("   А теперь я пишу этот код с берега тёплого Индийского океана.")
	fmt.Println("   Спасибо тебе за каждый type switch,")
	fmt.Println("   за каждый час с Go вместо соцсетей.")
	fmt.Println("   Продолжай. Оно того стоит.'")

	fmt.Println("\n🚀 GO-ПУТЕШЕСТВИЕ ПРОДОЛЖАЕТСЯ!")
}
