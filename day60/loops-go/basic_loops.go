package main

import (
	"fmt"
	"time"
)

func demoBasicLoops() {
	printSeparator()

	fmt.Println("🔄 Классический for цикл (как в C):")
	for i := 0; i < 5; i++ {
		fmt.Printf("🚀 Итерация %d\n", i)
	}

	fmt.Println("\n🌀 While-подобный цикл (только условие):")
	counter := 3
	for counter > 0 {
		fmt.Printf("⏳ Осталось: %d\n", counter)
		counter--
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("🎉 Обратный отсчет завершен!")

	fmt.Println("\n♾️ Бесконечный цикл (с контролируемым выходом):")
	attempts := 0
	for {
		attempts++
		fmt.Printf("🔍 Попытка №%d\n", attempts)
		if attempts >= 3 {
			fmt.Println("✅ Найдено решение!")
			break
		}
	}

	fmt.Println("\n🎯 Цикл с continue и break:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Пропускаем четные числа
		}
		if i > 7 {
			break // Выходим при i > 7
		}
		fmt.Printf("🔢 Нечетное число: %d\n", i)
	}
}
