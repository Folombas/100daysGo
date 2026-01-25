package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("❄️  День 82: Sentinel Errors - Стражи Холодного Кода")
	fmt.Println("==================================================")
	fmt.Println("Дата: 25 января 2026, воскресенье")
	fmt.Println("Миссия: Использовать Sentinel Errors для защиты фокуса")
	fmt.Println()

	// Инициализация системы
	thermostat := NewThermostat()
	resistance := NewResistanceSystem()

	// Стартовые условия
	fmt.Println("🌡️  СТАРТОВЫЕ УСЛОВИЯ:")
	fmt.Printf("- Комнатная температура: %d°C\n", thermostat.RoomTemp)
	fmt.Printf("- Фокус температура: %d°C\n", thermostat.FocusTemp)
	fmt.Printf("- Дисбаланс: %d°C\n", thermostat.GetImbalance())
	fmt.Println()

	// Симуляция дня
	events := []struct {
		time    string
		event   string
		tempChange int
	}{
		{"11:00", "Проснулся в ледяной комнате", -15},
		{"11:30", "Умывание ледяной водой", +5},
		{"12:00", "Завтрак", +10},
		{"13:00", "Заказик Химки→Выхино (+1000₽)", +20},
		{"15:00", "Пустая электричка домой", -5},
		{"17:00", "Поход в магазин", +5},
		{"19:00", "Ужин с макарошками", +15},
		{"20:00", "Горячий шиповник", +10},
		{"20:30", "ВКЛЮЧЕНИЕ КОМПЬЮТЕРА", +25},
	}

	for _, e := range events {
		fmt.Printf("⏰ %s: %s\n", e.time, e.event)
		thermostat.AdjustFocus(e.tempChange)
		time.Sleep(500 * time.Millisecond)

		// Проверка на искушения
		if challenge := resistance.CheckDailyChallenge(); challenge != nil {
			fmt.Printf("   ⚠️  Вызов: %s\n", challenge.Description)
			if success := resistance.FaceTemptation(challenge); success {
				fmt.Println("   ✅ Успешно сопротивлен!")
				thermostat.AdjustFocus(+15)
			}
		}
	}

	// Финальная проверка Sentinel Errors
	fmt.Println("\n🔍 ПРОВЕРКА SENTINEL ERRORS:")
	CheckAllSentinels(resistance)

	// Итоги
	fmt.Println("\n📊 ИТОГИ ДНЯ 82:")
	fmt.Printf("Финальная температура фокуса: %d°C\n", thermostat.FocusTemp)
	fmt.Printf("Преодолено искушений: %d\n", resistance.TemptationsResisted)
	fmt.Printf("Дофамин заработан: +%d\n", resistance.CalculateDopamine())

	if thermostat.FocusTemp >= 100 {
		fmt.Println("🎉 ПОБЕДА! Достигнут кипящий фокус!")
	} else {
		fmt.Println("💪 ХОРОШО! Завтра будет лучше!")
	}

	// Коммит
	fmt.Println("\n💾 Коммит: 'feat: Day 82 - освоил Sentinel Errors для защиты фокуса'")
}
