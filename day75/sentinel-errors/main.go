package main

import (
	"errors"
	"fmt"
	"time"
)

// Sentinel Errors - наши "сторожевые" ошибки для тракторного сервиса
var (
	ErrNoOil          = errors.New("масло в двигателе закончилось")
	ErrLowPressure    = errors.New("давление в шинах слишком низкое")
	ErrBatteryDead    = errors.New("аккумулятор разряжен")
	ErrFuelEmpty      = errors.New("топливный бак пуст")
	ErrTransmission   = errors.New("проблемы с коробкой передач")
	ErrStarterFailure = errors.New("неисправность стартера")
)

// Tractor представляет наш трактор
type Tractor struct {
	Model       string
	OilLevel    float64 // уровень масла (0.0 - 1.0)
	TirePressure float64 // давление в шинах (бар)
	Battery     float64 // заряд аккумулятора (0.0 - 1.0)
	Fuel        float64 // уровень топлива (0.0 - 1.0)
}

// NewTractor создает новый трактор
func NewTractor(model string) *Tractor {
	return &Tractor{
		Model:       model,
		OilLevel:    0.8,
		TirePressure: 2.5,
		Battery:     0.9,
		Fuel:        0.7,
	}
}

// CheckEngine проверяет двигатель трактора
func (t *Tractor) CheckEngine() error {
	fmt.Printf("🔧 Проверяем двигатель %s...\n", t.Model)
	time.Sleep(1 * time.Second)

	if t.OilLevel < 0.1 {
		return ErrNoOil
	}

	if t.OilLevel < 0.3 {
		return fmt.Errorf("уровень масла низкий: %.1f%%, нужно долить", t.OilLevel*100)
	}

	fmt.Println("✅ Двигатель в порядке!")
	return nil
}

// CheckTires проверяет шины
func (t *Tractor) CheckTires() error {
	fmt.Printf("🎯 Проверяем шины...\n")
	time.Sleep(1 * time.Second)

	if t.TirePressure < 1.5 {
		return ErrLowPressure
	}

	if t.TirePressure < 2.0 {
		return fmt.Errorf("давление в шинах низкое: %.1f бар", t.TirePressure)
	}

	fmt.Println("✅ Шины в норме!")
	return nil
}

// CheckElectrical проверяет электрическую систему
func (t *Tractor) CheckElectrical() error {
	fmt.Printf("⚡ Проверяем электрику...\n")
	time.Sleep(1 * time.Second)

	if t.Battery < 0.1 {
		return ErrBatteryDead
	}

	if t.Battery < 0.5 {
		return fmt.Errorf("заряд аккумулятора низкий: %.1f%%", t.Battery*100)
	}

	fmt.Println("✅ Электрика в порядке!")
	return nil
}

// CheckFuel проверяет топливную систему
func (t *Tractor) CheckFuel() error {
	fmt.Printf("⛽ Проверяем топливную систему...\n")
	time.Sleep(1 * time.Second)

	if t.Fuel < 0.05 {
		return ErrFuelEmpty
	}

	if t.Fuel < 0.2 {
		return fmt.Errorf("топлива мало: %.1f%%", t.Fuel*100)
	}

	fmt.Println("✅ Топливо в норме!")
	return nil
}

// Start пытается завести трактор
func (t *Tractor) Start() error {
	fmt.Printf("\n🚜 Пытаемся завести трактор %s...\n", t.Model)
	time.Sleep(2 * time.Second)

	// Проверяем все системы перед запуском
	if err := t.CheckFuel(); err != nil {
		return fmt.Errorf("не удалось завести: %w", err)
	}

	if err := t.CheckElectrical(); err != nil {
		return fmt.Errorf("не удалось завести: %w", err)
	}

	if err := t.CheckEngine(); err != nil {
		return fmt.Errorf("не удалось завести: %w", err)
	}

	fmt.Println("✅✅✅ ТРАКТОР ЗАВЕЛСЯ! Можно ехать в поле!")
	return nil
}

// Diagnose выполняет полную диагностику
func (t *Tractor) Diagnose() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🏥 ПОЛНАЯ ДИАГНОСТИКА ТРАКТОРА")
	fmt.Println(strings.Repeat("=", 50))

	checks := []func() error{
		t.CheckFuel,
		t.CheckElectrical,
		t.CheckEngine,
		t.CheckTires,
	}

	for _, check := range checks {
		if err := check(); err != nil {
			fmt.Printf("❌ Обнаружена проблема: %v\n", err)
		}
	}
}

// RepairService сервис по ремонту тракторов
type RepairService struct {
	Name string
}

// NewRepairService создает новый сервис
func NewRepairService(name string) *RepairService {
	return &RepairService{Name: name}
}

