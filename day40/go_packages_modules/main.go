package main

import (
	"fmt"

	"go_packages_modules/calculator"
	"go_packages_modules/greeter"
	"go_packages_modules/internal/tools"
	"go_packages_modules/quotes"
)

func main() {
	fmt.Println("📦 День 40: Пакеты и модули в Go")
	fmt.Println("======================================")

	// Демонстрация пакета greeter
	tools.PrintInfo("Демонстрация пакета greeter")
	greeting := greeter.Greet("Гоша")
	fmt.Println(greeting)
	fmt.Println()

	// Демонстрация пакета calculator
	tools.PrintInfo("Демонстрация пакета calculator")

	a, b := 15.0, 3.0
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, calculator.Add(a, b))
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, calculator.Subtract(a, b))
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, calculator.Multiply(a, b))

	result, err := calculator.Divide(a, b)
	if err != nil {
		tools.PrintError(err.Error())
	} else {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
	}

	// Проверка на простые числа
	tools.PrintInfo("Проверка чисел на простоту")
	numbers := []int{2, 7, 10, 17, 25}
	for _, num := range numbers {
		if calculator.IsPrime(num) {
			fmt.Printf("%d - простое число\n", num)
		} else {
			fmt.Printf("%d - составное число\n", num)
		}
	}
	fmt.Println()

	// Демонстрация пакета quotes
	tools.PrintInfo("Демонстрация пакета quotes")
	fmt.Println("Случайная мотивационная цитата:")
	quote := quotes.GetRandomQuote()
	fmt.Println(quote.String())
	fmt.Println()

	// Показать цитаты из категории "Программирование"
	tools.PrintInfo("Цитаты о программировании")
	programmingQuotes := quotes.GetQuotesByCategory("Программирование")
	for i, q := range programmingQuotes {
		fmt.Printf("%d. %s\n— %s\n\n", i+1, q.Text, q.Author)
	}

	tools.PrintSuccess("Все пакеты успешно продемонстрированы!")
	fmt.Println("\n🎉 Вы освоили работу с пакетами и модулями в Go!")
}
