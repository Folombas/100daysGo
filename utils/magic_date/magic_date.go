package main

import (
	"fmt"
	"time"
)

func main() {
	// Магическая дата Go
	magicDate := time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)

	// Старт марафона Гоши
	marathonStart := time.Date(2025, 7, 25, 0, 0, 0, 0, time.UTC)

	// Вычисляем разницу
	diff := marathonStart.Sub(magicDate)
	daysBetween := int(diff.Hours() / 24)
	years := daysBetween / 365

	// Визуальное представление
	fmt.Println("🔥 Тайная связь дат в Go")
	fmt.Println("════════════════════════════════")
	fmt.Println("Магическая дата Go: 02.01.2006")
	fmt.Println("Старт 100-дневного айти-марафона Гоши:    25.07.2025")
	fmt.Println("────────────────────────────────")
	fmt.Printf("Между ними: %d дней\n", daysBetween)
	fmt.Printf("Это ровно %d лет и %d дней\n", years, daysBetween-years*365)
	fmt.Println("════════════════════════════════")

	// Философская связь
	fmt.Println("Это означает:")
	fmt.Println("1. Ты начал обучение через", years, "лет после установления стандарта дат в Go")
	fmt.Println("2. Каждая дата в твоих программах будет использовать этот магический шаблон")
	fmt.Println("3. Ты становишься частью истории языка программирования, который выбрал для фокусированного и продвинутого обучения.")

	// Практический урок
	fmt.Println("\n💡 Запомни навсегда:")
	fmt.Println("В Go форматирование дат всегда использует шаблон:")
	fmt.Println("   currentTime.Format(\"02.01.2006\")")
	fmt.Println("где 02=день, 01=месяц, 2006=год")

	// ASCII-арт
	fmt.Println(`
    02.01.2006         25.07.2025
		|                    |
		│    19 лет          │
		╰───────────────┬────╯
						│
				Твой код на Go`)
}
