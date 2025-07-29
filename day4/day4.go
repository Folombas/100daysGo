package main

import (
	"fmt"
	"time"
)

func main() {
	// Получаем текущую дату
	now := time.Now()

	// Словарь для перевода месяцев на русский (в родительном падеже)
	months := map[time.Month]string{
		time.January:   "января",
		time.February:  "февраля",
		time.March:     "марта",
		time.April:     "апреля",
		time.May:       "мая",
		time.June:      "июня",
		time.July:      "июля",
		time.August:    "августа",
		time.September: "сентября",
		time.October:   "октября",
		time.November:  "ноября",
		time.December:  "декабря",
	}

	// Извлекаем компоненты даты
	day := now.Day()
	month := months[now.Month()]
	year := now.Year()

	// Форматируем вывод согласно примеру
	fmt.Printf("Сегодня %d %s %d г.\n", day, month, year)
	fmt.Println("This is Day 4 of 100 days of daily Go coding!")
}
