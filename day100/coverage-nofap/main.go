package main

import (
	"fmt"
	"time"
)

// NoFapWarrior представляет участника челленджа
type NoFapWarrior struct {
	Name        string
	Age         int
	DaysClean   int
	GoSkills    []string
	EnergyLevel float64
	Birthday    time.Time
}

// NewWarrior создает нового воина
func NewWarrior(name string, age int) *NoFapWarrior {
	return &NoFapWarrior{
		Name:        name,
		Age:         age,
		DaysClean:   0,
		GoSkills:    []string{},
		EnergyLevel: 50.0,
		Birthday:    time.Date(2024, 11, 30, 0, 0, 0, 0, time.UTC),
	}
}

// AddCleanDay добавляет день чистоты
func (w *NoFapWarrior) AddCleanDay() {
	w.DaysClean++
	w.EnergyLevel += 2.5

	// Каждые 3 дня изучаем новый навык Go
	if w.DaysClean%3 == 0 {
		skills := []string{"Go basics", "Testing", "Coverage", "Concurrency", "APIs", "Microservices"}
		if len(w.GoSkills) < len(skills) {
			w.GoSkills = append(w.GoSkills, skills[len(w.GoSkills)])
		}
	}
}

// GetCoverageReport возвращает отчет о покрытии жизни
func (w *NoFapWarrior) GetCoverageReport() string {
	lifeCoverage := float64(w.DaysClean) / 30.0 * 100 // Ноябрь = 30 дней

	var status string
	switch {
	case lifeCoverage >= 90:
		status = "ЭЛИТНЫЙ УРОВЕНЬ"
	case lifeCoverage >= 70:
		status = "ПРОДВИНУТЫЙ"
	case lifeCoverage >= 50:
		status = "СТАБИЛЬНЫЙ"
	default:
		status = "НАЧАЛЬНЫЙ"
	}

	return fmt.Sprintf("Покрытие жизни: %.1f%% | Уровень: %s", lifeCoverage, status)
}

// CalculateTransformation рассчитывает трансформацию
func (w *NoFapWarrior) CalculateTransformation() string {
	daysUntilBirthday := int(time.Until(w.Birthday).Hours() / 24)

	if w.DaysClean >= daysUntilBirthday {
		return "🎉 К 45 годам станешь SENIOR разработчиком!"
	}
	return "💪 Продолжай! Трансформация в процессе..."
}

func main() {
	fmt.Println("🚀 DAY 100: Coverage & NoFap Challenge!")
	fmt.Println("=========================================")

	goshа := NewWarrior("Гоша", 37)

	// Симуляция 15 дней челленджа (с 1 по 15 ноября)
	for day := 1; day <= 15; day++ {
		goshа.AddCleanDay()

		fmt.Printf("День %d ноября:\n", day)
		fmt.Printf("  Дней чистоты: %d\n", goshа.DaysClean)
		fmt.Printf("  Уровень энергии: %.1f\n", goshа.EnergyLevel)
		fmt.Printf("  Навыки Go: %v\n", goshа.GoSkills)
		fmt.Printf("  %s\n", goshа.GetCoverageReport())

		if day == 15 {
			fmt.Printf("  %s\n", goshа.CalculateTransformation())
			fmt.Println("  🎊 ПОЗДРАВЛЯЕМ С ЗАВЕРШЕНИЕМ 100 ДНЕЙ GO!")
		}
		fmt.Println()
	}

	fmt.Println("💡 НОВЫЙ ЧЕЛЛЕНДЖ: Ноябрь-Недрочабрь АКТИВИРОВАН!")
	fmt.Println("Цель: 100% coverage жизни кодом вместо adult-контента!")
}
