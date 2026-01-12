package game

import (
	"fmt"
	"strings"
)

type Achievement struct {
	Name        string
	Points      int
	Description string
	Unlocked    bool
}

type Gamification struct {
	TotalPoints  int
	Level        int
	Achievements []Achievement
	Dopamine     int // виртуальный дофамин
}

func NewGamification() *Gamification {
	return &Gamification{
		Level: 1,
		Achievements: []Achievement{
			{"Утренний ритуал", 0, "Зарядка + завтрак", false},
			{"Деловой путешественник", 0, "Заряженная проездная карта на месяц", false},
			{"Доставщик-профи", 0, "3+ доставки за день", false},
			{"Код-магнит", 0, "Программирование > развлечения", false},
			{"Generics-мастер", 0, "Освоил дженерики", false},
		},
		Dopamine: 50, // стартовый дофамин
	}
}

func (g *Gamification) AddPoints(points int, reason string) {
	oldPoints := g.TotalPoints
	g.TotalPoints += points
	g.Dopamine += points / 5 // дофамин за достижения

	fmt.Printf("  🎯 +%d очков: %s\n", points, reason)
	fmt.Printf("    💊 Дофамин: +%d (текущий: %d)\n", points/5, g.Dopamine)

	// Проверка новых уровней
	newLevel := g.TotalPoints/100 + 1
	if newLevel > g.Level {
		fmt.Printf("    ⬆️  НОВЫЙ УРОВЕНЬ: %d!\n", newLevel)
		g.Level = newLevel
		g.Dopamine += 20 // бонус дофамина за уровень
	}

	// Разблокировка достижений
	for i := range g.Achievements {
		if !g.Achievements[i].Unlocked && g.TotalPoints >= (i+1)*50 {
			g.Achievements[i].Unlocked = true
			fmt.Printf("    🏆 ДОСТИЖЕНИЕ: %s - %s\n",
				g.Achievements[i].Name, g.Achievements[i].Description)
		}
	}

	// Визуализация прогресса
	if oldPoints/100 != g.TotalPoints/100 {
		g.showProgressBar()
	}
}

func (g *Gamification) showProgressBar() {
	barWidth := 20
	progress := g.TotalPoints % 100
	if progress == 0 && g.TotalPoints > 0 {
		progress = 100
	}

	filled := progress * barWidth / 100
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

	fmt.Printf("    [%s] %d%% к уровню %d\n", bar, progress, g.Level+1)
}

func (g *Gamification) ShowProgress() {
	fmt.Printf("\n📊 ПРОГРЕСС:\n")
	fmt.Printf("  Очков: %d\n", g.TotalPoints)
	fmt.Printf("  Уровень: %d\n", g.Level)
	fmt.Printf("  Дофамин: %d/100\n", g.Dopamine)

	fmt.Println("\n🏆 ДОСТИЖЕНИЯ:")
	for _, a := range g.Achievements {
		status := "🔒"
		if a.Unlocked {
			status = "✅"
		}
		fmt.Printf("  %s %s\n", status, a.Name)
	}
}
