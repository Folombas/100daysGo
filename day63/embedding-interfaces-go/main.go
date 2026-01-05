package main

import (
	"fmt"
	"strings"
)

// Базовый курьер
type Courier struct {
	Name     string
	Speed    float64
	Rate     int
	Location string
}

func (c Courier) Move(dest string) string {
	return fmt.Sprintf("%s: %s → %s (%.1f км/ч)", c.Name, c.Location, dest, c.Speed)
}

func (c Courier) CalculateCost(distance float64) int {
	return int(distance) * c.Rate
}

// 1. Embedding структур: Велокурьер
type BikeCourier struct {
	Courier
	BikeType string
}

func (bc BikeCourier) Move(dest string) string {
	base := bc.Courier.Move(dest)
	return fmt.Sprintf("%s на %s велосипеде", base, bc.BikeType)
}

// 2. Embedding структур: Автокурьер
type CarCourier struct {
	Courier
	CarModel   string
	WinterTire bool
}

func (cc CarCourier) CheckWinterReady() string {
	if cc.WinterTire {
		return fmt.Sprintf("%s готов к гололёду", cc.CarModel)
	}
	return fmt.Sprintf("%s без зимней резины", cc.CarModel)
}

// 3. Интерфейс для полиморфизма
type Deliverer interface {
	Move(string) string
	CalculateCost(float64) int
}

func main() {
	fmt.Println("=== Day 63: Embedding Interfaces ===")
	fmt.Printf("Дата: 05.01.2026 | День 63/100\n")
	fmt.Printf("Контекст: Мороз, гололёд, 0 заказов\n\n")

	fmt.Println("1. ВЕЛО-КУРЬЕР (embedding структур):")
	bike := BikeCourier{
		Courier: Courier{"Гоша-велосипедист", 15.5, 50, "Ховрино"},
		BikeType: "городском",
	}
	fmt.Printf("   • %s\n", bike.Move("Реутов"))
	fmt.Printf("   • Стоимость 25 км: %d руб\n\n", bike.CalculateCost(25))

	fmt.Println("2. АВТО-КУРЬЕР (embedding + методы):")
	car := CarCourier{
		Courier:   Courier{"Гоша-автомобилист", 45.0, 100, "Ховрино"},
		CarModel:  "Lada Granta",
		WinterTire: true,
	}
	fmt.Printf("   • %s\n", car.Move("Реутов"))
	fmt.Printf("   • %s\n\n", car.CheckWinterReady())

	fmt.Println("3. ПОЛИМОРФИЗМ (интерфейсы):")
	deliverers := []Deliverer{bike, car}
	for i, d := range deliverers {
		fmt.Printf("   %d. %s\n", i+1, d.Move("Центр"))
		fmt.Printf("      Стоимость 30 км: %d руб\n", d.CalculateCost(30))
	}

	fmt.Println("\n" + strings.Repeat("-", 50))
	fmt.Println("КЛЮЧЕВЫЕ КОНЦЕПЦИИ EMBEDDING:")
	fmt.Println(strings.Repeat("-", 50))

	points := []string{
		"1. Композиция структур (не наследование)",
		"2. Методы встроенной структуры промоутятся",
		"3. Можно переопределять методы",
		"4. Можно добавлять новые методы",
		"5. Embedding интерфейсов = полиморфизм",
		"",
		"ПРИМЕНЕНИЕ:",
		"• Встроить навык Go в свою идентичность",
		"• Добавить 'зимнюю резину' знаний",
		"• Композиция скиллов > наследование диплома",
	}

	for _, p := range points {
		fmt.Println(p)
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("ВЫВОД: Embedding = зимняя резина для карьеры")
	fmt.Println(strings.Repeat("=", 50))
}
