package main

import (
	"fmt"
	"strings"
)

// DopamineLevel - уровень дофамина от программирования
type DopamineLevel int

const (
	NoDopamine DopamineLevel = iota
	LowDopamine
	MediumDopamine
	HighDopamine
	DopamineFlood
)

// TrollMessage - тип для сообщений троллей с уровнем троттинга
type TrollMessage struct {
	Text     string
	Level    int
	IsFriend bool
}

// GenericResponse - обобщенный ответ на разные типы раздражителей
type Response[T any] struct {
	Input     T
	Processed bool
	Dopamine  DopamineLevel
	Action    string
}

// GenericProcessor - дженерик-обработчик для любого типа ввода
type GenericProcessor[T any] interface {
	Process(input T) Response[T]
	GetDopamineReward() DopamineLevel
}

// WeatherProcessor - обработчик погодных условий
type WeatherProcessor struct{}

func (wp WeatherProcessor) Process(weather string) Response[string] {
	var response Response[string]
	response.Input = weather
	
	if strings.Contains(strings.ToLower(weather), "метель") || 
	   strings.Contains(strings.ToLower(weather), "снег") ||
	   strings.Contains(strings.ToLower(weather), "сугроб") {
		response.Processed = true
		response.Dopamine = HighDopamine
		response.Action = "Сижу дома, изучаю Go!"
	} else {
		response.Dopamine = LowDopamine
		response.Action = "Можно выйти, но лучше почитать про type inference"
	}
	
	return response
}

func (wp WeatherProcessor) GetDopamineReward() DopamineLevel {
	return MediumDopamine
}

// TrollProcessor - обработчик троллинга с использованием type inference
type TrollProcessor struct {
	IgnoredCount int
}

func (tp *TrollProcessor) Process(message TrollMessage) Response[TrollMessage] {
	var response Response[TrollMessage]
	response.Input = message
	
	// Автоматический вывод типа и обработка
	if message.Level >= 80 && message.IsFriend {
		response.Processed = true
		response.Dopamine = DopamineFlood
		response.Action = "Тролль уровня 80! Удалить сообщение и писать код"
		tp.IgnoredCount++
	} else if message.Level > 50 {
		response.Dopamine = HighDopamine
		response.Action = "Игнорировать тролля, добавить +10 строк кода"
		tp.IgnoredCount++
	} else {
		response.Dopamine = MediumDopamine
		response.Action = "Прочитать и забыть, продолжить изучение дженериков"
	}
	
	return response
}

func (tp *TrollProcessor) GetDopamineReward() DopamineLevel {
	return HighDopamine
}

// FoodProcessor - обработчик продуктов питания
type FoodProcessor struct{}

func (fp FoodProcessor) Process(food string) Response[string] {
	var response Response[string]
	response.Input = food
	
	favoriteFoods := []string{"беляш", "холодец", "холопеньё", "чай"}
	for _, fav := range favoriteFoods {
		if strings.Contains(strings.ToLower(food), fav) {
			response.Processed = true
			response.Dopamine = MediumDopamine
			response.Action = "Подкрепился, теперь можно учить type inference!"
			return response
		}
	}
	
	response.Dopamine = LowDopamine
	response.Action = "Похоже, пора в магазин... но сначала коммит!"
	return response
}

func (fp FoodProcessor) GetDopamineReward() DopamineLevel {
	return LowDopamine
}

// Generic function с type inference
func ProcessInput[T any, P GenericProcessor[T]](input T, processor P) Response[T] {
	fmt.Printf("🎯 Обработка ввода типа: %T\n", input)
	response := processor.Process(input)
	
	// Геймификация: дофаминовая награда
	dopamineReward := processor.GetDopamineReward()
	fmt.Printf("💊 Дофамин за обработку: %v\n", dopamineReward)
	
	// Автоматический вывод типа в действии
	fmt.Printf("🔧 Type inference определил: %T -> %T\n", input, response)
	
	return response
}

func main() {
	fmt.Println("🚀 DAY 67: Generics in Go - Type Inference")
	fmt.Println(strings.Repeat("=", 50))
	
	// 1. Обработка погоды
	fmt.Println("\n🌨️  СИТУАЦИЯ: Проснулся, смотрю в окно")
	weather := "На улице метель, сугробы по колено, соседи откапывают машины"
	weatherProc := WeatherProcessor{}
	weatherResponse := ProcessInput(weather, weatherProc)
	fmt.Printf("   Действие: %s\n", weatherResponse.Action)
	fmt.Printf("   Дофамин: %v\n\n", weatherResponse.Dopamine)
	
	// 2. Обработка тролля
	fmt.Println("👹 СИТУАЦИЯ: Тролль Рокки в личке")
	trollMsg := TrollMessage{
		Text:     "Гошик, почему ты 'друга' называешь Троллем?",
		Level:    80,
		IsFriend: true,
	}
	trollProc := &TrollProcessor{}
	trollResponse := ProcessInput(trollMsg, trollProc)
	fmt.Printf("   Действие: %s\n", trollResponse.Action)
	fmt.Printf("   Дофамин: %v\n", trollResponse.Dopamine)
	fmt.Printf("   Проигнорировано троллей сегодня: %d\n\n", trollProc.IgnoredCount)
	
	// 3. Обработка еды
	fmt.Println("🍲 СИТУАЦИЯ: Завтрак в метель")
	foods := []string{"последний беляш", "холодец с холопеньём", "горячий чай"}
	foodProc := FoodProcessor{}
	
	totalFoodDopamine := NoDopamine
	for _, food := range foods {
		foodResponse := ProcessInput(food, foodProc)
		fmt.Printf("   %s → %s\n", food, foodResponse.Action)
		if foodResponse.Dopamine > totalFoodDopamine {
			totalFoodDopamine = foodResponse.Dopamine
		}
	}
	
	// 4. Генерация дофамина от программирования
	fmt.Println("\n💻 СИТУАЦИЯ: Самое время для Go!")
	fmt.Println("   Автоматический вывод типов (type inference) в действии:")
	
	// Демонстрация type inference с дженериками
	processors := []interface{}{weatherProc, trollProc, foodProc}
	fmt.Printf("   Количество обработчиков: %d\n", len(processors))
	
	// Симуляция дофаминового цикла
	fmt.Println("\n🧠 ДОФАМИНОВЫЙ ЦИКЛ ПРОГРАММИСТА:")
	steps := []struct{
		action string
		dopamine DopamineLevel
	}{
		{"Проснулся в 7 утра", LowDopamine},
		{"Увидел метель за окном", MediumDopamine},
		{"Проигнорировал тролля уровня 80", HighDopamine},
		{"Изучил type inference в дженериках", DopamineFlood},
		{"Сделал коммит с обновлениями", HighDopamine},
		{"Предвкушение завтрашнего кода", MediumDopamine},
	}
	
	totalDopamine := NoDopamine
	for _, step := range steps {
		fmt.Printf("   ✅ %-40s → %v\n", step.action, step.dopamine)
		totalDopamine += step.dopamine
	}
	
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("📊 ИТОГ ДНЯ 67:\n")
	fmt.Printf("   Уровень дофамина: %v\n", totalDopamine)
	fmt.Printf("   Игнорировано троллей: %d\n", trollProc.IgnoredCount)
	fmt.Printf("   Вывод: Type inference определен успешно!\n")
	fmt.Println("\n🔥 ВЫВОД: Не трать время на оправдания троллям.")
	fmt.Println("   Потрать его на изучение дженериков в Go!")
	fmt.Println("   Каждая строка кода = +1 к дофамину, каждый тролль = +10 к фокусу")
}
