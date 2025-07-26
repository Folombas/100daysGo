package main

import "fmt"

func main() {
	fmt.Println("Это второй день ежедневного регулярного обучения программированию на Go!")

	var (
		deliveryCount int     = 6
		kmTotal       float64 = 250.5
		isExhausted   bool    = true
	)

	fmt.Printf("Заказов: %d\n", deliveryCount)
	fmt.Printf("Километраж: %.1f км\n", kmTotal)
	fmt.Printf("Усталось: %t\n", isExhausted)

	// Мои данные за сегодня
	earnings := 1400.0 // заработал рублей
	food := 200.0      // купил еды: ватрушка с творогом, пирожок с малиной и Айран 1% 0,5 л.
	bikeRepair := 0.0  // Сегодня не тратил деньги на ремонт велика

	// Вычисления
	netIncome := earnings - (food - bikeRepair)
	hourlyRate := netIncome / 6 // 6 часов работы

	// Вывод
	fmt.Println("═════════════════════════════")
	fmt.Printf(" Заработано: \t%.2f руб\n", earnings)
	fmt.Printf(" Расходы: \t%.2f руб\n", food+bikeRepair)
	fmt.Printf(" Чистый доход: \t%.2f руб\n", netIncome)
	fmt.Println("──────────────────────────────")
	fmt.Printf(" Почасовая ставка: %.2f руб/час\n", hourlyRate)
	fmt.Println("═════════════════════════════")
	fmt.Println("Цель: чтобы эта цифра была > 500 руб/час")
}
