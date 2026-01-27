package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

// ========== DISCLAIMER: ХУДОЖЕСТВЕННЫЙ ВЫМЫСЕЛ ==========
// SnowTrace Debugger — метафора пути в IT. Все персонажи вымышлены.
// Совпадения случайны. Это программная поэзия о преодолении.
// ======================================================

// Ошибки системы
var (
	// SnowDepthError возникает при невозможности пройти через сугроб
	SnowDepthError = errors.New("снежный сугроб глубиной 1+ метр")

	// ClientNotFoundErr возникает когда клиент недоступен
	ClientNotFoundErr = errors.New("клиент не доступен")

	// DeliveryFailedErr общая ошибка доставки
	DeliveryFailedErr = errors.New("доставка провалена")
)

// CustomError — ошибка с контекстом и стектрейсом
type CustomError struct {
	Msg, Ctx, Trace string
	Time            time.Time
}

func (e CustomError) Error() string {
	return fmt.Sprintf("❌ [%s] %s\n   Контекст: %s\n   Трассировка:\n%s",
		e.Time.Format("15:04:05"), e.Msg, e.Ctx, e.Trace)
}

// getStackTrace возвращает стек вызовов
func getStackTrace() string {
	buf := make([]byte, 1024)
	return string(buf[:runtime.Stack(buf, false)])
}

// Courier — курьер Гоша
type Courier struct {
	Name, Role               string
	Focus, Stamina           int // 0-100
	Knowledge, Motivation    int
	DaysAsCourier            int
	IsDebugMode              bool
	ErrorsHandled, ErrorsIgnored int
}

// NewCourier создаёт курьера в состоянии "утро понедельника"
func NewCourier(name string) *Courier {
	return &Courier{
		Name:       name,
		Role:       "courier",
		Focus:      30,
		Stamina:    85,
		Knowledge:  42,
		Motivation: 25,
		IsDebugMode: false,
	}
}

// Deliver — попытка доставки с обработкой ошибок
func (c *Courier) Deliver(address string, hasSnowdrift bool) error {
	if address == "" {
		// Явно используем ClientNotFoundErr
		return fmt.Errorf("%w: %v (адрес пустой)", ClientNotFoundErr, address)
	}

	if !c.IsDebugMode {
		fmt.Println("⚠️  Режим отладки выключен")
	}

	if hasSnowdrift {
		// Явно используем SnowDepthError
		err := CustomError{
			Msg:   SnowDepthError.Error(),
			Ctx:   fmt.Sprintf("адрес: %s, курьер: %s", address, c.Name),
			Trace: getStackTrace(),
			Time:  time.Now(),
		}

		if c.IsDebugMode {
			log.Printf("🚨 ДЕТАЛЬНАЯ ОТЛАДКА:\n%s\n", err)
			c.Knowledge += 15
			fmt.Printf("📈 Понимание +15: %d/100\n", c.Knowledge)
			c.ErrorsHandled++
		} else {
			fmt.Printf("❌ Ошибка доставки: %s\n", address)
			c.Motivation -= 10
			c.ErrorsIgnored++
		}

		// Явно используем DeliveryFailedErr и SnowDepthError
		return fmt.Errorf("%w: %v | %s", DeliveryFailedErr, SnowDepthError, address)
	}

	fmt.Printf("✅ Доставлено: %s\n", address)
	c.Motivation += 5
	return nil
}

// EnableDebug включает режим отладки
func (c *Courier) EnableDebug() {
	c.IsDebugMode = true
	fmt.Println("\n🔍 ВКЛЮЧЁН РЕЖИМ ОТЛАДКИ")
	fmt.Println("   Теперь все ошибки показывают стектрейс")
	c.Knowledge += 20
}

// HandlePanic обрабатывает панику с восстановлением
func (c *Courier) HandlePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\n🚑 ВОССТАНОВЛЕНИЕ ПАНИКИ: %v\n", r)
			fmt.Println("Стектрейс:", getStackTrace())
			c.Focus = max(c.Focus-20, 0)
		}
	}()

	if c.Motivation < 20 {
		panic("критический уровень мотивации: депрессия")
	}
}

// Status показывает статус курьера
func (c *Courier) Status() {
	fmt.Println("\n" + strings.Repeat("▬", 50))
	fmt.Println("📊 СТАТУС КУРЬЕРА:")
	fmt.Printf("   Фокус: %d/100 | Выносливость: %d/100\n", c.Focus, c.Stamina)
	fmt.Printf("   Знания: %d/100 | Мотивация: %d/100\n", c.Knowledge, c.Motivation)
	fmt.Printf("   Режим отладки: %v\n", c.IsDebugMode)
	fmt.Printf("   Ошибок обработано: %d | Проигнорировано: %d\n", c.ErrorsHandled, c.ErrorsIgnored)
	fmt.Println(strings.Repeat("▬", 50))
}

