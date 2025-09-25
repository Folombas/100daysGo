package main

import (
	"fmt"
	"unsafe"
)

// 🏗️ Структура для демонстрации указателей на структуры
type Player struct {
	Name     string
	Level    int
	Health   float64
	IsActive bool
}

// 🎯 Функция, принимающая указатель для модификации структуры
func levelUp(player *Player) {
	player.Level++
	player.Health += 10.0
	fmt.Printf("🎮 %s повысил уровень до %d! Здоровье: %.1f\n",
		player.Name, player.Level, player.Health)
}

// 📊 Функция для демонстрации разницы между значением и указателем
func demonstrateValueVsPointer() {
	fmt.Println("\n🔍 ДЕМОНСТРАЦИЯ: ЗНАЧЕНИЕ vs УКАЗАТЕЛЬ")
	fmt.Println("=========================================")

	// Создаем переменную
	score := 100
	fmt.Printf("🎯 Исходное значение score: %d (адрес: %p)\n", score, &score)

	// Передаем по значению (копия)
	modifyValue(score)
	fmt.Printf("📝 После modifyValue(score): %d\n", score)

	// Передаем по указателю (оригинал)
	modifyPointer(&score)
	fmt.Printf("📍 После modifyPointer(&score): %d\n", score)
}

// 🔄 Модификация по значению (работает с копией)
func modifyValue(value int) {
	value = 200
	fmt.Printf("   🔄 modifyValue: изменили копию на %d\n", value)
}

// 🔄 Модификация по указателю (работает с оригиналом)
func modifyPointer(pointer *int) {
	*pointer = 300
	fmt.Printf("   🔄 modifyPointer: изменили оригинал на %d\n", *pointer)
}

// 🧠 Демонстрация указателей на различные типы данных
func demonstrateDifferentTypes() {
	fmt.Println("\n🎪 УКАЗАТЕЛИ НА РАЗЛИЧНЫЕ ТИПЫ ДАННЫХ")
	fmt.Println("=====================================")

	// Указатель на string
	name := "Гофер"
	namePtr := &name
	fmt.Printf("📝 String: значение=%s, указатель=%p, через указатель=%s\n",
		name, namePtr, *namePtr)

	// Указатель на int
	age := 25
	agePtr := &age
	fmt.Printf("🔢 Int: значение=%d, указатель=%p, через указатель=%d\n",
		age, agePtr, *agePtr)

	// Указатель на float
	pi := 3.14159
	piPtr := &pi
	fmt.Printf("📐 Float64: значение=%.5f, указатель=%p, через указатель=%.5f\n",
		pi, piPtr, *piPtr)

	// Указатель на bool
	isReady := true
	isReadyPtr := &isReady
	fmt.Printf("✅ Bool: значение=%t, указатель=%p, через указатель=%t\n",
		isReady, isReadyPtr, *isReadyPtr)
}

// 💾 Демонстрация работы с указателями на структуры
func demonstrateStructPointers() {
	fmt.Println("\n🏗️ УКАЗАТЕЛИ НА СТРУКТУРЫ")
	fmt.Println("=========================")

	// Создаем структуру
	player := Player{
		Name:     "Алексей",
		Level:    1,
		Health:   100.0,
		IsActive: true,
	}

	fmt.Printf("🎮 Исходный игрок: %+v\n", player)
	fmt.Printf("📍 Адрес структуры: %p\n", &player)

	// Передаем указатель на структуру
	levelUp(&player)
	fmt.Printf("🎮 Игрок после levelUp: %+v\n", player)
}

