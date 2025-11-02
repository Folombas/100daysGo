package main

import (
	"fmt"
	"time"
)

// LifeTransformation отслеживает трансформацию жизни
type LifeTransformation struct {
	Name           string
	Age            int
	DaysClean      int
	DailyStudyTime time.Duration
	Skills         []string
	Motivation     float64
	BankAccount    float64
}

// Константы - неизменные цели
const (
	TargetSalary    = 250000.0  // руб/мес
	TargetSkills    = 15        // навыков Go
	TransformationDays = 90     // дней для трансформации
)

// Глобальные переменные для отслеживания прогресса
var (
	currentDay    = 0
	totalProgress = 0.0
)

func main() {
	fmt.Println("🚀 DAY 100: Variables & Constants - Основа стабильности!")
	fmt.Println("========================================================")

	// Инициализация переменных
	goshа := LifeTransformation{
		Name:           "Гоша",
		Age:            37,
		DaysClean:      15, // Уже 15 дней без игр/сериалов
		DailyStudyTime: 3 * time.Hour,
		Skills:         []string{"variables", "constants", "functions", "structs"},
		Motivation:     85.5,
		BankAccount:    15000.50,
	}

	// Демонстрация работы с переменными
	fmt.Printf("👤 Имя: %s\n", goshа.Name)
	fmt.Printf("🎂 Возраст: %d лет\n", goshа.Age)
	fmt.Printf("📅 Дней без игр/сериалов: %d\n", goshа.DaysClean)
	fmt.Printf("⏰ Ежедневное обучение: %v\n", goshа.DailyStudyTime)
	fmt.Printf("💪 Мотивация: %.1f%%\n", goshа.Motivation)
	fmt.Printf("💰 Счёт в банке: %.2f руб\n", goshа.BankAccount)
	fmt.Printf("🛠 Навыки Go: %v\n", goshа.Skills)

	fmt.Println("\n🎯 КОНСТАНТЫ - НЕИЗМЕННЫЕ ЦЕЛИ:")
	fmt.Printf("Целевая зарплата: %.0f руб/мес\n", TargetSalary)
	fmt.Printf("Целевое количество навыков: %d\n", TargetSkills)
	fmt.Printf("Дней для трансформации: %d\n", TransformationDays)

	// Расчет прогресса
	progress := calculateProgress(&goshа)
	fmt.Printf("\n📊 ПРОГРЕСС ТРАНСФОРМАЦИИ: %.1f%%\n", progress)

	// Демонстрация изменения переменных
	fmt.Println("\n🔄 ИЗМЕНЕНИЕ ПЕРЕМЕННЫХ:")
	goshа.DaysClean++
	goshа.Motivation += 2.5
	goshа.BankAccount -= 2450.75 // траты на жизнь
	goshа.Skills = append(goshа.Skills, "interfaces")

	fmt.Printf("Новое количество дней: %d\n", goshа.DaysClean)
	fmt.Printf("Новая мотивация: %.1f%%\n", goshа.Motivation)
	fmt.Printf("Новый баланс: %.2f руб\n", goshа.BankAccount)
	fmt.Printf("Новые навыки: %v\n", goshа.Skills)

	// Константы остаются неизменными
	fmt.Println("\n⭐ КОНСТАНТЫ НЕИЗМЕННЫ:")
	fmt.Printf("Цель зарплаты всё ещё: %.0f руб\n", TargetSalary)

	showMemoryUsage()
}

func calculateProgress(lt *LifeTransformation) float64 {
	// Используем локальные переменные
	daysProgress := float64(lt.DaysClean) / float64(TransformationDays) * 50
	skillsProgress := float64(len(lt.Skills)) / float64(TargetSkills) * 30
	motivationProgress := lt.Motivation / 100 * 20

	totalProgress = daysProgress + skillsProgress + motivationProgress
	return totalProgress
}

func showMemoryUsage() {
	// Демонстрация разных типов переменных
	var (
		smallNumber  int8   = 127
		bigNumber    int64  = 9223372036854775807
		price        float32 = 299.99
		productName  string = "Курс Go Pro"
		isCompleted  bool   = false
	)

	fmt.Println("\n💾 ТИПЫ ПЕРЕМЕННЫХ И ПАМЯТЬ:")
	fmt.Printf("smallNumber (int8): %d\n", smallNumber)
	fmt.Printf("bigNumber (int64): %d\n", bigNumber)
	fmt.Printf("price (float32): %.2f\n", price)
	fmt.Printf("productName (string): %s\n", productName)
	fmt.Printf("isCompleted (bool): %t\n", isCompleted)

	fmt.Println("\n🎉 ВЫВОД: Переменные меняются, константы задают направление!")
	fmt.Println("Так и в жизни - доход переменный, а цели должны быть постоянными!")
}