// CheckErrors демонстрирует работу с ошибками
func (c *Courier) CheckErrors() {
	fmt.Println("\n🔎 ПРОВЕРКА ТИПОВ ОШИБОК:")

	// Тест SnowDepthError
	if err := c.Deliver("ул. Тестовая, 1", true); err != nil {
		if errors.Is(err, SnowDepthError) || strings.Contains(err.Error(), SnowDepthError.Error()) {
			fmt.Println("✅ Обнаружен SnowDepthError")
		}
	}

	// Тест ClientNotFoundErr
	if err := c.Deliver("", false); err != nil {
		if errors.Is(err, ClientNotFoundErr) || strings.Contains(err.Error(), ClientNotFoundErr.Error()) {
			fmt.Println("✅ Обнаружен ClientNotFoundErr")
		}
	}

	// Тест DeliveryFailedErr
	if err := c.Deliver("ул. Тестовая, 2", true); err != nil {
		if errors.Is(err, DeliveryFailedErr) || strings.Contains(err.Error(), DeliveryFailedErr.Error()) {
			fmt.Println("✅ Обнаружен DeliveryFailedErr")
		}
	}
}

// printMotivation — 10 мотивационных фраз для изучения Go
func printMotivation() {
	phrases := []string{
		"1. 🔥 Каждая ошибка в Go — шаг от 'почему?' к 'я знаю почему'",
		"2. 🧠 Стектрейс в коде = причинно-следственные связи в жизни",
		"3. 🛡️ Go учит принимать ответственность, а не прятаться",
		"4. ⚡ 100 дней Go = 100 дней перепрошивки 'я не могу' → 'я разберусь'",
		"5. 🎯 Debug режим — суперсила, превращающая проблемы в рост",
		"6. 💎 Контекст ошибок — записка себе из будущего с решением",
		"7. 🚀 Система типов Go — карта в метели багов",
		"8. 📈 Коммит с обработкой ошибок = +1 к ценности разработчика",
		"9. 🏆 Ошибка — не провал, а данные для правильного решения",
		"10. 🌅 Утро с `go test` лучше утра с 'опять эти сугробые'",
	}

	fmt.Println("\n" + strings.Repeat("💎", 30))
	fmt.Println("МОТИВАЦИЯ НА 84-Й ДЕНЬ:")
	for _, p := range phrases {
		fmt.Printf("   %s\n", p)
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Println(strings.Repeat("💎", 30))
}

// runExperiment демонстрирует разные сценарии
func runExperiment(title string, fn func()) {
	fmt.Println("\n" + strings.Repeat("═", 50))
	fmt.Println(title)
	fmt.Println(strings.Repeat("═", 50))
	fn()
}

func main() {
	fmt.Println("🌨️ SNOWTRACE DEBUGGER — День 84")
	fmt.Println("   Метафора: Как читать стектрейс своей жизни")

	gopher := NewCourier("Гоша")
	gopher.Status()
	printMotivation()

	// Эксперимент 1: Ошибка без отладки
	runExperiment("ЭКСПЕРИМЕНТ 1: Доставка БЕЗ отладки", func() {
		if err := gopher.Deliver("ул. Снежная, 15", true); err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			fmt.Println("❓ Где проблема? Как исправить? Непонятно.")
		}
	})

	// Эксперимент 2: С отладкой
	gopher.EnableDebug()
	runExperiment("ЭКСПЕРИМЕНТ 2: Доставка В РЕЖИМЕ ОТЛАДКИ", func() {
		if err := gopher.Deliver("ул. Ледяная, 8", true); err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			fmt.Println("✅ Виден контекст и стектрейс! Можно исправить.")
		}
	})

	// Эксперимент 3: Успешная доставка
	runExperiment("ЭКСПЕРИМЕНТ 3: Успешная доставка", func() {
		gopher.Deliver("ул. Тёплая, 3", false)
	})

	// Эксперимент 4: Клиент не найден
	runExperiment("ЭКСПЕРИМЕНТ 4: Клиент не найден", func() {
		if err := gopher.Deliver("", false); err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			fmt.Println("✅ Клиент не найден - стандартная бизнес-ошибка")
		}
	})

	// Эксперимент 5: Обработка паники
	runExperiment("ЭКСПЕРИМЕНТ 5: Обработка паники (recover)", func() {
		gopher.Motivation = 15
		gopher.HandlePanic()
	})

	// Проверка всех типов ошибок
	gopher.CheckErrors()

	// Статистика
	gopher.Status()

	// Финальный вывод
	fmt.Println("\n" + strings.Repeat("✨", 50))
	fmt.Println("ВЫВОД ДНЯ 84:")
	fmt.Println("   Ошибки в Go — учителя, а не враги.")
	fmt.Println("   Стектрейс — карта для навигации, а не обвинение.")
	fmt.Println("   Каждый день с Go — урок анализа вместо паники.")
	fmt.Println("   SnowDepthError и ClientNotFoundErr — примеры типизированных ошибок.")
	fmt.Println(strings.Repeat("✨", 50))

	// Явное использование всех переменных ошибок в выводе (для линтера)
	fmt.Printf("\n📋 ИСПОЛЬЗУЕМЫЕ ТИПЫ ОШИБОК:\n")
	fmt.Printf("   • %v\n", SnowDepthError)
	fmt.Printf("   • %v\n", ClientNotFoundErr)
	fmt.Printf("   • %v\n", DeliveryFailedErr)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
