package main

import (
	"fmt"
	"reflect"
)

// Haystack представляет стог сена с различными предметами
type Haystack struct {
	Items map[string]interface{}
}

// NewHaystack создает новый стог сена
func NewHaystack() *Haystack {
	return &Haystack{
		Items: make(map[string]interface{}),
	}
}

// AddItem добавляет предмет в стог сена
func (h *Haystack) AddItem(name string, item interface{}) {
	h.Items[name] = item
}

// FindNeedle пытается найти иголку в стоге сена (Comma-Ok в действии!)
func (h *Haystack) FindNeedle() (string, bool) {
	// Проверяем наличие ключа в map
	if needle, ok := h.Items["needle"]; ok {
		if str, ok := needle.(string); ok {
			return str, true
		}
	}
	return "", false
}

// CountHayTypes подсчитывает различные типы сена
func (h *Haystack) CountHayTypes() {
	fmt.Println("🌾 Анализируем типы сена в стоге:")

	hayTypes := map[string]int{
		"timothy":   150,
		"clover":    75,
		"alfalfa":   200,
		"brome":     50,
	}

	// Comma-Ok для проверки существования ключа в map
	if count, exists := hayTypes["timothy"]; exists {
		fmt.Printf("✓ Тимофеевка: %d кг\n", count)
	}

	if count, exists := hayTypes["rye"]; exists {
		fmt.Printf("✓ Рожь: %d кг\n", count)
	} else {
		fmt.Println("✗ Рожь не найдена в стоге")
	}

	// Итерация по map с Comma-Ok
	totalWeight := 0
	for hayType, weight := range hayTypes {
		if weight > 0 {
			fmt.Printf("  - %s: %d кг\n", hayType, weight)
			totalWeight += weight
		}
	}
	fmt.Printf("📊 Общий вес сена: %d кг\n\n", totalWeight)
}

// FarmTools представляет инструменты фермера
type FarmTools struct {
	Tools map[string]interface{}
}

// CheckToolSafety проверяет безопасность инструментов
func CheckToolSafety(tools map[string]interface{}) {
	fmt.Println("🛠️ Проверка безопасности инструментов:")

	// Comma-Ok для проверки типа через type assertion
	for name, tool := range tools {
		if sharpness, ok := tool.(int); ok {
			if sharpness > 5 {
				fmt.Printf("⚠️  %s: острота %d/10 (опасно!)\n", name, sharpness)
			} else {
				fmt.Printf("✓ %s: острота %d/10 (безопасно)\n", name, sharpness)
			}
		} else if isDangerous, ok := tool.(bool); ok {
			if isDangerous {
				fmt.Printf("⚠️  %s: опасный инструмент\n", name)
			} else {
				fmt.Printf("✓ %s: безопасный инструмент\n", name)
			}
		} else {
			fmt.Printf("? %s: неизвестный тип инструмента\n", name)
		}
	}
	fmt.Println()
}

// AnimalCare демонстрирует работу с интерфейсами
type Animal interface {
	MakeSound() string
}

type Cow struct{ Name string }
type Chicken struct{ Name string }

func (c Cow) MakeSound() string     { return "Муууу!" }
func (c Chicken) MakeSound() string { return "Куд-куда!" }

func HandleAnimal(animal interface{}) {
	// Comma-Ok для проверки реализации интерфейса
	if cow, ok := animal.(Cow); ok {
		fmt.Printf("🐄 %s говорит: %s\n", cow.Name, cow.MakeSound())
	} else if chicken, ok := animal.(Chicken); ok {
		fmt.Printf("🐔 %s говорит: %s\n", chicken.Name, chicken.MakeSound())
	} else {
		fmt.Printf("❓ Неизвестное животное: %v\n", animal)
	}
}

// ChannelOperations демонстрирует Comma-Ok с каналами
func DemonstrateChannels() {
	fmt.Println("📡 Работа с каналами (сенокосные сигналы):")

	hayReady := make(chan string, 2)
	hayReady <- "Первая партия сена готова!"
	hayReady <- "Вторая партия сена готова!"
	close(hayReady)

	for {
		if message, ok := <-hayReady; ok {
			fmt.Printf("  📢 %s\n", message)
		} else {
			fmt.Println("  ✅ Все сообщения о сене получены!")
			break
		}
	}
	fmt.Println()
}

