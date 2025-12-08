package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

// Person - пример структуры с различными тегами JSON
type Person struct {
	Name        string `json:"name"`                       // Простой тег
	Age         int    `json:"age,omitempty"`              // omitempty - не выводить если нулевое значение
	Email       string `json:"email,omitempty"`            // omitempty работает и для строк
	SecretField string `json:"-"`                          // Минус - полностью игнорировать поле
	Salary      int    `json:"salary,string"`              // string - сериализовать как строку
	IsActive    bool   `json:"is_active"`                  // Снейк-кейс для JSON
	Nickname    string `json:"nickname,omitempty"`         // Будет опущено если пустое
	Rating      *float64 `json:"rating,omitempty"`         // Указатель для отличия отсутствия от нуля
}

// Product - структура с вложенными структурами
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Metadata struct {
		CreatedBy string `json:"created_by"`
		CreatedAt string `json:"created_at"`
	} `json:"metadata"`
}

func main() {
	fmt.Println("=== День 35: Теги структур и JSON ===")

	// Пример 1: Базовая сериализация
	fmt.Println("1. Базовая сериализация:")
	person1 := Person{
		Name:     "Гоша",
		Age:      30,
		Email:    "gosha@golang_gopher.com",
		Salary:   50000,
		IsActive: true,
		Rating:   nil, // nil указатель
	}

	// Маршалинг (Go структура -> JSON)
	jsonData, err := json.MarshalIndent(person1, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка маршалинга: %v\n", err)
		return
	}
	fmt.Printf("Структура -> JSON:\n%s\n\n", string(jsonData))

	// Пример 2: Десериализация
	fmt.Println("2. Десериализация:")
	jsonStr := `{
		"name": "Мария",
		"age": 25,
		"salary": "75000",
		"is_active": true,
		"nickname": "Маша"
	}`

	var person2 Person
	err = json.Unmarshal([]byte(jsonStr), &person2)
	if err != nil {
		fmt.Printf("Ошибка анмаршалинга: %v\n", err)
		return
	}
	fmt.Printf("JSON -> Структура:\n%+v\n\n", person2)

	// Пример 3: omitempty в действии
	fmt.Println("3. omitempty демонстрация:")
	person3 := Person{
		Name:        "Иван",
		Age:         0, // Будет опущено из-за omitempty
		Email:       "", // Будет опущено
		SecretField: "секрет",
		Salary:      0,
		IsActive:    false,
	}

	jsonData3, _ := json.MarshalIndent(person3, "", "  ")
	fmt.Printf("С omitempty:\n%s\n\n", string(jsonData3))

	// Пример 4: Сложная структура с вложенными объектами
	fmt.Println("4. Сложные структуры:")
	product := Product{
		ID:    1,
		Name:  "Ноутбук",
		Price: 999.99,
		Tags:  []string{"электроника", "техника"},
	}
	product.Metadata.CreatedBy = "admin"
	product.Metadata.CreatedAt = "2025-12-08"

	productJSON, _ := json.MarshalIndent(product, "", "  ")
	fmt.Printf("Продукт:\n%s\n\n", string(productJSON))

	// Пример 5: Динамическое чтение тегов
	fmt.Println("5. Чтение тегов через reflection:")
	t := reflect.TypeOf(person1)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		fmt.Printf("Поле: %-15s JSON тег: %s\n", field.Name, jsonTag)
	}

	// Пример 6: Запись в файл и чтение из файла
	fmt.Println("\n6. Работа с файлами:")

	// Запись в файл
	file, _ := os.Create("day35/data.json")
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(person1)
	file.Close()
	fmt.Println("Данные записаны в data.json")

	// Чтение из файла
	file, _ = os.Open("day35/data.json")
	var personFromFile Person
	decoder := json.NewDecoder(file)
	decoder.Decode(&personFromFile)
	file.Close()
	fmt.Printf("Прочитано из файла: %+v\n", personFromFile)
}