// 🎭 Демонстрация nil указателей и безопасной работы с ними
func demonstrateNilPointers() {
	fmt.Println("\n⚠️ РАБОТА С NIL УКАЗАТЕЛЯМИ")
	fmt.Println("=========================")

	var nilPtr *int
	fmt.Printf("📝 Nil указатель: %p\n", nilPtr)

	// Безопасная проверка перед разыменованием
	if nilPtr != nil {
		fmt.Printf("📍 Значение: %d\n", *nilPtr)
	} else {
		fmt.Println("❌ Ошибка: попытка разыменовать nil указатель!")
	}

	// Создаем валидный указатель
	value := 42
	validPtr := &value

	if validPtr != nil {
		fmt.Printf("✅ Валидный указатель: %p, значение: %d\n", validPtr, *validPtr)
	}
}

// 🔬 Продвинутая демонстрация: арифметика указателей (через unsafe)
func demonstratePointerArithmetic() {
	fmt.Println("\n🔬 ПРОДВИНУТАЯ ДЕМОНСТРАЦИЯ: МАССИВЫ И УКАЗАТЕЛИ")
	fmt.Println("===============================================")

	numbers := [3]int{10, 20, 30}
	fmt.Printf("📦 Массив: %v\n", numbers)

	// Указатель на первый элемент
	firstPtr := &numbers[0]
	fmt.Printf("📍 Указатель на numbers[0]: %p, значение: %d\n", firstPtr, *firstPtr)

	// Используем unsafe для арифметики указателей (осторожно!)
	secondPtr := unsafe.Pointer(uintptr(unsafe.Pointer(firstPtr)) + unsafe.Sizeof(numbers[0]))
	fmt.Printf("📍 Указатель на numbers[1] (через арифметику): %p\n", secondPtr)

	// Безопасный способ получить указатель на второй элемент
	secondPtrSafe := &numbers[1]
	fmt.Printf("📍 Указатель на numbers[1] (безопасно): %p, значение: %d\n",
		secondPtrSafe, *secondPtrSafe)
}

// 🎯 Демонстрация возврата указателей из функций
func createPlayer(name string, level int) *Player {
	fmt.Printf("🏗️ Создаем нового игрока: %s\n", name)
	return &Player{
		Name:     name,
		Level:    level,
		Health:   100.0,
		IsActive: true,
	}
}

func demonstrateFunctionPointers() {
	fmt.Println("\n🔄 УКАЗАТЕЛИ И ФУНКЦИИ")
	fmt.Println("=====================")

	// Функция возвращает указатель
	playerPtr := createPlayer("Мария", 5)
	fmt.Printf("🎮 Игрок через указатель: %+v\n", *playerPtr)
	fmt.Printf("📍 Адрес в памяти: %p\n", playerPtr)

	// Модифицируем через указатель
	playerPtr.Level = 10
	fmt.Printf("🎮 После модификации: %+v\n", *playerPtr)
}

func main() {
	fmt.Println("🎉 ДЕМО УКАЗАТЕЛЕЙ В GO!")
	fmt.Println("========================")
	fmt.Println("💡 Указатели — это переменные, хранящие адреса памяти других переменных")
	fmt.Println("💡 Символ & — получить адрес переменной")
	fmt.Println("💡 Символ * — получить значение по адресу (разыменование)")

	// 🎯 Основная демонстрация
	demonstrateValueVsPointer()
	demonstrateDifferentTypes()
	demonstrateStructPointers()
	demonstrateNilPointers()
	demonstratePointerArithmetic()
	demonstrateFunctionPointers()

	// 🏆 Итоги
	fmt.Println("\n🎯 КЛЮЧЕВЫЕ ВЫВОДЫ:")
	fmt.Println("✅ Указатели позволяют работать напрямую с памятью")
	fmt.Println("✅ &variable — получить адрес переменной")
	fmt.Println("✅ *pointer — получить значение по адресу")
	fmt.Println("✅ Указатели экономят память при работе с большими структурами")
	fmt.Println("✅ Всегда проверяйте указатели на nil перед разыменованием")
	fmt.Println("✅ Используйте указатели для модификации оригинальных значений")
	fmt.Println("🚫 В Go нет арифметики указателей (кроме пакета unsafe)")

	fmt.Println("\n🎊 Поздравляю с освоением указателей в Go! 🎊")
}