// AdvancedTypeChecking показывает расширенную проверку типов
func AdvancedTypeChecking() {
	fmt.Println("🔍 Расширенная проверка типов с помощью reflect:")

	items := []interface{}{
		"вилы",
		42,
		3.14,
		Cow{Name: "Бурёнка"},
		true,
		[]string{"сеноворошилка", "грабли"},
	}

	for i, item := range items {
		fmt.Printf("  %d. ", i+1)

		// Используем reflect для определения типа
		switch v := item.(type) {
		case string:
			fmt.Printf("Строка: %s\n", v)
		case int:
			fmt.Printf("Целое число: %d\n", v)
		case float64:
			fmt.Printf("Дробное число: %.2f\n", v)
		case Cow:
			fmt.Printf("Животное: %s\n", v.Name)
		case bool:
			fmt.Printf("Булево значение: %t\n", v)
		default:
			fmt.Printf("Неизвестный тип: %v\n", reflect.TypeOf(v))
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("🚜 Day 73: Comma-Ok Idiom в фермерских условиях!")
	fmt.Println("================================================")

	// Создаем стог сена
	haystack := NewHaystack()
	haystack.AddItem("needle", "золотая иголка")
	haystack.AddItem("hay", "свежее сено")
	haystack.AddItem("tools", map[string]int{"вилы": 3, "грабли": 2})

	// Демонстрация 1: Поиск в map
	fmt.Println("1. 🔎 Ищем иголку в стоге сена:")
	if needle, found := haystack.FindNeedle(); found {
		fmt.Printf("   🎉 Найдена: %s\n\n", needle)
	} else {
		fmt.Println("   😞 Иголка не найдена")
	}

	// Демонстрация 2: Анализ типов сена
	haystack.CountHayTypes()

	// Демонстрация 3: Проверка инструментов
	tools := map[string]interface{}{
		"вилы":       7,
		"грабли":     3,
		"трактор":    true,
		"сенокосилка": 9,
	}
	CheckToolSafety(tools)

	// Демонстрация 4: Работа с интерфейсами
	fmt.Println("4. 🐮 Работа с животными через интерфейсы:")
	animals := []interface{}{
		Cow{Name: "Зорька"},
		Chicken{Name: "Ряба"},
		"не животное",
		42,
	}

	for _, animal := range animals {
		HandleAnimal(animal)
	}
	fmt.Println()

	// Демонстрация 5: Каналы
	DemonstrateChannels()

	// Демонстрация 6: Расширенная проверка типов
	AdvancedTypeChecking()

	// Практический пример: Симуляция фермерской работы
	fmt.Println("🏁 Практический пример: Симуляция рабочего дня")
	SimulateWorkDay()
}

// SimulateWorkDay показывает комбинированное использование Comma-Ok
func SimulateWorkDay() {
	workLog := map[string]interface{}{
		"hours_worked": 8,
		"hay_bales":    150,
		"weather":      "sunny",
		"equipment_used": []string{"трактор", "сеноворошилка", "пресс-подборщик"},
		"animals_fed":  true,
	}

	fmt.Println("📝 Отчет о рабочем дне:")

	// Проверяем различные типы данных в логе
	if hours, ok := workLog["hours_worked"].(int); ok {
		fmt.Printf("   ⏰ Отработано часов: %d\n", hours)
	}

	if bales, ok := workLog["hay_bales"].(int); ok {
		fmt.Printf("   🌾 Собрано тюков сена: %d\n", bales)
	}

	if weather, ok := workLog["weather"].(string); ok {
		fmt.Printf("   🌤️  Погода: %s\n", weather)
	}

	if equipment, ok := workLog["equipment_used"].([]string); ok {
		fmt.Printf("   🛠️  Использованное оборудование: %v\n", equipment)
	}

	if fed, ok := workLog["animals_fed"].(bool); ok && fed {
		fmt.Println("   🐮 Животные накормлены: да")
	}

	fmt.Println("\n🎉 Отлично! Comma-Ok Idiom освоен в полевых условиях!")
}