// HandleError обрабатывает ошибки трактора
func (rs *RepairService) HandleError(err error) {
	fmt.Printf("\n🔧 Сервис '%s' начинает работу...\n", rs.Name)

	// Используем errors.Is для проверки sentinel errors
	switch {
	case errors.Is(err, ErrNoOil):
		fmt.Println("💧 Проблема: Нет масла. Решение: Заливаем новое масло!")

	case errors.Is(err, ErrLowPressure):
		fmt.Println("🎯 Проблема: Низкое давление. Решение: Накачиваем шины!")

	case errors.Is(err, ErrBatteryDead):
		fmt.Println("⚡ Проблема: Разряжен аккумулятор. Решение: Заряжаем или меняем!")

	case errors.Is(err, ErrFuelEmpty):
		fmt.Println("⛽ Проблема: Нет топлива. Решение: Заправляем трактор!")

	case errors.Is(err, ErrTransmission):
		fmt.Println("🔩 Проблема: Коробка передач. Решение: Вызываем Пахомыча!")

	case errors.Is(err, ErrStarterFailure):
		fmt.Println("🔄 Проблема: Стартер. Решение: Ремонтируем систему запуска!")

	default:
		fmt.Printf("❓ Неизвестная проблема: %v\n", err)
		fmt.Println("🤔 Нужна дополнительная диагностика...")
	}
}

// DemonstrateSentinelErrors показывает различные сценарии
func DemonstrateSentinelErrors() {
	fmt.Println("🎬 ДЕМОНСТРАЦИЯ SENTINEL ERRORS")
	fmt.Println(strings.Repeat("=", 40))

	// Сценарий 1: Трактор без масла
	fmt.Println("\n1. Трактор 'Ударник' - проблемы с маслом:")
	tractor1 := NewTractor("Ударник")
	tractor1.OilLevel = 0.05 // Очень мало масла

	if err := tractor1.CheckEngine(); err != nil {
		fmt.Printf("   Обнаружена: %v\n", err)
		if errors.Is(err, ErrNoOil) {
			fmt.Println("   ✅ Это именно та ошибка, которую мы ожидали!")
		}
	}

	// Сценарий 2: Трактор с разряженным аккумулятором
	fmt.Println("\n2. Трактор 'Быстрый' - проблемы с аккумулятором:")
	tractor2 := NewTractor("Быстрый")
	tractor2.Battery = 0.05 // Почти разряжен

	if err := tractor2.CheckElectrical(); err != nil {
		fmt.Printf("   Обнаружена: %v\n", err)
		if errors.Is(err, ErrBatteryDead) {
			fmt.Println("   ✅ Аккумулятор требует замены!")
		}
	}

	// Сценарий 3: Цепочка ошибок
	fmt.Println("\n3. Трактор 'Старый' - множественные проблемы:")
	tractor3 := NewTractor("Старый")
	tractor3.Fuel = 0.0
	tractor3.OilLevel = 0.0

	if err := tractor3.Start(); err != nil {
		fmt.Printf("   Ошибка запуска: %v\n", err)

		// Проверяем корневые причины
		if errors.Is(err, ErrFuelEmpty) {
			fmt.Println("   💡 Основная причина: нет топлива!")
		}
		if errors.Is(err, ErrNoOil) {
			fmt.Println("   💡 Основная причина: нет масла!")
		}
	}
}

func main() {
	fmt.Println("🚜 Day 75: Sentinel Errors - Ремонт трактора с Пахомычем!")
	fmt.Println(strings.Repeat("=", 60))

	// Демонстрация возможностей sentinel errors
	DemonstrateSentinelErrors()

	// Создаем сервис ремонта
	service := NewRepairService("Пахомыч и Компания")

	// Тестируем различные случаи
	fmt.Println("\n" + strings.Repeat("💥", 25))
	fmt.Println("РЕАЛЬНЫЕ СЛУЧАИ ИЗ ЖИЗНИ:")
	fmt.Println(strings.Repeat("💥", 25))

	testCases := []struct {
		name  string
		tractor *Tractor
	}{
		{"Трактор без топлива", &Tractor{Model: "Голодарь", Fuel: 0.0}},
		{"Трактор с севшим аккумулятором", &Tractor{Model: "Тихоня", Battery: 0.0}},
		{"Трактор без масла", &Tractor{Model: "Сухарь", OilLevel: 0.0}},
		{"Трактор со спущенными шинами", &Tractor{Model: "Призрак", TirePressure: 1.0}},
	}

	for _, tc := range testCases {
		fmt.Printf("\n📋 Случай: %s\n", tc.name)
		if err := tc.tractor.Start(); err != nil {
			service.HandleError(err)
		}
		time.Sleep(1 * time.Second)
	}

	// Заключительная диагностика
	fmt.Println("\n" + strings.Repeat("📊", 20))
	fmt.Println("ИТОГИ ОБУЧЕНИЯ:")
	fmt.Println(strings.Repeat("📊", 20))

	workingTractor := NewTractor("Идеальный")
	workingTractor.Diagnose()

	fmt.Println("\n🎯 Преимущества Sentinel Errors:")
	fmt.Println("   ✅ Ясность - сразу понятен тип ошибки")
	fmt.Println("   ✅ Сравнение - используем errors.Is()")
	fmt.Println("   ✅ Документация - ошибки как контракты")
	fmt.Println("   ✅ Тестирование - легко тестировать конкретные случаи")

	fmt.Println("\n💪 Отлично! Теперь ты знаешь Sentinel Errors!")
	fmt.Println("   Помни: каждая ошибка - это возможность научиться чему-то новому!")
}

// Вспомогательная функция для strings.Repeat
var strings = struct {
	Repeat func(string, int) string
}{
	Repeat: func(s string, count int) string {
		result := ""
		for i := 0; i < count; i++ {
			result += s
		}
		return result
	},
}

