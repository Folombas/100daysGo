package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Balance float64
	IsAdmin bool
	Skills  []string
	Details map[string]interface{}
}

func main() {
	fmt.Println("🚀 Day6: Zero Values Exploration")
	fmt.Println("=================================")

	// Демонстрация нулевых значений для базовых типов
	var integer int
	var decimal float64
	var text string
	var flag bool
	var slice []string
	var mapping map[string]int
	var pointer *int

	fmt.Println("\n📊 Zero Values для базовых типов:")
	fmt.Printf("int: %d\n", integer)
	fmt.Printf("float64: %f\n", decimal)
	fmt.Printf("string: '%s'\n", text)
	fmt.Printf("bool: %t\n", flag)
	fmt.Printf("slice: %v (nil: %t)\n", slice, slice == nil)
	fmt.Printf("map: %v (nil: %t)\n", mapping, mapping == nil)
	fmt.Printf("pointer: %v (nil: %t)\n", pointer, pointer == nil)

	// Демонстрация для структуры
	fmt.Println("\n👤 Zero Values для структуры Person:")
	var person Person
	fmt.Printf("Name: '%s'\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Balance: %.2f\n", person.Balance)
	fmt.Printf("IsAdmin: %t\n", person.IsAdmin)
	fmt.Printf("Skills: %v (nil: %t)\n", person.Skills, person.Skills == nil)
	fmt.Printf("Details: %v (nil: %t)\n", person.Details, person.Details == nil)

	// Практическое применение
	fmt.Println("\n💡 Практический пример:")
	users := make([]Person, 3) // Создаем срез из 3 Person
	for i, user := range users {
		fmt.Printf("User%d: {Name:'%s', Age:%d, Balance:%.2f}\n",
			i+1, user.Name, user.Age, user.Balance)
	}

	fmt.Println("\n🎯 Вывод: Go каждую переменную инициализирует нулевым значением!")
	fmt.Println("Это делает код безопаснее и предсказуемее ❤️")
}
