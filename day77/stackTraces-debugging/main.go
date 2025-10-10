package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

// CoffeePotatoHybrid представляет наш кофе-картофельный эксперимент
type CoffeePotatoHybrid struct {
	Name        string
	GrowthStage int
}

// NewCoffeePotato создает новый гибрид кофе с картофелем - Кофертофель
func NewCoffeePotato(name string) *CoffeePotatoHybrid {
	return &CoffeePotatoHybrid{
		Name:        name,
		GrowthStage: 0,
	}
}

// Grow пытается вырастить гибрид - может вызвать проблемы!
func (cp *CoffeePotatoHybrid) Grow() {
	cp.GrowthStage++
	fmt.Printf("🌱 %s растет... этап %d\n", cp.Name, cp.GrowthStage)

	switch cp.GrowthStage {
	case 1:
		cp.simulateRootProblem()
	case 2:
		cp.simulateNutritionProblem()
	case 3:
		cp.simulateGeneticProblem()
	default:
		fmt.Printf("✅ %s растет нормально\n", cp.Name)
	}
}

func (cp *CoffeePotatoHybrid) simulateRootProblem() {
	fmt.Println("   🚨 Проблема с корневой системой!")
	fmt.Printf("   📍 Stack Trace:\n")
	debug.PrintStack()
}

func (cp *CoffeePotatoHybrid) simulateNutritionProblem() {
	fmt.Println("   🚨 Дисбаланс питательных веществ!")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("   💥 Паника перехвачена: %v\n", r)
			fmt.Printf("   📍 Stack Trace:\n%s\n", debug.Stack())
		}
	}()

	panic("СЛИШКОМ МНОГО КОФЕ! Гибрид перевозбужден!")
}

func (cp *CoffeePotatoHybrid) simulateGeneticProblem() {
	fmt.Println("   🚨 Генетическая несовместимость!")

	buf := make([]byte, 1024)
	n := runtime.Stack(buf, true)
	fmt.Printf("   📍 Подробный Stack Trace:\n%s\n", string(buf[:n]))
}

// BrazilianResearchTeam представляет бразильских исследователей
type BrazilianResearchTeam struct {
	Members []string
}

func (brt *BrazilianResearchTeam) AnalyzeHybrid(hybrid *CoffeePotatoHybrid) {
	fmt.Printf("\n🇧🇷 Бразильская команда анализирует %s...\n", hybrid.Name)
	brt.analyzeMemory()
	brt.analyzeGoroutines()
}

func (brt *BrazilianResearchTeam) analyzeMemory() {
	fmt.Println("   🧠 Анализ памяти:")

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("   • Память: %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("   • Горутины: %d\n", runtime.NumGoroutine())
}

func (brt *BrazilianResearchTeam) analyzeGoroutines() {
	fmt.Println("   🧵 Анализ горутин:")

	buf := make([]byte, 1024)
	n := runtime.Stack(buf, true)
	fmt.Printf("   • Stack всех горутин:\n%s\n", string(buf[:n]))
}

// AdvancedDebugging демонстрирует продвинутые техники
func AdvancedDebugging() {
	fmt.Println("\n🔧 ПРОДВИНУТЫЕ ТЕХНИКИ DEBUGGING:")

	// 1. Чтение stack trace из panic
	fmt.Println("1. Анализ stack trace из panic:")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("   💥 Перехваченная паника: %v\n", r)
				fmt.Printf("   📍 Stack trace:\n%s\n", debug.Stack())
			}
		}()

		panic("ЭКСПЕРИМЕНТАЛЬНАЯ ПАНИКА ДЛЯ АНАЛИЗА")
	}()
}

// DebuggingTools демонстрирует встроенные инструменты
func DebuggingTools() {
	fmt.Println("\n🛠️ ВСТРОЕННЫЕ ИНСТРУМЕНТЫ DEBUGGING:")

	fmt.Println("1. runtime.Caller() - информация о вызове:")
	for i := 0; i < 2; i++ {
		if pc, file, line, ok := runtime.Caller(i); ok {
			fn := runtime.FuncForPC(pc)
			fmt.Printf("   • Уровень %d: %s (%s:%d)\n", i, fn.Name(), file, line)
		}
	}

	fmt.Printf("\n2. Версия Go: %s\n", runtime.Version())
	fmt.Printf("3. GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}

func main() {
	fmt.Println("🌱☕ Day 77: Stack Traces & Debugging - Кофе-картофельные эксперименты!")
	fmt.Println("======================================================================")

	// Создаем наш экспериментальный гибрид
	hybrid := NewCoffeePotato("Кофертофель-2030")

	fmt.Println("🧪 НАЧИНАЕМ ЭКСПЕРИМЕНТ:")

	// Выращиваем гибрид через несколько этапов
	for i := 0; i < 4; i++ {
		hybrid.Grow()
		time.Sleep(500 * time.Millisecond)
	}

	// Бразильская команда анализирует результаты
	brazilianTeam := &BrazilianResearchTeam{
		Members: []string{"Карлос", "Мария", "Фернандо"},
	}
	brazilianTeam.AnalyzeHybrid(hybrid)

	// Демонстрация продвинутых техник
	AdvancedDebugging()

	// Показываем инструменты debugging
	DebuggingTools()

	// Практические советы
	fmt.Println("\n💡 ПРАКТИЧЕСКИЕ СОВЕТЫ:")
	fmt.Println("   • debug.PrintStack() - быстрая отладка")
	fmt.Println("   • runtime.Stack() - больше контроля")
	fmt.Println("   • debug.Stack() - возвращает как []byte")
	fmt.Println("   • runtime.Caller() - полезно для логирования")

	fmt.Println("\n🎯 ВЫВОДЫ:")
	fmt.Println("   • Stack traces ведут к корню проблемы")
	fmt.Println("   • Debugging требует терпения и инструментов")
	fmt.Println("   • Бразильский подход: работа в команде!")

	fmt.Println("\n💪 Отлично! Теперь ты умеешь читать следы ошибок!")
}
