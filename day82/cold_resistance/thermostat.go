package main

import "fmt"

type Thermostat struct {
	RoomTemp  int // Температура комнаты
	FocusTemp int // Температура фокуса (0-100)
	Imbalance int // Дисбаланс между комфортом и фокусом
}

func NewThermostat() *Thermostat {
	return &Thermostat{
		RoomTemp:  -5, // Холодная комната после ночи с открытым балконом
		FocusTemp: 40, // Начальный фокус
		Imbalance: -45,
	}
}

func (t *Thermostat) AdjustFocus(change int) {
	oldTemp := t.FocusTemp
	t.FocusTemp += change

	// Ограничиваем диапазон
	if t.FocusTemp < 0 {
		t.FocusTemp = 0
	}
	if t.FocusTemp > 100 {
		t.FocusTemp = 100
	}

	// Обновляем дисбаланс
	t.Imbalance = t.FocusTemp - t.RoomTemp

	// Логирование изменений
	if change != 0 {
		trend := "↑"
		if change < 0 {
			trend = "↓"
		}
		fmt.Printf("   🌡️  Фокус: %d°C %s%d → %d°C (дисбаланс: %d°C)\n",
			oldTemp, trend, change, t.FocusTemp, t.Imbalance)
	}
}

func (t *Thermostat) GetImbalance() int {
	return t.Imbalance
}

func (t *Thermostat) GetFocusStatus() string {
	if t.FocusTemp >= 90 {
		return "КИПЯЩИЙ ФОКУС 🔥"
	} else if t.FocusTemp >= 70 {
		return "ГОРЯЧИЙ ФОКУС 🔥"
	} else if t.FocusTemp >= 50 {
		return "ТЕПЛЫЙ ФОКУС 🌡️"
	} else if t.FocusTemp >= 30 {
		return "ПРОХЛАДНЫЙ ФОКУС 💨"
	} else {
		return "ХОЛОДНЫЙ ФОКУС ❄️"
	}
}
