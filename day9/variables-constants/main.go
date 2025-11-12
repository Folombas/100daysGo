package main

import (
	"fmt"
	"math"
)

func main() {
	// Примеры переменных
	fmt.Println("=== Переменные в Go ===")

	// Объявление переменных с явным типом
	var age int = 37
	var name string = "Гоша"
	var height float64 = 1.82
	var isActive bool = true

	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Возраст: %d\n", age)
	fmt.Printf("Рост: %.2fм\n", height)
	fmt.Printf("Студент: %t\n", isActive)

	// Объявление переменных с неявным типом
	year := 2025
	temp := 36.6
	city := "Москва"
	isStudent := false

	fmt.Printf("\nГод: %d\n", year)
	fmt.Printf("Температура: %.1f°C\n", temp)
	fmt.Printf("Город: %s\n", city)
	fmt.Printf("Студент: %t\n", isStudent)

	// Объявление нескольких переменных одновременно
	var (
		temperature = 22.5
		humidity    = 65
		pressure    = 1013
	)

	fmt.Printf("\nТемпература: %.1f°C\n", temperature)
	fmt.Printf("Влажность: %d%%\n", humidity)
	fmt.Printf("Давление: %d гПа\n", pressure)

	// Примеры констант
	fmt.Println("\n=== Константы в Go ===")

	const pi = 3.14159
	const appName = "MyGoApp"
	const maxUsers = 1000

	fmt.Printf("Число π: %.5f\n", pi)
	fmt.Printf("Название приложения: %s\n", appName)
	fmt.Printf("Максимальное количество пользователей: %d\n", maxUsers)

	// Константы с явным типом
	const (
		minAge int = 18
		maxAge int = 120
	)

	fmt.Printf("Минимальный возраст: %d\n", minAge)
	fmt.Printf("Максимальный возраст: %d\n", maxAge)

	// Константы с использованием math
	const (
		eulersNumber = math.E
		sqrt2        = math.Sqrt2
	)

	fmt.Printf("Число Эйлера: %.5f\n", eulersNumber)
	fmt.Printf("Квадратный корень из 2: %.5f\n", sqrt2)

	// Примеры типов переменных
	fmt.Println("\n=== Типы переменных ===")

	var intVar int = 42
	var int8Var int8 = 127
	var int16Var int16 = 32767
	var int32Var int32 = 2147483647
	var int64Var int64 = 9223372036854775807

	var uintVar uint = 42
	var uint8Var uint8 = 255
	var uint16Var uint16 = 65535
	var uint32Var uint32 = 4294967295
	var uint64Var uint64 = 18446744073709551615

	var float32Var float32 = 3.14159
	var float64Var float64 = 3.14159265359

	var complex64Var complex64 = 1 + 2i
	var complex128Var complex128 = 1 + 2i

	var runeVar rune = 'A'
	var byteVar byte = 'B'

	fmt.Printf("int: %d\n", intVar)
	fmt.Printf("int8: %d\n", int8Var)
	fmt.Printf("int16: %d\n", int16Var)
	fmt.Printf("int32: %d\n", int32Var)
	fmt.Printf("int64: %d\n", int64Var)

	fmt.Printf("uint: %d\n", uintVar)
	fmt.Printf("uint8: %d\n", uint8Var)
	fmt.Printf("uint16: %d\n", uint16Var)
	fmt.Printf("uint32: %d\n", uint32Var)
	fmt.Printf("uint64: %d\n", uint64Var)

	fmt.Printf("float32: %.5f\n", float32Var)
	fmt.Printf("float64: %.10f\n", float64Var)

	fmt.Printf("complex64: %v\n", complex64Var)
	fmt.Printf("complex128: %v\n", complex128Var)

	fmt.Printf("rune: %c\n", runeVar)
	fmt.Printf("byte: %c\n", byteVar)

	// Примеры присваивания и изменения значений
	fmt.Println("\n=== Изменение значений переменных ===")

	var counter int = 0
	fmt.Printf("Счетчик до: %d\n", counter)
	counter = 10
	fmt.Printf("Счетчик после присваивания: %d\n", counter)
	counter += 5
	fmt.Printf("Счетчик после увеличения на 5: %d\n", counter)

	// Обмен значений переменных
	a := 10
	b := 20
	fmt.Printf("До обмена: a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("После обмена: a=%d, b=%d\n", a, b)
}
