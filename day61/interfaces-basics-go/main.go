// day61/interfaces-basics-go/main.go
package main

import (
	"fmt"
	"time"
)

// 🎯 День 61: Основы интерфейсов в Go
// 2026: ГОД ПОГРУЖЕНИЯ В GO | УРОВЕНЬ ГЛУБИНЫ: 1

func main() {
	fmt.Println("🐹 2026: ПОГРУЖЕНИЕ В GO | День 61/100")
	fmt.Println("🎯 Тема: Основы интерфейсов")
	fmt.Println("📅", time.Now().Format("02.01.2006"))
	fmt.Println("========================================")

	// Демонстрация основных концепций интерфейсов
	demoBasicInterface()
	demoInterfaceComposition()
	demoEmptyInterface()
	demoTypeAssertion()
	demoStringerInterface()
}

// 🔹 1. Базовый интерфейс
type Speaker interface {
	Speak() string
	Volume() int
}

type Human struct {
	Name   string
	Age    int
	Energy int
}

func (h Human) Speak() string {
	return fmt.Sprintf("Привет, я %s, мне %d лет", h.Name, h.Age)
}

func (h Human) Volume() int {
	return h.Energy / 10
}

// 🔹 2. Композиция интерфейсов
type Walker interface {
	Walk() string
}

type Runner interface {
	Run() string
}

type Athlete interface {
	Speaker
	Walker
	Runner
}

// 🔹 3. Пустой интерфейс
type UniversalContainer struct {
	Data interface{}
}

func demoBasicInterface() {
	fmt.Println("\n🔸 1. Базовый интерфейс:")

	gosha := Human{Name: "Гоша", Age: 38, Energy: 100}
	var speaker Speaker = gosha

	fmt.Printf("Говорит: %s\n", speaker.Speak())
	fmt.Printf("Громкость: %d%%\n", speaker.Volume())
}

func demoInterfaceComposition() {
	fmt.Println("\n🔸 2. Композиция интерфейсов:")

	// Пример полиморфизма
	var speakers []Speaker = []Speaker{
		Human{Name: "Гоша", Age: 38, Energy: 100},
		Human{Name: "Наставник", Age: 35, Energy: 90},
	}

	for _, s := range speakers {
		fmt.Printf("- %s (громкость: %d%%)\n", s.Speak(), s.Volume())
	}
}

func demoEmptyInterface() {
	fmt.Println("\n🔸 3. Пустой интерфейс:")

	container := UniversalContainer{}
	container.Data = "Строка данных"
	fmt.Printf("Тип: %T, Значение: %v\n", container.Data, container.Data)

	container.Data = 42
	fmt.Printf("Тип: %T, Значение: %v\n", container.Data, container.Data)

	container.Data = []string{"Go", "Интерфейсы", "Погружение"}
	fmt.Printf("Тип: %T, Значение: %v\n", container.Data, container.Data)
}

func demoTypeAssertion() {
	fmt.Println("\n🔸 4. Утверждение типа (Type Assertion):")

	var something interface{} = "Это строка"

	if str, ok := something.(string); ok {
		fmt.Printf("Успешно: %s (длина: %d)\n", str, len(str))
	}

	// Type switch
	processValue(42)
	processValue("Глубина вместо ширины")
	processValue(3.14)
}

func processValue(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("Целое число: %d\n", val)
	case string:
		fmt.Printf("Строка: %s\n", val)
	default:
		fmt.Printf("Неизвестный тип: %T\n", val)
	}
}

func demoStringerInterface() {
	fmt.Println("\n🔸 5. Интерфейс Stringer (из стандартной библиотеки):")

	person := Human{Name: "Гофер", Age: 38, Energy: 85}
	fmt.Println(person) // Автоматически вызовется String(), если он есть
}

// Реализуем Stringer для Human
func (h Human) String() string {
	return fmt.Sprintf("[Человек: %s, Возраст: %d, Энергия: %d%%]",
		h.Name, h.Age, h.Energy)
}

// 🔹 Утилитарная функция для проверки интерфейса
func checkInterfaceImplementation() {
	fmt.Println("\n🔍 Проверка реализации интерфейса:")

	var _ Speaker = (*Human)(nil) // Компилятор проверит, что Human реализует Speaker

	fmt.Println("✅ Human корректно реализует интерфейс Speaker")
}
