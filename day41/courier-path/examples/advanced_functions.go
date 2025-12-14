package examples

import "fmt"

// Функция как параметр другой функции
func ApplyDiscount(originalPrice float64, discountFunc func(float64) float64) float64 {
	return discountFunc(originalPrice)
}

// Функция с неопределенным количеством параметров
func CalculateTotalRevenue(deliveries ...float64) float64 {
	total := 0.0
	for _, amount := range deliveries {
		total += amount
	}
	return total
}

// Рекурсивная функция
func CalculateCareerYears(year int) int {
	currentYear := 2025
	if year >= currentYear {
		return 0
	}
	return 1 + CalculateCareerYears(year+1)
}

// Функция с отложенным вызовом (defer)
func ProcessDelivery(deliveryID string) {
	fmt.Printf("📦 Начало обработки доставки #%s\n", deliveryID)
	defer fmt.Printf("✅ Завершение обработки доставки #%s\n", deliveryID)

	// Имитация процесса обработки
	for i := 1; i <= 3; i++ {
		fmt.Printf("  ⏳ Шаг %d для доставки #%s\n", i, deliveryID)
	}
}

// Демонстрация продвинутых функций
func DemonstrateAdvancedFunctions() {
	// Функция как параметр
	regularDiscount := func(price float64) float64 {
		return price * 0.95 // 5% скидка
	}

	specialDiscount := func(price float64) float64 {
		return price * 0.85 // 15% скидка
	}

	basePrice := 100.0
	fmt.Printf("🏷️  Обычная скидка для заказа %.0f руб: %.2f руб\n", basePrice, ApplyDiscount(basePrice, regularDiscount))
	fmt.Printf("🏷️  Специальная скидка для заказа %.0f руб: %.2f руб\n", basePrice, ApplyDiscount(basePrice, specialDiscount))

	// Функция с неопределенным количеством параметров
	total := CalculateTotalRevenue(250.50, 150.75, 300.00, 425.25)
	fmt.Printf("💰 Общий доход от доставок: %.2f руб\n", total)

	// Рекурсивная функция
	courierYears := CalculateCareerYears(2005)
	fmt.Printf("📈 Опыт работы курьером: %d лет\n", courierYears)

	// Функция с отложенным вызовом
	ProcessDelivery("DEL-2025-12-14")
}
