package main

import "fmt"

// Демонстрация структур данных
func demoDataStructures() {
	fmt.Println("Структуры данных:")

	// Структуры (очень похоже на C)
	type Person struct {
		Name    string
		Age     int
		Address string
	}

	// Инициализация структур (похоже на C)
	person1 := Person{"Иван Иванов", 30, "Москва"}
	person2 := Person{
		Name:    "Мария Петрова",
		Age:     25,
		Address: "Санкт-Петербург",
	}

	fmt.Printf("Person1: %+v\n", person1)
	fmt.Printf("Person2: %+v\n", person2)

	// Массивы (похожи на C, но с дополнительными возможностями)
	fmt.Println("\nМассивы:")
	var arr [5]int // Как в C: int arr[5];
	for i := 0; i < len(arr); i++ {
		arr[i] = i * i
	}
	fmt.Printf("Массив: %v\n", arr)

	// Срезы (улучшение относительно C-массивов)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Срез: %v, длина: %d, емкость: %d\n",
		slice, len(slice), cap(slice))

	// Указатели на структуры (как в C)
	pPerson := &person1
	fmt.Printf("Имя через указатель: %s\n", pPerson.Name)
}

// Константы и перечисления (похоже на C)
const (
	Red   = iota // 0
	Green        // 1
	Blue         // 2
)

const (
	ReadPermission  = 1 << iota // 1
	WritePermission             // 2
	ExecutePermission           // 4
)
