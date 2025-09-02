package main

import (
	"fmt"
	"time"
)

func main() {
	startDate := time.Date(2025, time.July, 25, 0, 0, 0, 0, time.UTC)
	currentDate := time.Now()
	daysPassed := int(currentDate.Sub(startDate).Hours() / 24)
	currentDay := daysPassed + 1

	// Zero-based индексация
	days := []string{"Day0", "Day1", "Day2", "Day3", "Day4", "Day5", "Day6", "Day7"}

	var dayLabel string
	if daysPassed < len(days) {
		dayLabel = days[daysPassed]
	} else {
		dayLabel = fmt.Sprintf("Day%d", daysPassed)
	}

	if daysPassed < 0 {
		fmt.Println("Марафон 100 дней Go начнётся", startDate.Format("02.01.2006"))
	} else {
		fmt.Printf("Сегодня: %s\n", currentDate.Format("02.01.2006"))
		fmt.Printf("День марафона: %s (%d/100)\n", dayLabel, currentDay)
		fmt.Println("-------------------------------")

		switch {
		case daysPassed == 0:
			fmt.Printf("%s: Старт дан! Твой путь в Go начинается.\n", dayLabel)
		case daysPassed < 7:
			fmt.Printf("%s: Формируешь привычку. Код -> Коммит -> Push!\n", dayLabel)
		case daysPassed < 30:
			fmt.Printf("%s: Прогресс ощущается? Продолжай в том же духе!\n", dayLabel)
		case daysPassed == 49:
			fmt.Printf("%s: Полпути пройдено! Ты теперь Junior Go-разработчик с опытом.\n", dayLabel)
		case daysPassed == 99:
			fmt.Printf("%s: Финиш рядом! Готовь резюме и начни откликаться на вакансии.\n", dayLabel)
		case daysPassed >= 100:
			fmt.Printf("%s: Марафон завершён! Добро пожаловать в IT.\n", dayLabel)
		default:
			fmt.Printf("%s: Идеальный день для изучения чего-то ещё в Go!\n", dayLabel)
		}
	}
	fmt.Println("Сегодня днём 2 сентября 2025 года мы изучаем Terminal: Создаем консольные утилиты на Go.")
}
