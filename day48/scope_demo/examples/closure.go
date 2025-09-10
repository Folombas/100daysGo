package examples

import "fmt"

// DemoClosures демонстрирует замыкания и их область видимости
func DemoClosures() {
	// Переменная уровня функции
	counter := 0
	message := "Исходное сообщение"
	
	fmt.Println("🔄 Начальное значение counter:", counter)
	fmt.Println("💬 Начальное сообщение:", message)
	
	// Создаем замыкание, которое имеет доступ к переменным функции
	increment := func() {
		counter++ // Изменяем переменную внешней функции
		message = "Изменено замыканием"
		fmt.Println("   📦 Замыкание: counter =", counter, "message =", message)
	}
	
	// Вызываем замыкание несколько раз
	increment()
	increment()
	increment()
	
	fmt.Println("🔄 Конечное значение counter:", counter)
	fmt.Println("💬 Конечное сообщение:", message)
	
	// Замыкание с параметрами
	multiplier := createMultiplier(2)
	fmt.Println("✖️  Умножитель 2 * 5 =", multiplier(5))
	fmt.Println("✖️  Умножитель 2 * 10 =", multiplier(10))
}

// createMultiplier создает и возвращает замыкание
func createMultiplier(factor int) func(int) int {
	// factor "захватывается" замыканием
	return func(x int) int {
		return x * factor
	}
}