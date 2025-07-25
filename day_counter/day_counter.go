package main

import (
	"fmt"
	"time"
)

func main() {
	// Стартовая дата марафона (25.07.2025)
	startDate := time.Date(2025, time.July, 25, 0, 0, 0, 0, time.UTC)
	currentDate := time.Now()

	// Вычисляем сколько дней прошло
	daysPassed := int(currentDate.Sub(startDate).Hours() / 24)

	// Проверяем, не начался ли марафон сегодня
	if daysPassed < 0 {
		fmt.Println("Марафон 100 дней Go начнётся", startDate.Format("02.01.2006"))
	} else {
		// Выводим результат с мотивационным сообщением
		fmt.Printf("Сегодня: %s\n", currentDate.Format("02.01.2006"))
		fmt.Printf("Дней в марафоне: %d/100\n", daysPassed+1)
		fmt.Println("-------------------------------")

		switch {
		case daysPassed == 0:
			fmt.Println("Старт дан! Сегодня 25.7.25 - начало твоего 100-дневного айти-марафона! Теперь твой путь — Go.")
		case daysPassed < 7:
			fmt.Println("Формируешь привычку. Каждый день — новый коммит в GitHub!")
		case daysPassed < 30:
			fmt.Println("Создал первый проект? Уже видишь структуры вместо снов? Так держать!")
		case daysPassed == 49:
			fmt.Println("Половина пути! Теперь ты не новичок, а junior с опытом.")
		case daysPassed == 99:
			fmt.Println("Финишная прямая! Готовь резюме — пора штурмовать вакансии.")
		case daysPassed >= 100:
			fmt.Println("Ты завершил марафон! Теперь Go — твой язык, а IT — твой дом.")
		default:
			fmt.Println("Идеальный день, чтобы изучить что-то новое в Go!")
		}
	}
}
