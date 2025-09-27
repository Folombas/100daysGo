package main

import (
	"fmt"
)

// 🎯 1. BASIC FUNCTION
func greet() {
	fmt.Println("👋 Привет, мир!")
}

// 🎯 2. FUNCTION WITH PARAMETERS AND RETURN
func add(a int, b int) int {
	return a + b
}

// 🎯 3. MULTIPLE RETURN VALUES
func calculator(a, b float64) (sum float64, product float64) {
	sum = a + b
	product = a * b
	return
}

// 🎯 4. VARIADIC FUNCTION
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 🎯 5. FUNCTION AS VALUE
var square = func(x int) int {
	return x * x
}

// 🎯 6. CLOSURE
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 🎯 7. FUNCTION WITH DEFER
func workWithFile(filename string) {
	fmt.Printf("📁 Открываем файл: %s\n", filename)
	defer fmt.Printf("✅ Закрываем файл: %s\n", filename)
	fmt.Println("✍️  Работаем с файлом...")
}

// 🎯 8. METHOD
type User struct {
	Name string
	Age  int
}

func (u User) Introduce() string {
	return fmt.Sprintf("👋 Я %s, мне %d лет", u.Name, u.Age)
}

func main() {
	fmt.Println("🎉 ДЕМОНСТРАЦИЯ ФУНКЦИЙ В GO!")
	fmt.Println("=============================")

	// 1. Basic function
	greet()

	// 2. Function with parameters
	result := add(5, 3)
	fmt.Printf("➕ 5 + 3 = %d\n", result)

	// 3. Multiple returns
	sum, product := calculator(4, 5)
	fmt.Printf("🧮 Сумма: %.1f, Произведение: %.1f\n", sum, product)

	// 4. Variadic function
	total := sumNumbers(1, 2, 3, 4, 5)
	fmt.Printf("🔢 Сумма чисел: %d\n", total)

	// 5. Function as value
	fmt.Printf("🔲 Квадрат 7: %d\n", square(7))

	// 6. Closure
	counter := createCounter()
	fmt.Printf("🔢 Счетчик: %d, %d, %d\n", counter(), counter(), counter())

	// 7. Defer
	workWithFile("документ.txt")

	// 8. Method
	user := User{Name: "Саша", Age: 25}
	fmt.Println(user.Introduce())

	fmt.Println("\n🎯 Основные принципы функций в Go:")
	principles := []string{
		"✅ Объявляются с помощью func",
		"✅ Могут возвращать несколько значений",
		"✅ Поддерживают variadic параметры",
		"✅ Являются first-class гражданами",
		"✅ Поддерживают замыкания",
		"✅ Используют defer для отложенного выполнения",
	}

	for _, principle := range principles {
		fmt.Println(principle)
	}
}
