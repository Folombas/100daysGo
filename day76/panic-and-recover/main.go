package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Cow представляет умную корову с AI-датчиками
type Cow struct {
	Name     string
	Mood     string
	Milk     float64
	AIStatus string
}

// NewCow создает новую AI-корову
func NewCow(name string) *Cow {
	return &Cow{
		Name:     name,
		Mood:     "спокойная",
		Milk:     5.0,
		AIStatus: "активен",
	}
}

// MilkCow пытается подоить корову - может вызвать панику!
func MilkCow(cow *Cow) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("🚨 CRITICAL: AI-доярка в панике! Причина: %v\n", r)
			fmt.Println("🔄 Активируем протокол восстановления...")
			cow.AIStatus = "восстановление"
			time.Sleep(2 * time.Second)
			cow.AIStatus = "стабилен"
			fmt.Println("✅ Система восстановлена! Продолжаем работу.")
		}
	}()

	fmt.Printf("🤖 AI-доярка начинает доить %s...\n", cow.Name)
	time.Sleep(1 * time.Second)

	// Симуляция различных сценариев
	switch rand.Intn(5) {
	case 0:
		panic("КОРОВА УДАРИЛА НОГОЙ! AI-сенсоры повреждены!")
	case 1:
		panic("ПЕРЕГРУЗКА AI-СИСТЕМЫ: слишком много данных о настроении коровы!")
	case 2:
		panic("СБОЙ МОЛОКОПРОВОДА: давление превышено!")
	case 3:
		// Успешная дойка
		milk := cow.Milk * (0.8 + rand.Float64()*0.4)
		fmt.Printf("✅ Успешно подоено: %.1f литров молока\n", milk)
		cow.Mood = "довольная"
	default:
		// Еще одна успешная дойка
		milk := cow.Milk * (0.7 + rand.Float64()*0.3)
		fmt.Printf("✅ Успешно подоено: %.1f литров молока\n", milk)
		cow.Mood = "расслабленная"
	}
}

// SmartMilkingSystem представляет умную систему доения
type SmartMilkingSystem struct {
	Cows         []*Cow
	PanicCount   int
	SuccessCount int
}

// NewSmartMilkingSystem создает новую систему
func NewSmartMilkingSystem() *SmartMilkingSystem {
	return &SmartMilkingSystem{
		Cows: []*Cow{
			NewCow("Бурёнка-3000"),
			NewCow("AI-Зорька"),
			NewCow("Кибер-Ромашка"),
			NewCow("Нано-Манька"),
		},
	}
}

// StartMilkingSession запускает сессию доения с защитой от паники
func (sms *SmartMilkingSystem) StartMilkingSession() {
	fmt.Println("\n🐄 ЗАПУСК AI-СИСТЕМЫ ДОЕНИЯ КОРОВ")
	fmt.Println("==========================================")

	for _, cow := range sms.Cows {
		// Защищаем каждую операцию доения recover'ом
		func() {
			defer func() {
				if r := recover(); r != nil {
					sms.PanicCount++
					fmt.Printf("📊 Статистика: паник - %d, успехов - %d\n\n",
						sms.PanicCount, sms.SuccessCount)
				}
			}()

			MilkCow(cow)
			sms.SuccessCount++
		}()
		time.Sleep(1 * time.Second)
	}
}

// TestVariousPanics демонстрирует разные типы паник
func TestVariousPanics() {
	fmt.Println("\n🔬 ДЕМОНСТРАЦИЯ РАЗНЫХ ТИПОВ PANIC:")
	fmt.Println("==================================")

	// 1. Паника с nil pointer
	fmt.Println("1. Nil pointer dereference:")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("   Перехвачено: %v\n", r)
			}
		}()
		var cow *Cow
		fmt.Println(cow.Name) // Здесь будет паника!
	}()

	// 2. Паника с индексом за пределами массива
	fmt.Println("\n2. Выход за границы массива:")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("   Перехвачено: %v\n", r)
			}
		}()
		cows := make([]*Cow, 2)
		fmt.Println(cows[10]) // Здесь будет паника!
	}()

	// 3. Паника с делением на ноль
	fmt.Println("\n3. Деление на ноль:")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("   Перехвачено: %v\n", r)
			}
		}()
		x := 10
		y := 0
		fmt.Println(x / y) // Здесь будет паника!
	}()

	// 4. Явный вызов panic
	fmt.Println("\n4. Явный вызов panic():")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("   Перехвачено: %v\n", r)
			}
		}()
		panic("ЭКСПЕРИМЕНТАЛЬНАЯ ПАНИКА: тестируем систему!")
	}()
}

