package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/GoReborn/day69/generics-type-interfaces-go/internal/delivery"
	"github.com/GoReborn/day69/generics-type-interfaces-go/internal/game"
	"github.com/GoReborn/day69/generics-type-interfaces-go/internal/legend"
	"github.com/GoReborn/day69/generics-type-interfaces-go/internal/trolls"
)

func main() {
	log.Println("🚀 Day 69: Generics Type Interfaces GO - ГОША REBORN")
	fmt.Println(strings.Repeat("=", 60))

	// Часть 1: Легенда дня
	fmt.Println("\n📖 ЛЕГЕНДА ДНЯ:")
	legend.PrintStory()

	// Часть 2: Геймификация
	fmt.Println("\n🎮 ГЕЙМИФИКАЦИЯ ДНЯ:")
	gamification := game.NewGamification()
	gamification.AddPoints(25, "Утренняя зарядка + завтрак с мамой")
	gamification.AddPoints(50, "Зарядил проездную карту Тройка безлимит + Пригород для путешествий по Ближайшему Подмосковью")
	gamification.AddPoints(75, "3 успешные доставки в снежную метель")
	gamification.AddPoints(100, "Программирование вместо бара с фриками")
	gamification.ShowProgress()

	// Часть 3: Дженерики в действии
	fmt.Println("\n🔧 GENERICS В ДЕЙСТВИИ:")
	fmt.Println("Демонстрация Type Parameters и Constraints...")

	// Система доставки с дженериками
	deliverySystem := delivery.NewSystem[string]()

	orders := []delivery.Order[string]{
		{ID: "ORD-001", Item: "Документы из МФЦ", Status: "delivered"},
		{ID: "ORD-002", Item: "Пакет с Авито", Status: "in_transit"},
		{ID: "ORD-003", Item: "Оборудование для офиса", Status: "pending"},
	}

	for _, order := range orders {
		deliverySystem.AddOrder(order)
	}

	fmt.Println("\n📦 СТАТУС ДОСТАВОК:")
	deliverySystem.ProcessOrders(func(o delivery.Order[string]) {
		fmt.Printf("  • %s: %s [%s]\n", o.ID, o.Item, o.Status)
	})

	// Часть 4: Игнорирование троллей (интерактивный режим)
	fmt.Println("\n🛡️  ИММУНИТЕТ К ТРОЛЛЯМ:")
	fmt.Println("Гоша тренирует игнор-мышцу...")

	trollShield := trolls.NewShield()
	fmt.Println("Тролли атакуют:")
	trollShield.PrintAttacks(3)

	fmt.Println("\nХейтеры нападают:")
	trollShield.PrintHateAttacks(3)

	fmt.Println("\n✅ Игнор-мышца прокачана! Достижение: 'Невозмутимый Кодер'")

	// Часть 5: Достижения дня
	fmt.Println("\n🏆 ДОСТИЖЕНИЯ ДНЯ 69:")
	achievements := []string{
		"🎖️  Преодоление непогоды",
		"🎖️  3 успешные доставки",
		"🎖️  Выбор кода вместо кутеха в барах",
		"🎖️  Освоение Generics",
		"🎖️  Иммунитет к троллям",
		"🔒  Скрытое достижение: 'Тройка безлимит на 30 дней + Пригород'",
	}

	for i, achievement := range achievements {
		fmt.Printf("  %d. %s\n", i+1, achievement)
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("💪 ГОША: 'Generics изучал, доставлял, игнорировал троллей.'")
	fmt.Println("💻 КОД: 'Скомпилировано успешно. Никаких undefined behavior!'")
}
