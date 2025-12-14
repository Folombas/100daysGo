package examples

import "fmt"

// Функция без параметров и возвращаемых значений
func Greet() {
	fmt.Println("Привет от функции Greet!")
}

// Функция с параметрами
func CalculateDeliveryPrice(distance float64, city string) float64 {
	basePrice := 2.5
	pricePerKm := 0.8

	if city == "Москва" {
		pricePerKm = 1.2
	}

	return basePrice + distance*pricePerKm
}

// Функция с несколькими возвращаемыми значениями
func GetDeliveryStats(deliveries int) (int, float64, string) {
	avgTime := 25.5
	rating := "отличный"

	if deliveries < 10 {
		rating = "хороший"
	} else if deliveries > 50 {
		rating = "превосходный"
	}

	return deliveries, avgTime, rating
}

// Анонимная функция
var WelcomeMessage = func(name string) string {
	return "Добро пожаловать, " + name + "! Пусть твой код будет чистым, а баги — редкими!"
}

// Демонстрация базовых функций
func DemonstrateBasicFunctions() {
	// Вызов функции без параметров
	Greet()

	// Вызов функции с параметрами
	price := CalculateDeliveryPrice(5.5, "Москва")
	fmt.Printf("💰 Стоимость доставки для 5.5 км в Москве: %.2f руб.\n", price)

	// Вызов функции с несколькими возвращаемыми значениями
	count, avgTime, rating := GetDeliveryStats(35)
	fmt.Printf("📊 Статистика: %d доставок, среднее время: %.1f мин, рейтинг: %s\n",
		count, avgTime, rating)

	// Вызов анонимной функции
	message := WelcomeMessage("будущий Go-разработчик")
	fmt.Println("💌 " + message)
}