// AdvancedRecovery демонстрирует продвинутое восстановление
func AdvancedRecovery() {
	fmt.Println("\n🎯 ПРОДВИНУТАЯ СИСТЕМА ВОССТАНОВЛЕНИЯ:")
	fmt.Println("=====================================")

	processCow := func(name string) (result string, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("восстановлено после паники: %v", r)
			}
		}()

		// Имитация работы с возможной паникой
		if rand.Intn(2) == 0 {
			panic("внезапный сбой в AI-анализе поведения коровы")
		}

		return fmt.Sprintf("Корова %s успешно обработана AI-системой", name), nil
	}

	// Обрабатываем несколько коров
	cows := []string{"AI-Белка", "Кибер-Стрелка", "Нано-Пятнашка"}
	for _, cow := range cows {
		result, err := processCow(cow)
		if err != nil {
			fmt.Printf("❌ %s: %v\n", cow, err)
		} else {
			fmt.Printf("✅ %s\n", result)
		}
	}
}

// ChineseStudentAI представляет китайских студентов с их AI
type ChineseStudentAI struct {
	Name string
}

func (cs *ChineseStudentAI) AnalyzeCowBehavior(cow *Cow) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("🇨🇳 %s: 不用担心! 我们修复了问题: %v\n", cs.Name, r)
		}
	}()

	fmt.Printf("\n🇨🇳 %s анализирует поведение %s...\n", cs.Name, cow.Name)

	// Китайские студенты тестируют сложные AI-алгоритмы
	complexCalculations := []func(){
		func() { panic("神经网络过载: 太多的奶牛 эмоций!") },
		func() { panic("数据溢出: 牛奶流量 превысил ожидания!") },
		func() { fmt.Println("分析成功: 奶牛 счастлива!") },
	}

	complexCalculations[rand.Intn(len(complexCalculations))]()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🤖 Day 76: Panic and Recover - AI-доярка и непокорные коровы!")
	fmt.Println("==========================================================")

	// Демонстрация различных типов паник
	TestVariousPanics()

	// Запускаем систему доения
	system := NewSmartMilkingSystem()

	// Проводим несколько сессий доения
	fmt.Println("\n🏁 ЗАПУСК ПРОИЗВОДСТВЕННОЙ СИСТЕМЫ:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("\n--- Сессия доения #%d ---\n", i)
		system.StartMilkingSession()
		time.Sleep(2 * time.Second)
	}

	// Демонстрация продвинутого восстановления
	AdvancedRecovery()

	// Китайские студенты тестируют свои AI-алгоритмы
	fmt.Println("\n👨‍🎓 КИТАЙСКИЕ СТУДЕНТЫ ТЕСТИРУЮТ AI:")
	fmt.Println("=================================")

	students := []*ChineseStudentAI{
		{Name: "张伟"},
		{Name: "李娜"},
		{Name: "王鹏"},
	}

	cow := NewCow("Экспериментальная корова")
	for _, student := range students {
		student.AnalyzeCowBehavior(cow)
		time.Sleep(1 * time.Second)
	}

	// Итоговая статистика
	fmt.Println("\n📊 ИТОГОВАЯ СТАТИСТИКА:")
	fmt.Printf("   Успешных доек: %d\n", system.SuccessCount)
	fmt.Printf("   Паник восстановлено: %d\n", system.PanicCount)
	fmt.Printf("   Общая эффективность: %.1f%%\n",
		float64(system.SuccessCount)/float64(system.SuccessCount+system.PanicCount)*100)

	fmt.Println("\n🎯 ВЫВОДЫ О PANIC/RECOVER:")
	fmt.Println("   • Panic - для действительно критических ситуаций")
	fmt.Println("   • Recover - только в defer-функциях")
	fmt.Println("   • Не злоупотреблять - обычные ошибки лучше через error")
	fmt.Println("   • Идеально для: горутин, middleware, обработчиков запросов")

	fmt.Println("\n💪 Отлично! Теперь твоя AI-доярка неуязвима для коровьих капризов!")
}
